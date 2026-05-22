# direct-teams-bridge

`direct-teams-bridge` bridges multiple direct4b accounts to Microsoft Teams.

v1 assumes the first message starts on direct. The bridge creates one Teams channel thread for each direct talk, then sends later messages as replies in that thread. Teams replies are sent back to direct only when they are in a mapped thread and mention the configured bridge user.

## Configuration

Copy `config.example.yaml` and fill the Teams bot app ID.

Direct tokens are not stored in the config. Use either:

- `token_env` for runtime environment variables.
- `token_ref` for `login-direct` and local `op read` fallback.

## Bridge v1 Local Test Runbook

This runbook exposes the local bridge through Cloudflare Tunnel.
Use a stable Cloudflare public hostname for Azure Bot's messaging endpoint:

```text
https://bridge.example.com/api/messages
```

This document assumes Cloudflare Tunnel routes
`bridge.example.com` to `http://localhost:5173` on the host
running the bridge. A Quick Tunnel URL also works for a short smoke test, but
the URL is temporary and must be copied into the Azure Bot messaging endpoint
every time it changes.

### 1. Create the Teams bot registration

In Microsoft Entra / Azure Bot registration:

1. Create or open the Teams bot app registration.
2. Copy the Microsoft App ID.
3. Create a client secret / app password.
4. Set the Messaging endpoint to the Cloudflare Tunnel public hostname:

```text
https://bridge.example.com/api/messages
```

5. Ensure the bot is available to Microsoft Teams and installable into the target Team.

### 2. Prepare config

Create `config.yaml` from the example:

```bash
cp config.example.yaml config.yaml
```

For local testing, use a local state path and the Cloudflare Tunnel public URL:

```yaml
bot:
  app_id: "YOUR_MICROSOFT_APP_ID"
  app_password_env: "MICROSOFT_APP_PASSWORD"
  tenant_id: "YOUR_MICROSOFT_TENANT_ID"
  endpoint_path: "/api/messages"
  allow_emulator: false
  allowed_service_urls: []

teams_channels:
  support: {}
  trial: {}

accounts:
  - id: "account-a"
    token_env: "DIRECT_TOKEN_ACCOUNT_A"
    token_ref: "op://path/to/direct_access_token"
    endpoint: "wss://api.direct4b.com/albero-app-server/api"
    proxy_url: ""
    teams_channel: "support"

state:
  path: "./state/direct-teams-bridge.json"

server:
  listen_addr: ":5173"
  public_base_url: "https://bridge.example.com"

attachments:
  file_proxy_ttl: "24h"
```

If the Azure Bot is configured as single tenant, `tenant_id` is required so the
bridge requests Bot Connector tokens from your tenant instead of the legacy
`botframework.com` tenant. Without it, replies can fail with `AADSTS700016:
Application with identifier ... was not found in the directory 'Bot Framework'`.

Incoming Teams activities are authenticated with the Bot Framework JWT in the
`Authorization: Bearer ...` header. The bridge verifies the JWT signature,
issuer, audience, expiry, Teams channel endorsement, and that the token
`serviceurl` claim matches the activity `serviceUrl`.

`allowed_service_urls` is optional additional hardening. Leave it empty until the
first successful Teams activity shows the service URL you want to pin, then set
that exact URL.

`disable_auth_validation` is only accepted for local loopback development. It is
rejected when `server.public_base_url` is a public URL. For Bot Framework
Emulator testing, prefer `allow_emulator: true` instead.

Direct attachment links posted into Teams use `/files/direct` proxy URLs signed
with the Teams bot secret and an expiry controlled by `attachments.file_proxy_ttl`.

### 3. Prepare secrets

Either export raw values:

```bash
export MICROSOFT_APP_PASSWORD='...'
export DIRECT_TOKEN_ACCOUNT_A='...'
```

Or use 1Password secret references with host-side `op run`:

```env
DIRECT_TOKEN_ACCOUNT_A=op://path/to/direct_access_token
MICROSOFT_APP_PASSWORD=op://path/to/app_password
```

For first-time direct token creation:

```bash
PATH="$HOME/bin/go/bin:$PATH" go run ./cmd/direct-teams-bridge login-direct --config config.yaml --account account-a
```

### 3b. direct account token flow

The bridge does not store direct passwords. It uses the direct account password
only once to create an access token, then writes that token to the configured
1Password field.

