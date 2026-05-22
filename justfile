set dotenv-load := false

go := env_var_or_default("GO", env_var("HOME") + "/bin/go/bin/go")
config := env_var_or_default("CONFIG", "config.yaml")
account := env_var_or_default("DIRECT_ACCOUNT", "bot-trial")
op_item := env_var_or_default("OP_ITEM", "direct-teams-bridge-example")

default:
    @just --list

# Create or rotate the direct access token for an account and save it to token_ref.
login account +extra:
    PATH="{{ env_var("HOME") }}/bin/go/bin:$PATH" {{ go }} run ./cmd/direct-teams-bridge login-direct --config {{ config }} --account {{ account }} {{ extra }}

# Run the bridge directly with the current shell environment.
run config=config:
    PATH="{{ env_var("HOME") }}/bin/go/bin:$PATH" {{ go }} run ./cmd/direct-teams-bridge run --config {{ config }}

# Run the bridge through opz, loading secrets from the configured 1Password item.
run-op item=op_item config=config:
    opz run {{ item }} -- go run ./cmd/direct-teams-bridge run --config {{ config }}

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
