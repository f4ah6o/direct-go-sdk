# Runtime comparison: daab-go vs daab

This directory contains a reproducible harness for comparing the Go SDK bot
runtime against upstream `daab` on a real Direct4B connection.

The comparison target is the smallest equivalent ping bot:

- Go: `bench/runtime/go-ping`
- Node/daab: generated under `bench/runtime/.generated/node-daab` from
  `daab@0.7.0`

## Prerequisites

- Go 1.25+
- Node.js 18+ and npm 10+
- A Direct4B bot token prepared by `daab login` in the generated Node/daab app
- Optional: `HUBOT_DIRECT_ENDPOINT`, `HUBOT_DIRECT_PROXY_URL`

Do not commit generated files under `bench/runtime/.generated` or result files
under `bench/runtime/results`.

## Setup

Build the Go ping bot:

```bash
bench/runtime/scripts/build-go.sh
```

Generate the Node/daab ping bot:

```bash
bench/runtime/scripts/setup-node-daab.sh
```

Log in with upstream daab. This writes `HUBOT_DIRECT_TOKEN` to
`bench/runtime/.generated/node-daab/.env`:

```bash
cd bench/runtime/.generated/node-daab
node_modules/.bin/daab login
```

Record artifact sizes:

```bash
bench/runtime/scripts/size-report.sh | tee bench/runtime/results/size-report.csv
```

## Posting during measurement

Use the Go SDK sender to include message posting in the same sampling window.
The bot under test should use the `HUBOT_DIRECT_TOKEN` in
`bench/runtime/.generated/node-daab/.env`; the sender should use a different
account, normally `bot-trial2` from `config.yaml`. `direct-post` first checks
the account's `token_env`, then falls back to `op read <token_ref>`.

Required environment:

```bash
export DIRECT_BENCH_TALK_ID='<direct-talk-id>'
```

Manual sender check:

```bash
bench/runtime/.generated/go/direct-post-stripped \
  --config config.yaml \
  --account bot-trial2 \
  --talk-id "$DIRECT_BENCH_TALK_ID" \
  --text ping
```

## Live lookup test

Use the Go test below to verify Direct API lookup behavior against the real
service. It resolves:

- `talk_id -> room name + domain_id` through `get_talks`
- `domain_id + user_id -> display_name/name` through `get_users`

The test uses the same account config and token resolution as `direct-post`.
It is skipped unless `DIRECT_LOOKUP_LIVE=1` is set.

```bash
cd bench/runtime/go-ping
DIRECT_LOOKUP_LIVE=1 \
DIRECT_LOOKUP_CONFIG=../../../config.yaml \
DIRECT_LOOKUP_ACCOUNT=bot-trial2 \
DIRECT_LOOKUP_TALK_ID="$DIRECT_BENCH_TALK_ID" \
go test -run TestLiveDirectLookupResolvesUserAndRoomNames -v
```

If `DIRECT_LOOKUP_USER_ID` is omitted, the test tries to choose a non-self user
from the target talk's `user_ids`.

## Memory, CPU, and startup measurement

Run both processes from `bench/runtime/.generated/node-daab` so they read the
same `.env` created by `daab login`.

Go:

```bash
cd bench/runtime/.generated/node-daab
../../scripts/measure-process.sh \
  go-ping 'READY runtime-go-ping' 30 \
  ../go/runtime-go-ping-stripped
```

Go, including three Go SDK posts at sample second 3. The run stops early when
three PONG sends are logged, otherwise it stops at 30 seconds:

```bash
cd bench/runtime/.generated/node-daab
export BENCH_PING_TEXT="bench-ping-$(date +%s)"
BENCH_POST_AFTER_SECONDS=3 \
BENCH_POST_COMMAND='../go/direct-post-stripped --config ../../../../config.yaml --account bot-trial2 --talk-id "$DIRECT_BENCH_TALK_ID" --text "$BENCH_PING_TEXT" --count 3 --interval 10s --timeout 90s' \
BENCH_DONE_REGEX='RUNTIME pong sent' \
BENCH_DONE_COUNT=3 \
../../scripts/measure-process.sh \
  go-ping-with-post 'READY runtime-go-ping' 30 \
  ../go/runtime-go-ping-stripped
```

