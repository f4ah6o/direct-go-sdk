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
│   ├── events.go           # Event handling (with SelectStamp support)
│   ├── users.go            # User management API
│   ├── domains.go          # Domain/organization API
│   ├── talks.go            # Talk/room management API
│   ├── message_operations.go  # Message operations (search, favorites, reactions)
│   ├── files.go            # File upload/download API
│   ├── departments.go      # Department hierarchy API
│   ├── announcements.go    # Announcements API
│   ├── conference.go       # Video/audio conference API
│   ├── debuglog/           # Debug logging infrastructure
│   ├── logserver/          # LLM-friendly log server module
│   ├── testutil/           # Test utilities and mock server
│   ├── tools/
│   │   ├── coverage/       # Porting coverage analysis tool
│   │   └── testcov/        # Runtime code coverage tool
│   ├── direct-js-source/   # Synced JS source for reference
│   └── examples/simple/    # Simple example
├── daab-go/                # Bot framework CLI
│   ├── cmd/daabgo/         # Main CLI entry point
│   ├── cmd/logserver/      # LLM-friendly log server with JSON API & SSE
│   ├── internal/cli/       # CLI commands (cobra-based)
│   │   ├── root.go         # Root command
│   │   ├── init.go         # Initialize bot project (generates high-level bot framework code)
│   │   ├── login.go        # Login to Direct4B
│   │   ├── logout.go       # Logout
│   │   ├── run.go          # Run bot (foreground, uses high-level framework)
│   │   ├── start.go        # Start bot as daemon
│   │   ├── stop.go         # Stop daemon
│   │   ├── invites.go      # Manage domain invites
│   │   ├── daemon.go       # Daemon management utilities
│   │   └── version.go      # Show version
│   ├── bot/                # Public bot framework API (Hubot-like)
│   ├── webhook/            # n8n webhook integration (public API)
│   │   ├── client.go       # HTTP webhook client
│   │   ├── types.go        # Webhook payload/response types
│   │   └── webhook_test.go # Webhook tests
│   └── daab-source/        # Synced daab JS source for reference
├── daab-go-examples/       # Bot examples (separate module)
│   ├── ping/               # Simple ping bot
│   ├── n8n-proxy/          # n8n webhook proxy bot
│   ├── selectstamp/        # SelectStamp interactive message example
│   └── teams-bridge/       # Teams message bridge bot
├── cmd/                    # Command binaries (root module: direct-teams-bridge)
│   ├── direct-mcp-server/  # OAuth-protected MCP server (read/send tools)
│   ├── direct-teams-bridge/# Multi-account direct4b ⇄ Teams bridge
│   └── direct-slack-compat/# Slack-compatible Web/Events API adapter
├── internal/               # Internal packages for the root module
│   ├── mcpserver/          # MCP server: JWT auth, JWKS, account authz, tools
│   ├── teams/              # Teams Bot Framework client
│   ├── bridge/, directworker/, codex/, codexbridge/, config/, model/, secrets/, store/
│   └── ...                 # Bridge worker plumbing
├── slackcompat/            # Slack-compat adapter package (root module)
├── bench/                  # Benchmarks (e.g. runtime/go-ping, its own module)
├── docs/                   # Module for compile-checked README snippets (gen/ is gitignored)
├── tools/
│   └── doccheck/           # Extracts ```go blocks from READMEs and builds them in docs/
├── Makefile                # test/vet/build/fmt/doccheck across all modules in ci-modules.json
└── .github/ci-modules.json # Manifest of every Go module CI validates (add new modules here)
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
cd ../daab-go-examples/ping
go run main.go

# Run n8n webhook proxy example
cd ../daab-go-examples/n8n-proxy
# Set up .env with HUBOT_DIRECT_TOKEN and N8N_WEBHOOK_URL
go run main.go

# Run SelectStamp example
cd ../daab-go-examples/selectstamp
go run main.go