Prepare the 1Password item and field first. The `token_ref` in `config.yaml`
must point to an existing item field:

```yaml
accounts:
  - id: "account-a"
    token_env: "DIRECT_TOKEN_ACCOUNT_A"
    token_ref: "op://path/to/direct_access_token"
    endpoint: "wss://api.direct4b.com/albero-app-server/api"
    teams_channel: "support"
```

Then run the interactive login command:

```bash
PATH="$HOME/bin/go/bin:$PATH" go run ./cmd/direct-teams-bridge login-direct --config config.yaml --account account-a
```

The command prompts for the direct account email and password:

```text
Email:
Password:
```

Internally it:

1. Connects to the direct WebSocket endpoint without a runtime access token.
2. Calls direct `create_access_token` with the entered email and password.
3. Extracts the returned access token.
4. Writes it to `token_ref` using `op item edit`.
5. Does not print the token to stdout.

At runtime, the bridge resolves the token in this order:

1. If `token_env` is set and the environment variable already has a value, use
   that value.
2. Otherwise, read `token_ref` with `op read`.

Each account also gets its own direct device ID in `state.direct_devices`.
This matches direct-js/daab's IDFV behavior and prevents a later login for one
account from invalidating another account's access token. If you need to rotate
the device ID as well as the token, run:

```bash
PATH="$HOME/bin/go/bin:$PATH" go run ./cmd/direct-teams-bridge login-direct --config config.yaml --account account-a --reset-device-id
```

For normal startup with 1Password service account token management:

```bash
export OP_SERVICE_ACCOUNT_TOKEN='...'
PATH="$HOME/bin/go/bin:$PATH" go run ./cmd/direct-teams-bridge run --config config.yaml
```

To rotate a direct token, rerun `login-direct` for that account. The running
bridge reloads `accounts` and `teams_channels` changes from `config.yaml`; if
only the 1Password token value changed, touch or resave `config.yaml` to trigger
the reload, or restart the bridge.

### 4. Start the bridge

From the repository root:

```bash
PATH="$HOME/bin/go/bin:$PATH" go run ./cmd/direct-teams-bridge run --config config.yaml
```

Expected log:

```text
[teams] bot endpoint listening on :5173/api/messages
[account-a] connected
```

### 4b. Dynamic config reload

The bridge polls `config.yaml` and reloads account routes and channel aliases
without restarting the process.

Reloaded live:

- `accounts` add/remove/change
- `teams_channels` alias add/remove
- direct token values resolved from `token_env` or `token_ref`

Restart required:

- Teams bot app settings
- server listen address / public base URL
- queue sizes
- state path

To add a new route while the bridge is running:

1. Add a new alias under `teams_channels`.
2. Add a new account under `accounts` with `teams_channel` set to that alias.
3. Save `config.yaml` and watch for `[config] reloaded ...`.
4. If the direct token has not been issued yet, the account is kept pending and
   the rest of the reload remains active.
5. Run `just login <account-id>` to create the token in 1Password.
6. Save or touch `config.yaml` again and watch for `[account-id] starting direct worker`.
7. In the target Teams channel, send `@direct bind <alias>`.

Pending token log example:

```text
[config] account pending token: account=account-a err=...
[config] reloaded accounts=3 active_accounts=2 pending_accounts=1 channels=3
```

To remove a Teams binding from a channel:

```text
@direct unbind <alias>
```

`unbind` must be sent as a new channel message. It removes the channel binding
and existing direct talk/thread mappings for that alias in that Teams channel.

### 5. Expose the bridge with Cloudflare Tunnel

For Azure Bot / Teams, the endpoint must be reachable from Microsoft's cloud.
`localhost`, LAN-only DNS, and private IP addresses such as `192.168.x.x` do not
work.

Recommended stable setup:

1. In Cloudflare Zero Trust, create a tunnel.
2. Add a Public Hostname:
   - Hostname: `bridge.example.com`
   - Service type: `HTTP`
   - URL: `localhost:5173`
3. Install or run the connector on the same machine as the bridge using the
   token command shown by Cloudflare.

The command usually looks like:

```bash
cloudflared tunnel run --token '<CLOUDFLARE_TUNNEL_TOKEN>'
```

If you use Docker on the VPS:

