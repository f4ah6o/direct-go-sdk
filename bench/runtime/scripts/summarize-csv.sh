#!/usr/bin/env bash
set -euo pipefail

if [[ $# -lt 1 ]]; then
  echo "usage: $0 <csv> [warmup-seconds]" >&2
  exit 2
fi

csv=$1
warmup_seconds=${2:-0}

awk -F, -v warmup="$warmup_seconds" '
NR == 1 { next }
$1 <= warmup { next }
{
  n++
  rss_values[n] = $3 + 0
  pss_values[n] = $4 + 0
  vsz_values[n] = $5 + 0
  cpu_values[n] = $6 + 0
  rss += rss_values[n]
  pss += pss_values[n]
  vsz += vsz_values[n]
  cpu += cpu_values[n]
  if (n == 1 || rss_values[n] < min_rss) min_rss = rss_values[n]
  if (n == 1 || rss_values[n] > max_rss) max_rss = rss_values[n]
  if (n == 1 || pss_values[n] < min_pss) min_pss = pss_values[n]
  if (n == 1 || pss_values[n] > max_pss) max_pss = pss_values[n]
  if (n == 1 || cpu_values[n] > max_cpu) max_cpu = cpu_values[n]
  last_second = $1
  last_rss = rss_values[n]
  last_pss = pss_values[n]
  last_vsz = vsz_values[n]
  last_cpu = cpu_values[n]
}
END {
  if (n == 0) {
    printf "samples=0 warmup_seconds=%d\n", warmup
    exit
  }
  printf "samples=%d warmup_seconds=%d last_second=%d avg_rss_kb=%.0f median_rss_kb=%.0f avg_pss_kb=%.0f median_pss_kb=%.0f avg_vsz_kb=%.0f median_vsz_kb=%.0f avg_cpu=%.2f median_cpu=%.2f min_rss_kb=%d max_rss_kb=%d min_pss_kb=%d max_pss_kb=%d max_cpu=%.2f last_rss_kb=%d last_pss_kb=%d last_vsz_kb=%d last_cpu=%.2f\n",
    n, warmup, last_second,
    rss / n, median(rss_values, n),
    pss / n, median(pss_values, n),
    vsz / n, median(vsz_values, n),
    cpu / n, median(cpu_values, n),
    min_rss, max_rss, min_pss, max_pss, max_cpu, last_rss, last_pss, last_vsz, last_cpu
}

function median(values, count, sorted, i, j, tmp) {
  for (i = 1; i <= count; i++) {
    sorted[i] = values[i]
  }
  for (i = 1; i <= count; i++) {
    for (j = i + 1; j <= count; j++) {
      if (sorted[j] < sorted[i]) {
        tmp = sorted[i]
        sorted[i] = sorted[j]
        sorted[j] = tmp
      }
    }
  }
  if (count % 2 == 1) {
    return sorted[(count + 1) / 2]
  }
  return (sorted[count / 2] + sorted[count / 2 + 1]) / 2
}
' "$csv"