# Run log server (for development)
cd cmd/logserver
go run main.go
```

### Module Dependencies

daab-go depends on direct-go using a local replace directive in `daab-go/go.mod`:

```text
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

The **coverage tool** (`direct-go/tools/coverage/`) tracks porting progress by comparing RPC method calls against the direct-js baseline and generates reports in JSON/Markdown/Text formats.

Run coverage analysis:

```bash
cd direct-go/tools/coverage
go run . -format markdown > ../../COVERAGE.md
```

`direct-go/COVERAGE.md` is the source of truth for current status, per-category
counts, and missing methods — do not hard-code its numbers into this file.

### Runtime Test Coverage

The **testcov tool** (`direct-go/tools/testcov/`) provides runtime code coverage analysis:

* Separate from porting coverage - analyzes actual test coverage
* PowerShell script for Windows: `run.ps1`
* Helps identify untested code paths

Run test coverage analysis:

```bash
cd direct-go/tools/testcov
./run.ps1  # Windows PowerShell
```

### Implemented RPC Methods (direct-go)

Current implementation status — per-category counts, implemented method
lists, and missing methods — is tracked in `direct-go/COVERAGE.md` and
regenerated by the coverage tool. Update that report when RPC methods are
added rather than duplicating numbers here.

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

**Supported Message Types** (`events.go`):
* Standard messages: Text (1), System (4), File (7), Sticker (11)
* Interactive messages: YesNo (13), Select (15), Task (17)
* Reply messages: YesNoReply (14), SelectReply (16), TaskDone (18)
* Closed states: YesNoClosed (19), SelectClosed (20), TaskClosed (21)

**Wire Types for Interactive Messages**:
* `WireTypeSelect` (502): Select menu question
* `WireTypeSelectReply` (503): User's select choice
* `WireTypeYesNo` (500): Yes/No question
* `WireTypeYesNoReply` (501): Yes/No response
* `WireTypeTask` (504): Task assignment
* `WireTypeTaskDone` (505): Task completion
* Closed states: 506 (YesNo), 507 (Select), 508 (Task)

### Bot Framework (daab-go)

Hubot-inspired API with pattern matching and context support:

```go
import "github.com/f4ah6o/direct-go-sdk/daab-go/bot"

// Create bot with functional options
robot := bot.New(
    bot.WithName("mybot"),
    bot.WithToken("optional-token"),
)

// Register handlers with context support
robot.Hear("pattern", func(ctx context.Context, res bot.Response) {
    // Handle any message matching pattern
})

robot.Respond("pattern", func(ctx context.Context, res bot.Response) {
    // Handle @bot mentions matching pattern
})

// Run with context for graceful shutdown
ctx := context.Background()
robot.Run(ctx)
```

**Key Features**:
* Context-aware handlers for cancellation and timeout control
* Functional options pattern for configuration
* Pattern matching with regex support
* Response helpers: `Send()`, `Reply()`, `SendSelect()`
* **Advanced API Methods**:
  * `Robot.SendText(roomID, text)`: Send text message to any room
  * `Robot.Call(method, params)`: Low-level RPC access for advanced operations
  * `Robot.On(event, handler)`: Lifecycle event handling
* **Event System**: Register handlers for robot lifecycle events
  * `EventConnected`: Emitted when robot connects to Direct
  * `EventReady`: Emitted when robot is ready to receive messages
  * `EventDisconnected`: Emitted when robot disconnects
* **SelectStamp Support**: Interactive select menus with `SendSelect(question, options)`
  * Automatically handles user responses with wire types 502/503
  * Returns message ID for tracking
  * Supports both single and multiple option selection
  * See `daab-go-examples/selectstamp/` for complete implementation

**Package Structure**:
* `bot/bot.go`: Public API (recommended)
* `internal/cli/`: CLI commands using cobra
  * `init.go`: Generates projects using high-level bot framework
  * `run.go`: Runs bot with high-level framework (not low-level client)
  * Support for foreground (`run`) and daemon mode (`start`/`stop`)
  * Domain invite management (`invites`)
  * PID and log file management in `~/.daabgo/`
