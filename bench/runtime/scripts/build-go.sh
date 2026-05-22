#!/usr/bin/env bash
set -euo pipefail

ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/../../.." && pwd)
OUT="$ROOT/bench/runtime/.generated/go"

mkdir -p "$OUT"
(
  cd "$ROOT/bench/runtime/go-ping"
  go mod tidy
  go build -o "$OUT/runtime-go-ping" .
  go build -ldflags="-s -w" -o "$OUT/runtime-go-ping-stripped" .
  go build -o "$OUT/direct-post" ./cmd/direct-post
  go build -ldflags="-s -w" -o "$OUT/direct-post-stripped" ./cmd/direct-post
)

echo "$OUT"