```bash
docker run --rm cloudflare/cloudflared:latest tunnel --no-autoupdate run --token '<CLOUDFLARE_TUNNEL_TOKEN>'
```

Check public routing:

```bash
curl -i https://bridge.example.com/healthz
```

Expected response:

```text
HTTP/2 200
ok
```

For a short local smoke test only, a Quick Tunnel can be used:

```bash
cloudflared tunnel --url http://localhost:5173
```

Copy the generated `https://*.trycloudflare.com/api/messages` URL into the Azure
Bot messaging endpoint. Do not use Quick Tunnel for normal operation because the
hostname is temporary.

### 6. Install and bind the Teams bot

Install the Teams app into the target Team. In the target channel, mention the bot once:

```text
@DirectBridge bind support
```

The bridge stores the Teams conversation reference in the state file. Confirm it:

```bash
PATH="$HOME/bin/go/bin:$PATH" go run ./cmd/direct-teams-bridge channels list --config config.yaml
```

### 7. Test direct to Teams

Send a message to the direct account configured as `account-a`.

Expected behavior:

1. The bridge receives the direct message.
2. It creates a Teams root thread in the bound `support` channel.
3. It stores the direct talk to Teams thread mapping.
4. Later direct messages in the same talk become Teams thread replies.

Confirm mapping:

```bash
PATH="$HOME/bin/go/bin:$PATH" go run ./cmd/direct-teams-bridge mappings list --config config.yaml
```

### 8. Test Teams to direct

Reply inside the mapped Teams thread and mention the bot:

```text
@DirectBridge reply text
```

Expected behavior:

1. Teams sends a message activity to `/api/messages`.
2. The bridge checks that it is in a mapped thread.
3. The bridge strips the bot mention.
4. The bridge sends the reply text to the mapped direct talk.

Unmapped Teams threads are ignored by design.

### 9. Reset local bindings if needed

Forget one direct talk mapping:

```bash
PATH="$HOME/bin/go/bin:$PATH" go run ./cmd/direct-teams-bridge mappings forget --config config.yaml --account account-a --talk-id DIRECT_TALK_ID
```

Forget the Teams channel binding:

```bash
PATH="$HOME/bin/go/bin:$PATH" go run ./cmd/direct-teams-bridge channels forget --config config.yaml --alias support
```

Then run `@DirectBridge bind support` again in Teams.

On a VPS, keep secret references in an env file:

```env
DIRECT_TOKEN_ACCOUNT_A=op://path/to/direct_access_token
MICROSOFT_APP_PASSWORD=op://path/to/app_password
```

Run the container through host-side 1Password CLI:

```bash
op run --env-file /etc/direct-teams-bridge/secrets.env -- \
  docker run --rm \
    --env DIRECT_TOKEN_ACCOUNT_A \
    --env MICROSOFT_APP_PASSWORD \
    -v /etc/direct-teams-bridge/config.yaml:/config.yaml:ro \
    -v /var/lib/direct-teams-bridge:/state \
    -p 127.0.0.1:5173:5173 \
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

Expose `/api/messages` over HTTPS with Caddy or Nginx and configure the Teams bot messaging endpoint to use that URL.

Install the Teams app into the target Team. Then mention the bot in the target channel once:

```text
@DirectBridge bind support
```

This saves the Teams conversation reference for the `support` alias. direct messages can only create Teams threads after the alias is bound.

## Mappings

```bash
direct-teams-bridge mappings list --config config.yaml
direct-teams-bridge mappings forget --config config.yaml --account account-a --talk-id 123
direct-teams-bridge channels list --config config.yaml
direct-teams-bridge channels forget --config config.yaml --alias support
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
  --env DIRECT_TOKEN_ACCOUNT_A --env MICROSOFT_APP_PASSWORD \
  -v /etc/direct-teams-bridge/config.yaml:/config.yaml:ro \
  -v /var/lib/direct-teams-bridge:/state \
  -p 127.0.0.1:5173:5173 \
  ghcr.io/OWNER/direct-teams-bridge@sha256:...
ExecStop=/usr/bin/docker stop direct-teams-bridge

[Install]
WantedBy=multi-user.target
```

## Limitations

- Unmapped Teams threads are ignored.
- Teams file attachments are best effort. If the bridge cannot download the file from the Bot activity URL, it sends a link.
- direct file upload uses the SDK's raw RPC helpers and may need adjustment if direct changes its file auth response shape.