* Credentials stored in `.env` file (handled by `direct-go/auth.go`)

**Daemon Mode**: Bot can run as background daemon with PID tracking:
* PID file: `~/.daabgo/daabgo.pid`
* Log file: `~/.daabgo/daabgo.log`

**Event System Example**:
```go
robot := bot.New(bot.WithName("mybot"))

// Register lifecycle event handlers
robot.On(bot.EventConnected, func() {
    log.Println("Bot connected!")
})

robot.On(bot.EventReady, func() {
    log.Println("Bot is ready to receive messages")
})

robot.On(bot.EventDisconnected, func() {
    log.Println("Bot disconnected")
})
```

**Advanced RPC Example**:
```go
// Use Call() for advanced operations not exposed by helper methods
result, err := robot.Call("get_messages", []interface{}{talkID, 0, 100})
if err != nil {
    log.Printf("RPC call failed: %v", err)
}
```

**SelectStamp Example**:
```go
// Interactive select menu
robot.Hear("menu", func(ctx context.Context, res bot.Response) {
    options := []string{"Option A", "Option B", "Option C"}
    messageID, err := res.SendSelect("Please choose:", options)
    if err != nil {
        log.Printf("Failed to send select: %v", err)
        return
    }
    log.Printf("Sent select menu with message ID: %s", messageID)
})

// Handle select responses (wire type 503) and other messages
robot.Hear(".*", func(ctx context.Context, res bot.Response) {
    // Check if this is a select response or regular message
    log.Printf("Received: %s", res.Text())
})
```

See `daab-go-examples/selectstamp/` for a complete working example with:
* Menu presentation with `SendSelect()` returning message ID
* Automatic handling of SelectReply messages (wire type 503)
* Support for msgpack integer types in responses
* UUID fortune telling and Mirasapo API integration
* Menu tracking to detect user responses

### n8n Webhook Integration (daab-go)

The `webhook` package provides n8n webhook integration for forwarding events to n8n workflows:

```go
import "github.com/f4ah6o/direct-go-sdk/daab-go/webhook"

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

**Package Availability**:
* Public API: `github.com/f4ah6o/direct-go-sdk/daab-go/webhook`
* Can be used independently in custom bot implementations

See `daab-go-examples/n8n-proxy/` for a complete example.

### Message Domain Resolution (direct-go)

The SDK automatically resolves domain IDs for incoming messages:

* Talk-to-domain mapping is cached during `start_notification`
* `ReceivedMessage` includes `DomainID` field for domain-scoped operations
* Enables user lookups with proper domain context

### Log Server (LLM-Friendly Debug Logging)

Both modules support structured debug logging to a separate HTTP server:

```go
direct.EnableDebugServer("http://localhost:3939")
```

**Log Server Features** (`daab-go/cmd/logserver/`):
* **JSON API**: Structured log entries with timestamps, levels, and metadata
* **SSE Streaming**: Real-time log streaming via Server-Sent Events
* **HTML UI**: Browser-based log viewer for easy debugging
* **LLM-Friendly**: Designed for analysis by Claude Code and other AI tools
* **Log Infrastructure**: `direct-go/logserver/` provides reusable server module

**Usage**:
```bash
# Start log server on port 3939
cd daab-go/cmd/logserver
go run main.go

# In your bot code, enable debug logging
direct.EnableDebugServer("http://localhost:3939")
```

**Endpoints**:
* `GET /` - HTML log viewer UI
* `GET /logs` - JSON array of all log entries
* `GET /stream` - SSE stream of real-time logs

## Common Commands

### Building

```bash
# Build daabgo CLI
cd daab-go
go build -o daabgo cmd/daabgo/main.go

