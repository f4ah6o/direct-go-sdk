#!/usr/bin/env bash
set -euo pipefail

ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/../../.." && pwd)
APP_DIR="$ROOT/bench/runtime/.generated/node-daab"

if [[ ! -d "$APP_DIR" ]]; then
  exit 0
fi

(
  cd "$APP_DIR"
  if [[ -x node_modules/.bin/forever ]]; then
    node_modules/.bin/forever stop node-daab >/dev/null 2>&1 || true
  fi
)

pkill -f "$APP_DIR/node_modules/.bin/hubot -a direct" >/dev/null 2>&1 || true
pkill -f "$APP_DIR/node_modules/.bin/forever --uid node-daab" >/dev/null 2>&1 || true
