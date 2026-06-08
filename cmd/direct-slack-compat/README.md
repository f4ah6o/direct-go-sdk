# direct-slack-compat

`direct-slack-compat` exposes a small Slack-compatible Web API surface backed by Direct4B.

This is an adapter for simple Slack bot implementations. It is not a full Slack clone.

## Supported subset

- `GET/POST /api/auth.test`
- `POST /api/chat.postMessage`
- `GET/POST /api/conversations.list`
- `GET/POST /api/conversations.history`
- `GET/POST /api/users.list`
- Direct incoming messages can be converted to Slack Events API-like `event_callback` payloads when `slack.event_callback_url` is configured.

## Non-goals

- Slack App OAuth install flow
- Socket Mode
- Block Kit compatibility
- Modals
- Full thread semantics
- Interactive components

## Run

```bash
cp slackcompat.config.example.yaml slackcompat.config.yaml
# Edit accounts and token refs.
go run ./cmd/direct-slack-compat --config slackcompat.config.yaml
```

Access tokens can come from environment variables or `op://...` references. Secret values are read at runtime and are not printed.

The default listener is `127.0.0.1:8091`. Configure `server.bearer_token_env` or
`server.bearer_token_ref` before exposing the adapter beyond localhost; API
requests then require `Authorization: Bearer <token>`.

Environment variables in the YAML file are expanded before parsing, matching
the repository's other bridge config loaders.

## Examples

List conversations:

```bash
curl http://localhost:8091/api/conversations.list
```

Send a message:

```bash
curl -X POST http://localhost:8091/api/chat.postMessage \
  -H 'Authorization: Bearer ...' \
  -d channel='C...' \
  -d text='hello from a Slack-compatible client'
```

The channel ID must come from `conversations.list`. It encodes the Direct account and talk ID with a deterministic prefix-based mapping.
