#!/usr/bin/env bash
set -euo pipefail

ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/../../.." && pwd)
GO_OUT="$ROOT/bench/runtime/.generated/go"
NODE_APP="$ROOT/bench/runtime/.generated/node-daab"

printf "artifact,bytes,path\n"
if [[ -f "$GO_OUT/runtime-go-ping" ]]; then
  printf "go_binary,%s,%s\n" "$(stat -c %s "$GO_OUT/runtime-go-ping")" "$GO_OUT/runtime-go-ping"
fi
if [[ -f "$GO_OUT/runtime-go-ping-stripped" ]]; then
  printf "go_binary_stripped,%s,%s\n" "$(stat -c %s "$GO_OUT/runtime-go-ping-stripped")" "$GO_OUT/runtime-go-ping-stripped"
fi
if [[ -f "$GO_OUT/direct-post" ]]; then
  printf "go_direct_post_binary,%s,%s\n" "$(stat -c %s "$GO_OUT/direct-post")" "$GO_OUT/direct-post"
fi
if [[ -f "$GO_OUT/direct-post-stripped" ]]; then
  printf "go_direct_post_binary_stripped,%s,%s\n" "$(stat -c %s "$GO_OUT/direct-post-stripped")" "$GO_OUT/direct-post-stripped"
fi
if [[ -d "$NODE_APP" ]]; then
  printf "node_daab_app_du,%s,%s\n" "$(du -sb "$NODE_APP" | awk '{print $1}')" "$NODE_APP"
fi
if [[ -d "$NODE_APP/node_modules" ]]; then
  printf "node_modules_du,%s,%s\n" "$(du -sb "$NODE_APP/node_modules" | awk '{print $1}')" "$NODE_APP/node_modules"
fi
