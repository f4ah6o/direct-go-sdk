# AGENTS.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a Go monorepo containing two main modules for Direct4B (direct4b.com) chat platform:

1. **direct-go**: Go SDK for Direct4B WebSocket/MessagePack RPC API
2. **daab-go**: CLI tool and bot framework built on top of direct-go (inspired by Hubot)

Both modules are being ported from upstream JavaScript implementations:
* direct-go ports from [lisb/direct-js](https://github.com/lisb/direct-js)
* daab-go ports from [lisb/daab](https://github.com/lisb/daab)

## Module Structure

```
direct-go-sdk/
├── direct-go/              # Direct4B Go SDK
│   ├── client.go           # WebSocket client with MessagePack RPC
│   ├── client_test.go      # Unit tests for client
│   ├── auth.go             # Authentication (.env-based)
│   ├── messages.go         # Message sending functions
│   ├── events.go           # Event handling
│   ├── users.go            # User management API
│   ├── domains.go          # Domain/organization API
│   ├── talks.go            # Talk/room management API
│   ├── message_operations.go  # Message operations (search, favorites, reactions)
│   ├── files.go            # File upload/download API
│   ├── departments.go      # Department hierarchy API
│   ├── announcements.go    # Announcements API
│   ├── conference.go       # Video/audio conference API
│   ├── debuglog/           # Debug logging to separate server
│   ├── testutil/           # Test utilities and mock server
│   ├── tools/coverage/     # Porting coverage analysis tool
│   ├── direct-js-source/   # Synced JS source for reference
│   └── examples/
└── daab-go/                # Bot framework CLI
    ├── cmd/daabgo/         # Main CLI entry point
    ├── cmd/debugserver/    # Debug log server
    ├── internal/cli/       # CLI commands (cobra-based)
    │   ├── root.go         # Root command
    │   ├── init.go         # Initialize bot project
    │   ├── login.go        # Login to Direct4B
    │   ├── logout.go       # Logout
    │   ├── run.go          # Run bot (foreground)
    │   ├── start.go        # Start bot as daemon
    │   ├── stop.go         # Stop daemon
    │   ├── invites.go      # Manage domain invites
    │   ├── daemon.go       # Daemon management utilities
    │   └── version.go      # Show version
    ├── internal/bot/       # Bot framework (Hubot-like)
    ├── internal/webhook/   # n8n webhook integration
    │   ├── client.go       # HTTP webhook client
    │   ├── types.go        # Webhook payload/response types
    │   └── webhook_test.go # Webhook tests
    ├── daab-source/        # Synced daab JS source for reference
    └── examples/
        ├── ping/           # Simple ping bot
        └── n8n-proxy/      # n8n webhook proxy bot
```

## Development Workflow

### Working with direct-go

```bash
cd direct-go

# Run tests
go test ./...
go test -v              # Verbose output
go test -cover          # With coverage report
go test -race           # With race detector

# Run example
cd examples/simple
go run main.go

# Build and run coverage tool
cd tools/coverage
go run . -format markdown -output ../../COVERAGE.md
go run . -use-baseline -format text  # Quick text summary
```

### Working with daab-go

```bash
cd daab-go

# Build CLI
go build -o daabgo cmd/daabgo/main.go

# Run CLI commands
./daabgo init      # Initialize bot project
./daabgo login     # Login to Direct4B
./daabgo invites   # Show and accept domain invites
./daabgo run       # Run bot (foreground)
./daabgo start     # Start bot as daemon
./daabgo stop      # Stop daemon
./daabgo logout    # Logout
./daabgo version   # Show version

# Run example bot
cd examples/ping
go run main.go

# Run n8n webhook proxy example
cd examples/n8n-proxy
# Set up .env with DIRECT_ACCESS_TOKEN and N8N_WEBHOOK_URL
go run main.go

# Run debug server (for development)
cd cmd/debugserver
go run main.go
```

### Module Dependencies

daab-go depends on direct-go using a local replace directive in `daab-go/go.mod`:

```go
replace github.com/f4ah6o/direct-go-sdk/direct-go => ../direct-go
```

When modifying direct-go, changes are immediately visible to daab-go.

## Porting from JavaScript

### Source Synchronization

GitHub Actions workflows automatically sync upstream JavaScript sources:

* `.github/workflows/sync-direct-js.yaml`: Syncs and deminifies `direct-node.min.js`
* `.github/workflows/sync-daab.yaml`: Syncs daab library files

These workflows are **manually triggered** only via GitHub Actions UI.

### Tracking Progress

The **coverage tool** (`direct-go/tools/coverage/`) tracks porting progress by comparing RPC method calls:

* JavaScript baseline: 82 RPC methods across 13 categories
* Current Go implementation: ~88% coverage (72/82 methods)
* Generates detailed reports in JSON/Markdown/Text formats

Run coverage analysis:

```bash
cd direct-go/tools/coverage
go run . -format markdown > ../../COVERAGE.md
```

View `direct-go/COVERAGE.md` for current status and missing methods.

### Implemented RPC Methods (direct-go)

72 out of 82 methods implemented (~88% coverage):

**Session & Auth (6/7)**
* `create_session`, `start_notification`, `reset_notification`, `update_last_used_at`
* `create_access_token`, `create_access_token_by_id`
* Missing: account control request methods

**User Management (10/11)**
* `get_me`, `get_users`, `get_profile`, `update_profile`, `update_user`
* `get_presences`, `get_user_identifiers`
* `get_friends`, `add_friend`, `delete_friend`, `get_acquaintances`

**Domain Management (7/7)** ✅
* `get_domains`, `get_domain_invites`, `accept_domain_invite`, `delete_domain_invite`
* `leave_domain`, `get_domain_users`, `search_domain_users`

**Department Management (3/3)** ✅
* `get_department_tree`, `get_department_users`, `get_department_user_count`

**Talk/Room Management (8/9)**
* `get_talks`, `get_talk_statuses`, `create_group_talk`, `create_pair_talk`
* `update_group_talk`, `add_talkers`, `delete_talker`
* `add_favorite_talk`, `delete_favorite_talk`

**Message Operations (15/17)**
* `create_message`, `get_messages`, `delete_message`, `search_messages`, `search_messages_around_datetime`
* `get_favorite_messages`, `add_favorite_message`, `delete_favorite_message`
* `get_scheduled_messages`, `schedule_message`, `delete_scheduled_message`, `reschedule_message`
* `get_available_message_reactions`, `set_message_reaction`, `reset_message_reaction`, `get_message_reaction_users`
* Missing: `get_read_status`, `update_read_status`

**File & Attachment Management (6/6)** ✅
* `create_upload_auth`, `get_attachments`, `delete_attachment`, `search_attachments`
* `create_file_preview`, `get_file_preview`

**Announcement Management (4/4)** ✅
* `create_announcement`, `get_announcements`
* `get_announcement_statuses`, `update_announcement_status`

**Push Notification Management (2/2)** ✅
* `enable_push_notification`, `disable_push_notification`

**Conference/Call Management (5/5)** ✅
* `get_conferences`, `get_conference_participants`
* `join_conference`, `leave_conference`, `reject_conference`

**Miscellaneous (2/5)**
* `authorize_device`
* Missing: note management (4 methods)

## Key Architecture Patterns

### MessagePack RPC Protocol

direct-go implements the MessagePack RPC wire protocol:

```
Request:  [0, msgID, "method_name", [arg1, arg2, ...]]
Response: [1, msgID, error, result]
```

* `client.go`: WebSocket connection, RPC request/response handling
* `client.Call()`: Low-level RPC method (blocking)
* `client.XxxWithContext()`: Context-aware API methods (recommended)
* Helper methods wrap `Call()` for type safety

**Context Support**: Most API methods have `WithContext` variants that accept `context.Context` for cancellation and timeout control.

Example:
```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
user, err := client.GetMeWithContext(ctx)
```

### Event System

Events from server are dispatched via registered handlers:

```go
client.OnMessage(func(msg Message) { ... })
client.On("event_type", func(data interface{}) { ... })
```

Event types defined in `events.go`.

### Bot Framework (daab-go)

Hubot-inspired API with pattern matching:

```go
robot := bot.New()
robot.Hear("pattern", handler)      // Match any message
robot.Respond("pattern", handler)   // Match @bot mentions
robot.Run()
```

* `internal/bot/bot.go`: Core framework
* `internal/cli/`: CLI commands using cobra
  * Support for foreground (`run`) and daemon mode (`start`/`stop`)
  * Domain invite management (`invites`)
  * PID and log file management in `~/.daabgo/`
* Credentials stored in `.env` file (handled by `direct-go/auth.go`)

**Daemon Mode**: Bot can run as background daemon with PID tracking:
* PID file: `~/.daabgo/daabgo.pid`
* Log file: `~/.daabgo/daabgo.log`

### n8n Webhook Integration (daab-go)

The `internal/webhook` package provides n8n webhook integration for forwarding events to n8n workflows:

```go
import "github.com/f4ah6o/direct-go-sdk/daab-go/internal/webhook"

// Create webhook client
client := webhook.NewClient("https://your-n8n.com/webhook/xxx", "botname")

// Create and send payload
payload := webhook.NewPayload("message_created", "botname", messageData)
response, err := client.Send(payload)

// Validate and handle response
if errCode := response.Validate(); errCode != webhook.ErrorCodeOK {
    // Handle error
}
```

**Webhook Actions**:
* `none`: No action
* `reply`: Reply to message
* `send`: Send message to specific room
* `send_select`, `send_yesno`, `send_task`: Interactive message types
* `reply_select`, `reply_yesno`, `reply_task`: Reply to interactive messages
* `close_select`, `close_yesno`: Close interactive messages

See `daab-go/examples/n8n-proxy/` for a complete example.

### Message Domain Resolution (direct-go)

The SDK automatically resolves domain IDs for incoming messages:

* Talk-to-domain mapping is cached during `start_notification`
* `ReceivedMessage` includes `DomainID` field for domain-scoped operations
* Enables user lookups with proper domain context

### Debug Logging

Both modules support debug logging to a separate HTTP server:

```go
direct.EnableDebugServer("http://localhost:3939")
```

Server implementation: `daab-go/cmd/debugserver/`

## Common Commands

### Building

```bash
# Build daabgo CLI
cd daab-go
go build -o daabgo cmd/daabgo/main.go

# Install globally
go install github.com/f4ah6o/daabgo/cmd/daabgo@latest
```

### Testing

direct-go has unit tests with mock server support:

```bash
# Run all tests in workspace
go test ./...

# Run tests for specific module
cd direct-go && go test ./...
cd direct-go && go test -v -cover  # With coverage

# Run tests with race detector
cd direct-go && go test -race

# daab-go doesn't have tests yet
cd daab-go && go test ./...
```

**Test utilities**: `direct-go/testutil/` provides a mock WebSocket server for testing RPC calls.

### Linting

No specific linter configuration exists yet. Standard Go tools:

```bash
go vet ./...
go fmt ./...
```

## Important Notes

### Module Paths

* Published module path: `github.com/f4ah6o/direct-go-sdk/{direct-go,daab-go}`
* Import direct-go in external code: `import direct "github.com/f4ah6o/direct-go"`
* Import daab-go bot: `import "github.com/f4ah6o/daabgo/bot"`
* Test utilities: `import "github.com/f4ah6o/direct-go-sdk/direct-go/testutil"`

### JavaScript Reference Sources

* `direct-go/direct-js-source/direct-node.js`: Deminified direct-js (read-only reference)
* `daab-go/daab-source/lib/*.js`: daab source files (read-only reference)

**Do not modify** these directories; they are managed by GitHub Actions.

### Coverage Status

Current implementation status by category:

1. ✅ Domain Management (7/7) - 100%
2. ✅ Department Management (3/3) - 100%
3. ✅ File & Attachment Management (6/6) - 100%
4. ✅ Announcement Management (4/4) - 100%
5. ✅ Push Notification Management (2/2) - 100%
6. ✅ Conference/Call Management (5/5) - 100%
7. 🟡 User Management (10/11) - 91%
8. 🟡 Talk/Room Management (8/9) - 89%
9. 🟡 Message Operations (15/17) - 88%
10. 🟡 Session & Auth (6/7) - 86%
11. 🔴 Note Management (0/6) - 0%
12. 🔴 Miscellaneous (1/5) - 20%

**Missing Methods (10/82)**:
* Note management: `create_note`, `get_notes`, `update_note`, `delete_note`, `get_note_comments`, `create_note_comment`
* Session: account control request methods (3)
* Message: `get_read_status`, `update_read_status`

## API Compatibility

* API Version: `1.128` (defined in `client.go`)
* Default Endpoint: `wss://api.direct4b.com/albero-app-server/api`
* Authentication: OAuth access token via `.env` file or `Options.AccessToken`
