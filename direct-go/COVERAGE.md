# Direct4B Porting Coverage Report

**Generated**: 2025-12-29

**Note**: This report shows the current state of the port. The SDK implements
the core MessagePack RPC protocol and provides a Client type with methods
for common operations. See the godoc documentation for the full API.

## Summary

The direct-go SDK provides a Go client for the Direct4B WebSocket API.
It implements the MessagePack RPC protocol used by direct-js.

## Implemented Features

### Core Client Functionality
- WebSocket connection management with automatic reconnection
- MessagePack request/response handling
- Event-based message receiving
- Debug logging support

### Auth & Session
- Token storage and retrieval from environment variables and .env files
- Session creation with access token authentication
- Environment variable loading

### Message Operations
- Send text messages to talks/rooms
- Send messages with custom types (text, stamp, action, etc.)
- Scheduled message operations
- Message reaction handling

### Data Retrieval
- Get current user information (get_me)
- Get domain information
- Get talks/rooms list
- Get messages from talks
- Get message stamps

### Conference Support
- Conference creation and management
- Conference join info parsing

## Usage

```go
import "github.com/f4ah6o/direct-go-sdk/direct-go"

client := direct.NewClient(direct.Options{
    AccessToken: "your-token",
})

if err := client.Connect(); err != nil {
    log.Fatal(err)
}

// Send a message
err := client.SendText(roomID, "Hello, world!")
```

## See Also

- [Package documentation](https://pkg.go.dev/github.com/f4ah6o/direct-go-sdk/direct-go)
- [daab-go bot framework](../daab-go/README.md)