Node/daab:

```bash
cd bench/runtime/.generated/node-daab
../../scripts/measure-process.sh \
  node-daab 'READY runtime-node-daab' 30 \
  bash -lc 'export PATH=/opt/codex-desktop/resources/node-runtime/bin:$PATH; touch hubot.log; tail -n +1 -F hubot.log & env DISABLE_NPM_INSTALL=1 node_modules/.bin/daab run'
```

Node/daab, including the same three Go SDK posts at sample second 3. The run
stops early when three PONG sends are logged, otherwise it stops at 30 seconds:

```bash
cd bench/runtime/.generated/node-daab
export BENCH_PING_TEXT="bench-ping-$(date +%s)"
BENCH_POST_AFTER_SECONDS=3 \
BENCH_POST_COMMAND='../go/direct-post-stripped --config ../../../../config.yaml --account bot-trial2 --talk-id "$DIRECT_BENCH_TALK_ID" --text "$BENCH_PING_TEXT" --count 3 --interval 10s --timeout 90s' \
BENCH_DONE_REGEX='RUNTIME pong sent' \
BENCH_DONE_COUNT=3 \
../../scripts/measure-process.sh \
  node-daab-with-post 'READY runtime-node-daab' 30 \
  bash -lc 'export PATH=/opt/codex-desktop/resources/node-runtime/bin:$PATH; touch hubot.log; tail -n +1 -F hubot.log & env DISABLE_NPM_INSTALL=1 node_modules/.bin/daab run'
```

After every Node/daab run, stop the `forever` supervisor before starting the Go
bot or another daab measurement:

```bash
../../scripts/stop-node-daab.sh
```

Each run writes:

- `bench/runtime/results/<timestamp>-<name>.summary`
- `bench/runtime/results/<timestamp>-<name>.csv`
- `bench/runtime/results/<timestamp>-<name>.log`

Run each measurement at least three times and use the median as the headline
number. Keep min/max as supporting context because real Direct4B connection
latency and server state can vary.

Summarize a CSV:

```bash
bench/runtime/scripts/summarize-csv.sh bench/runtime/results/<run>.csv
```

Summaries include average, median, min/max, and last sample values. Summarize
steady state by dropping the first 5 seconds of startup warmup:

```bash
bench/runtime/scripts/summarize-csv.sh bench/runtime/results/<run>.csv 5
```

Report both numbers when comparing runtimes:

- Full-window average: includes startup, script loading, and first connection work.
- Steady-state average: excludes startup warmup and better reflects idle/runtime cost.

## Metrics to compare

- Runtime memory footprint: process-tree RSS and PSS after 30 seconds idle
- Executable/distribution size: Go binary size, stripped Go binary size,
  Node app size, and `node_modules` size
- Startup time: process start to ready log
- Idle CPU: average CPU percent over the sampling window
- Optional ping-path check: send `ping` in the measured room and confirm the
  process replies with `PONG`

## Browser E2E check

Use Chrome for the user-facing E2E check after `daab login` has created the
token and one of the measured bots is running:

1. Open the Direct4B web app in Chrome.
2. Log in with the same Direct4B account that can access the bot test room.
3. Open the test room, send `ping`, and confirm the bot replies `PONG`.
4. Keep the measurement script running during this check so the CPU/RSS samples
   include the message handling path.

## Notes

The Node/daab command can spawn child processes. The measurement script follows
the process tree and sums child RSS/PSS/CPU so the comparison includes the full
runtime cost. `daab run` also uses `forever`, so it can leave a supervised Hubot
process alive after the measurement wrapper exits.

`DISABLE_NPM_INSTALL=1` is intentional. The generated `bin/hubot run` otherwise
runs `npm install` every time, which mixes install/build work into runtime
measurement. If `libxmljs` is missing after an interrupted install, run
`npm rebuild libxmljs` once in `bench/runtime/.generated/node-daab`.

The command also prepends the Codex-bundled Node.js 22 runtime to `PATH`.
`libxmljs` is a native module, so the Node.js version used for `npm install` /
`npm rebuild` and the Node.js version used for `daab run` must match.
