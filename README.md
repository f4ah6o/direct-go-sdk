# direct-go-sdk

[![CI](https://github.com/f4ah6o/direct-go-sdk/actions/workflows/ci.yaml/badge.svg)](https://github.com/f4ah6o/direct-go-sdk/actions/workflows/ci.yaml)
[![direct-go porting coverage](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/f4ah6o/direct-go-sdk/main/.github/badges/direct-go-porting-coverage.json)](./direct-go/COVERAGE.md)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

> [!IMPORTANT]
> これは**非公式**のSDKです。L is B社およびdirect公式チームとは関係ありません。

This is an **unofficial** SDK for the direct (direct4b.com) chat service. It provides a low-level SDK (`direct-go`) and a high-level bot framework (`daab-go`).

## Modules

- **[direct-go](./direct-go)**: Go SDK for Direct4B WebSocket/MessagePack RPC API
- **[daab-go](./daab-go)**: Bot framework and CLI tools built on direct-go
- **[direct-teams-bridge](./cmd/direct-teams-bridge)**: Multi-account direct4b ⇄ Microsoft Teams bridge

## Reference Repositories

This SDK is developed based on the following official repositories:

- [lisb/direct-js](https://github.com/lisb/direct-js) - Direct JavaScript SDK
- [lisb/daab](https://github.com/lisb/daab) - Direct as a Bot framework
- [lisb/hubot-direct](https://github.com/lisb/hubot-direct) - Hubot adapter for direct

## Installation

```bash
# For the SDK
go get github.com/f4ah6o/direct-go-sdk/direct-go

# For the bot framework
go get github.com/f4ah6o/direct-go-sdk/daab-go
```

## Quick Start

### Using the SDK directly

```go
package main

import (
    "context"
    "log"

    direct "github.com/f4ah6o/direct-go-sdk/direct-go"
    "github.com/f4ah6o/direct-go-sdk/direct-go/auth"
)

func main() {
    // Load token from .env file or environment
    a := auth.NewAuth()
    token, err := a.GetToken()
    if err != nil {
        log.Fatal(err)
    }

    // Create client
    client := direct.NewClient(direct.Options{
        Token:    token,
        Endpoint: direct.DefaultEndpoint,
    })

    // Connect
    ctx := context.Background()
    if err := client.Connect(); err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // Send a message
    err = client.SendText(roomID, "Hello, World!")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Using the Bot Framework

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/f4ah6o/direct-go-sdk/daab-go/bot"
)

func main() {
    // Create a new bot
    r := bot.New(
        bot.WithName("mybot"),
    )

    // Add a simple handler
    r.Hear("hello", func(ctx context.Context, res bot.Response) {
        res.Send("Hello there!")
    })

    // Add a response handler (only when directly addressed)
    r.Respond("ping", func(ctx context.Context, res bot.Response) {
        res.Reply("pong")
    })

    // Run the bot
    ctx := context.Background()
    if err := r.Run(ctx); err != nil {
        log.Fatal(err)
    }
}
```

### Using Middleware

```go
import (
    "log"
    "os"

    "github.com/f4ah6o/direct-go-sdk/daab-go/bot"
)

func main() {
    r := bot.New()

    // Add middleware for logging and recovery
    logger := log.New(os.Stdout, "[BOT] ", log.LstdFlags)
    r.Use(bot.LoggingMiddleware(logger))
    r.Use(bot.RecoveryMiddleware(logger))

    // Add rate limiting (1 request per 10 seconds per user)
    r.Use(bot.RateLimitMiddleware(10 * time.Second))

    // Add filter to only process messages from specific users
    r.Use(bot.FilterMiddleware(func(ctx context.Context, res bot.Response) bool {
        return res.UserID() == "allowed_user_id"
    }))

    // Your handlers...
    r.Hear("test", func(ctx context.Context, res bot.Response) {
        res.Send("Filtered and rate-limited response!")
    })

    r.Run(context.Background())
}
```

### Using the CLI

```bash
# Login to obtain access token
daabgo login

# Initialize a new bot project
daabgo init my-bot

# Run the bot
cd my-bot
daabgo run

# Show available commands
daabgo --help
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DIRECT_TOKEN` | Access token for authentication | - |
| `HUBOT_DIRECT_ENDPOINT` | WebSocket endpoint URL | `wss://api.direct4b.com/albero-app-server/api` |
| `HUBOT_DIRECT_PROXY_URL` | Proxy URL for connections | - |
| `HTTPS_PROXY` | HTTPS proxy URL (fallback) | - |
| `HTTP_PROXY` | HTTP proxy URL (fallback) | - |

## Documentation

- [AGENTS.md](./AGENTS.md) - Complete development guide
- [direct-go/COVERAGE.md](./direct-go/COVERAGE.md) - API porting status

## Examples

See [daab-go-examples](./daab-go-examples/) for complete example projects:

- `ping` - Simple echo bot
- `webhook` - Webhook server example

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Run tests: `go test ./...`
4. Run linter: `golangci-lint run`
5. Commit your changes
6. Push to the branch
7. Create a Pull Request

## Development

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detector
go test -race ./...

# Build the CLI
cd daab-go && go build -o daabgo cmd/daabgo/main.go
```

## License

[MIT License](./LICENSE)
