set dotenv-load := false

go := env_var_or_default("GO", env_var("HOME") + "/bin/go/bin/go")
config := env_var_or_default("CONFIG", "config.yaml")
account := env_var_or_default("DIRECT_ACCOUNT", "bot-trial")
lookup_account := env_var_or_default("DIRECT_LOOKUP_ACCOUNT", "bot-trial2")
lookup_talk_id := env_var_or_default("DIRECT_LOOKUP_TALK_ID", env_var_or_default("DIRECT_BENCH_TALK_ID", ""))
op_item := env_var_or_default("OP_ITEM", "direct-teams-bridge-example")

default:
    @just --list

# Create or rotate the direct access token for an account and save it to token_ref.
login account *extra:
    PATH="{{ env_var("HOME") }}/bin/go/bin:$PATH" {{ go }} run ./cmd/direct-teams-bridge login-direct --config {{ config }} --account {{ account }} {{ extra }}

# Run the bridge directly with the current shell environment.
run config=config:
    PATH="{{ env_var("HOME") }}/bin/go/bin:$PATH" {{ go }} run ./cmd/direct-teams-bridge run --config {{ config }}

# Run the bridge through opz, loading secrets from the configured 1Password item.
run-op item=op_item config=config:
    opz run {{ item }} -- go run ./cmd/direct-teams-bridge run --config {{ config }}

# Build the bridge, then run the binary through opz with 1Password secrets.
run-op-bin item=op_item config=config:
    PATH="{{ env_var("HOME") }}/bin/go/bin:$PATH" {{ go }} build -o ./bin/direct-teams-bridge ./cmd/direct-teams-bridge
    opz run {{ item }} -- ./bin/direct-teams-bridge run --config {{ config }}

# List Teams channel bindings.
channels config=config:
    PATH="{{ env_var("HOME") }}/bin/go/bin:$PATH" {{ go }} run ./cmd/direct-teams-bridge channels list --config {{ config }}

# List direct talk to Teams thread mappings.
mappings config=config:
    PATH="{{ env_var("HOME") }}/bin/go/bin:$PATH" {{ go }} run ./cmd/direct-teams-bridge mappings list --config {{ config }}

# Run root bridge tests.
test:
    PATH="{{ env_var("HOME") }}/bin/go/bin:$PATH" {{ go }} test ./...

# Run direct-go SDK tests.
test-direct:
    cd direct-go && PATH="{{ env_var("HOME") }}/bin/go/bin:$PATH" {{ go }} test ./...

# Run the live Direct lookup test for talk/user display name resolution.
test-live-lookup account=lookup_account talk_id=lookup_talk_id config=config:
    test -n "{{ talk_id }}" || (echo "DIRECT_LOOKUP_TALK_ID or DIRECT_BENCH_TALK_ID is required" >&2; exit 1)
    cd bench/runtime/go-ping && PATH="{{ env_var("HOME") }}/bin/go/bin:$PATH" DIRECT_LOOKUP_LIVE=1 DIRECT_LOOKUP_CONFIG=../../../{{ config }} DIRECT_LOOKUP_ACCOUNT={{ account }} DIRECT_LOOKUP_TALK_ID={{ talk_id }} {{ go }} test -run TestLiveDirectLookupResolvesUserAndRoomNames -v
