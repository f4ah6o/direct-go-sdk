#!/usr/bin/env bash
set -euo pipefail

ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/../../.." && pwd)
APP_DIR="$ROOT/bench/runtime/.generated/node-daab"

rm -rf "$APP_DIR"
mkdir -p "$APP_DIR"

(
  cd "$APP_DIR"
  npm init -y >/dev/null
  npm install daab@0.7.0 >/dev/null
  npx daab init --no-prompt >/dev/null
  npm pkg set devDependencies.daab=0.7.0 >/dev/null
  npm install >/dev/null
  cp "$ROOT/bench/runtime/node-daab/runtime-ping.js" scripts/runtime-ping.js
  rm -f scripts/ping.js
)

echo "$APP_DIR"