# Install globally
go install github.com/f4ah6o/direct-go-sdk/daab-go/cmd/daabgo@latest
```

### Testing

Every command below validates **all** modules via `.github/ci-modules.json`
(the same manifest CI uses) — prefer the Makefile over `go test ./...`, which
only covers the module you run it in:

```bash
make test      # go test ./... in every module
make vet       # go vet ./... in every module
make build     # go build ./... in every module
make fmt       # gofmt check on tracked files
make doccheck  # compile-check the Go snippets in README files

# Single-module equivalents
cd direct-go && go test -race ./...
cd daab-go && go test -race ./...

# Run CI locally (if act is installed)
act -j test
```

direct-go has comprehensive unit tests with mock server support:

**Test utilities**: `direct-go/testutil/` provides:
- `MockServer`: WebSocket mock server for RPC testing
- `OnSimple()`: Simple mock responses
- `OnDynamic()`: Parameter-based dynamic responses
- `GetCallCount()`: Method call verification
- `Reset()`: Test isolation

**CI/CD**: GitHub Actions workflow (`.github/workflows/ci.yaml`)
- Every Go module is tested, built, vetted, formatted, and vulnerability-scanned; the
  module list lives in `.github/ci-modules.json` (the Module Inventory job fails if a
  `go.mod` is missing from it — add new modules there, with their Go versions)
- Test matrix: per-module Go versions from `ci-modules.json` (minimum + latest supported)
- Coverage reporting per module
- Race detection
- Lint checks (go vet, gofmt)
- Security scan: `govulncheck` pinned via `GOVULNCHECK_VERSION` in the workflow env
  (do not use `@latest`; v1.8.0+ requires Go >= 1.26)

### Linting

No specific linter configuration exists yet. Standard Go tools, run across all
modules via the Makefile:

```bash
make vet    # go vet ./... per module
make fmt    # gofmt -l on tracked files
```

CI also compile-checks every fenced Go snippet in the READMEs:

```bash
go run ./tools/doccheck   # extracts ```go blocks → docs/gen/ and builds them
```

## Important Notes

### Module Paths

* Published module path: `github.com/f4ah6o/direct-go-sdk/{direct-go,daab-go}`
* Import direct-go in external code: `import direct "github.com/f4ah6o/direct-go-sdk/direct-go"`
* Import daab-go bot: `import "github.com/f4ah6o/direct-go-sdk/daab-go/bot"`
* Import webhook integration: `import "github.com/f4ah6o/direct-go-sdk/daab-go/webhook"`
* Test utilities: `import "github.com/f4ah6o/direct-go-sdk/direct-go/testutil"`
* Log server module: `import "github.com/f4ah6o/direct-go-sdk/direct-go/logserver"`

### Example Projects

Bot examples are maintained in a separate module at the repository root:

* **Location**: `daab-go-examples/` (root-level directory, separate from daab-go)
* **Module**: Independent `go.mod` with dependencies on direct-go and daab-go
* **Examples**:
  * `ping/`: Simple ping-pong bot demonstrating basic message handling
  * `n8n-proxy/`: n8n webhook integration for workflow automation
  * `selectstamp/`: Interactive select menu example (SelectStamp feature)
  * `teams-bridge/`: Message bridge between Direct4B and Microsoft Teams

**Running examples**:
```bash
cd daab-go-examples/ping
go run main.go
```

### JavaScript Reference Sources

* `direct-go/direct-js-source/direct-node.js`: Deminified direct-js (read-only reference)
* `daab-go/daab-source/lib/*.js`: daab source files (read-only reference)

**Do not modify** these directories; they are managed by GitHub Actions.

### Coverage Status

Per-category implementation status and missing methods are generated into
`direct-go/COVERAGE.md` by the coverage tool (`direct-go/tools/coverage/`).
See that file — it is checked into the repo and refreshed by CI — rather than
maintaining numbers here.

## API Compatibility

* API Version: `1.128` (defined in `client.go`)
* Default Endpoint: `wss://api.direct4b.com/albero-app-server/api`
* Authentication: OAuth access token via `.env` file or `Options.AccessToken`
