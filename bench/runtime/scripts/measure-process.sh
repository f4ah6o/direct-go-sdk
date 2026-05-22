#!/usr/bin/env bash
set -euo pipefail

if [[ $# -lt 4 ]]; then
  echo "usage: $0 <name> <ready-regex> <sample-seconds> <command> [args...]" >&2
  exit 2
fi

name=$1
ready_regex=$2
sample_seconds=$3
shift 3

ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/../../.." && pwd)
OUT_DIR="$ROOT/bench/runtime/results"
mkdir -p "$OUT_DIR"

stamp=$(date -u +"%Y%m%dT%H%M%SZ")
log_file="$OUT_DIR/${stamp}-${name}.log"
csv_file="$OUT_DIR/${stamp}-${name}.csv"
summary_file="$OUT_DIR/${stamp}-${name}.summary"

start_ns=$(date +%s%N)
"$@" >"$log_file" 2>&1 &
pid=$!

cleanup() {
  if kill -0 "$pid" 2>/dev/null; then
    kill "$pid" 2>/dev/null || true
    sleep 2
    kill -9 "$pid" 2>/dev/null || true
  fi
}
trap cleanup EXIT

ready_seconds=""
for _ in $(seq 1 120); do
  if ! kill -0 "$pid" 2>/dev/null; then
    echo "process exited before ready; see $log_file" >&2
    exit 1
  fi
  if grep -Eq "$ready_regex" "$log_file"; then
    now_ns=$(date +%s%N)
    ready_seconds=$(awk "BEGIN { printf \"%.3f\", ($now_ns - $start_ns) / 1000000000 }")
    break
  fi
  sleep 1
done

if [[ -z "$ready_seconds" ]]; then
  echo "ready pattern was not observed within 120s; see $log_file" >&2
  exit 1
fi

echo "second,pids,rss_kb,pss_kb,vsz_kb,cpu_percent" >"$csv_file"
post_ran=0
done_reason="timeout"
for second in $(seq 1 "$sample_seconds"); do
  if [[ "${BENCH_POST_COMMAND:-}" != "" && "$post_ran" -eq 0 && "$second" -ge "${BENCH_POST_AFTER_SECONDS:-3}" ]]; then
    {
      echo "---- BENCH_POST_COMMAND start second=$second ----"
      bash -lc "$BENCH_POST_COMMAND"
      echo "---- BENCH_POST_COMMAND end second=$second ----"
    } >>"$log_file" 2>&1 &
    post_ran=1
  fi
  mapfile -t pids < <("$ROOT/bench/runtime/scripts/process-tree.sh" "$pid")
  rss=0
  pss=0
  vsz=0
  cpu=0
  for p in "${pids[@]}"; do
    if [[ -r "/proc/$p/smaps_rollup" ]]; then
      pss=$((pss + $(awk '/^Pss:/ {print $2}' "/proc/$p/smaps_rollup")))
    fi
    read -r prss pvsz pcpu < <(ps -p "$p" -o rss=,vsz=,%cpu= 2>/dev/null || echo "0 0 0")
    rss=$((rss + ${prss:-0}))
    vsz=$((vsz + ${pvsz:-0}))
    cpu=$(awk "BEGIN { printf \"%.2f\", $cpu + ${pcpu:-0} }")
  done
  printf "%s,%s,%s,%s,%s,%s\n" "$second" "${pids[*]}" "$rss" "$pss" "$vsz" "$cpu" >>"$csv_file"
  if [[ "${BENCH_DONE_REGEX:-}" != "" ]]; then
    done_count=$(grep -Ec "$BENCH_DONE_REGEX" "$log_file" || true)
    if [[ "$done_count" -ge "${BENCH_DONE_COUNT:-1}" ]]; then
      done_reason="done_regex:${BENCH_DONE_REGEX}:$done_count"
      break
    fi
  fi
  sleep 1
done

{
  echo "name=$name"
  echo "pid=$pid"
  echo "ready_seconds=$ready_seconds"
  echo "log_file=$log_file"
  echo "csv_file=$csv_file"
  echo "done_reason=$done_reason"
  echo "go_version=$(go version 2>/dev/null || true)"
  echo "node_version=$(node --version 2>/dev/null || true)"
  echo "npm_version=$(npm --version 2>/dev/null || true)"
  echo "uname=$(uname -a)"
} >"$summary_file"

cat "$summary_file"
