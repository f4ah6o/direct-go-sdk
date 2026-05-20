# direct-teams-bridge

`direct-teams-bridge` bridges multiple direct4b accounts to Microsoft Teams.

v1 assumes the first message starts on direct. The bridge creates one Teams channel thread for each direct talk, then sends later messages as replies in that thread. Teams replies are sent back to direct only when they are in a mapped thread and mention the configured bridge user.

## Configuration

Copy `config.example.yaml` and fill the Graph and Teams IDs.

Direct tokens are not stored in the config. Use either:

- `token_env` for runtime environment variables.
- `token_ref` for `login-direct` and local `op read` fallback.

On a VPS, keep secret references in an env file:

```env
DIRECT_TOKEN_ACCOUNT_A=op://path/to/direct_access_token
GRAPH_ACCESS_TOKEN=op://path/to/graph_access_token
GRAPH_CLIENT_SECRET=op://path/to/graph_client_secret
```

Run the container through host-side 1Password CLI:

```bash
op run --env-file /etc/direct-teams-bridge/secrets.env -- \
  docker run --rm \
    --env DIRECT_TOKEN_ACCOUNT_A \
    --env GRAPH_ACCESS_TOKEN \
    --env GRAPH_CLIENT_SECRET \
    -v /etc/direct-teams-bridge/config.yaml:/config.yaml:ro \
    -v /var/lib/direct-teams-bridge:/state \
    -p 127.0.0.1:8080:8080 \
    ghcr.io/OWNER/direct-teams-bridge@sha256:...
```

## Login

Create or rotate a direct token and save it to 1Password:

```bash
direct-teams-bridge login-direct --config config.yaml --account account-a
```

The command calls direct `create_access_token` and writes the token to the configured `token_ref`. It does not print the token.

## Run

```bash
direct-teams-bridge run --config config.yaml
```

Expose `/graph/notifications` over HTTPS with Caddy or Nginx and configure Microsoft Graph change notifications to use that URL.

For normal Teams message send, prefer a delegated Graph access token supplied through `GRAPH_ACCESS_TOKEN`. The client credentials path is kept for Graph operations that support application permissions.

## Mappings

```bash
direct-teams-bridge mappings list --config config.yaml
direct-teams-bridge mappings forget --config config.yaml --account account-a --talk-id 123
```

## Build With ko

```bash
export KO_DOCKER_REPO=ghcr.io/OWNER/direct-teams-bridge
ko build ./cmd/direct-teams-bridge --bare
```

For a local Docker daemon build:

```bash
ko build --local ./cmd/direct-teams-bridge
```

## systemd Sketch

```ini
[Unit]
Description=direct Teams bridge
After=docker.service network-online.target
Wants=network-online.target

[Service]
Restart=always
Environment=OP_SERVICE_ACCOUNT_TOKEN=...
ExecStart=/usr/bin/op run --env-file /etc/direct-teams-bridge/secrets.env -- /usr/bin/docker run --rm --name direct-teams-bridge \
  --env DIRECT_TOKEN_ACCOUNT_A --env GRAPH_ACCESS_TOKEN --env GRAPH_CLIENT_SECRET \
  -v /etc/direct-teams-bridge/config.yaml:/config.yaml:ro \
  -v /var/lib/direct-teams-bridge:/state \
  -p 127.0.0.1:8080:8080 \
  ghcr.io/OWNER/direct-teams-bridge@sha256:...
ExecStop=/usr/bin/docker stop direct-teams-bridge

[Install]
WantedBy=multi-user.target
```

## Limitations

- Unmapped Teams threads are ignored.
- Teams file reference attachments are best effort. If the bridge cannot download the file through Graph, it sends a link.
- direct file upload uses the SDK's raw RPC helpers and may need adjustment if direct changes its file auth response shape.
