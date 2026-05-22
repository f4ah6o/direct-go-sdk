#!/usr/bin/env bash
set -euo pipefail

root_pid=${1:?pid is required}

walk() {
  local pid=$1
  if ! kill -0 "$pid" 2>/dev/null; then
    return
  fi
  echo "$pid"
  local child
  while read -r child; do
    if [[ -n "$child" ]]; then
      walk "$child"
    fi
  done < <(pgrep -P "$pid" 2>/dev/null || true)
}

walk "$root_pid" | sort -n | uniq
