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
│   ├── auth.go             # Authentication (.env-based)
│   ├── messages.go         # Message sending functions
│   ├── events.go           # Event handling
│   ├── debuglog/           # Debug logging to separate server
│   ├── tools/coverage/     # Porting coverage analysis tool
│   ├── direct-js-source/   # Synced JS source for reference
│   └── examples/
└── daab-go/                # Bot framework CLI
    ├── cmd/daabgo/         # Main CLI entry point
    ├── cmd/debugserver/    # Debug log server
    ├── internal/cli/       # CLI commands (cobra-based)
    ├── internal/bot/       # Bot framework (Hubot-like)
    ├── daab-source/        # Synced daab JS source for reference
    └── examples/
```

## Development Workflow

### Working with direct-go

```bash
cd direct-go

# Run tests (currently no test files)
go test ./...

# Run example
cd examples/simple
go run main.go

# Build and run coverage tool
cd tools/coverage
go run . -format markdown -output ../../COVERAGE.md
```

### Working with daab-go

```bash
cd daab-go

# Build CLI
go build -o daabgo cmd/daabgo/main.go

# Run CLI commands
./daabgo init    # Initialize bot project
./daabgo login   # Login to Direct4B
./daabgo run     # Run bot

# Run example bot
cd examples/ping
go run main.go

# Run debug server
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
* Current Go implementation: ~10% coverage (8 methods)
* Generates detailed reports in JSON/Markdown/Text formats

Run coverage analysis:

```bash
cd direct-go/tools/coverage
go run . -format markdown > ../../COVERAGE.md
```

View `direct-go/COVERAGE.md` for current status and missing methods.

### Implemented RPC Methods (direct-go)

* Session & Auth: `create_session`, `start_notification`, `reset_notification`, `update_last_used_at`
* User: `get_me`
* Domain: `get_domains`, `accept_domain_invite`
* Talk/Room: `create_talk`, `send_message`

## Key Architecture Patterns

### MessagePack RPC Protocol

direct-go implements the MessagePack RPC wire protocol:

```
Request:  [0, msgID, "method_name", [arg1, arg2, ...]]
Response: [1, msgID, error, result]
```

* `client.go`: WebSocket connection, RPC request/response handling
* `client.Call()`: Low-level RPC method
* Helper methods wrap `Call()` for type safety

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
* Credentials stored in `.env` file (handled by `direct-go/auth.go`)

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

Currently no test files exist. When adding tests:

```bash
# Run all tests in workspace
go test ./...

# Run tests for specific module
cd direct-go && go test ./...
cd daab-go && go test ./...
```

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

### JavaScript Reference Sources

* `direct-go/direct-js-source/direct-node.js`: Deminified direct-js (read-only reference)
* `daab-go/daab-source/lib/*.js`: daab source files (read-only reference)

**Do not modify** these directories; they are managed by GitHub Actions.

### Coverage Tool Categories

When implementing new RPC methods, check which category they belong to:

1. Session & Auth (7 methods)
2. User Management (11 methods)
3. Domain Management (7 methods)
4. Department Management (3 methods)
5. Talk/Room Management (9 methods)
6. Message Operations (17 methods)
7. File & Attachment Management (6 methods)
8. Note Management (6 methods)
9. Announcement Management (4 methods)
10. Push Notification Management (2 methods)
11. Conference/Call Management (5 methods)
12. Miscellaneous (5 methods)

Prioritize based on coverage gaps shown in `COVERAGE.md`.

## API Compatibility

* API Version: `1.128` (defined in `client.go`)
* Default Endpoint: `wss://api.direct4b.com/albero-app-server/api`
* Authentication: OAuth access token via `.env` file or `Options.AccessToken`
