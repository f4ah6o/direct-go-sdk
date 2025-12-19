# Claude Code Guide

This file provides quick reference for Claude Code sessions working with the direct-go-sdk repository.

## Primary Documentation

**See [AGENTS.md](AGENTS.md) for complete repository documentation**, including:
- Repository overview and structure
- Development workflows
- Module dependencies
- API patterns and examples
- Testing strategies
- Porting progress tracking

## Quick Start

This is a Go monorepo with two main modules:

1. **direct-go** - Go SDK for Direct4B WebSocket/MessagePack RPC API
2. **daab-go** - CLI tool and bot framework built on direct-go

### Prerequisites

- Go 1.24 or 1.25
- Git
- A Direct4B account with access token (for running examples)

### Environment Setup

The SessionStart hook automatically:
- Verifies Go installation and version
- Checks module dependencies
- Validates repository structure
- Reports any issues

### Common Development Commands

```bash
# Test direct-go SDK
cd direct-go && go test ./...

# Test daab-go framework
cd daab-go && go test ./...

# Build daabgo CLI
cd daab-go && go build -o daabgo cmd/daabgo/main.go

# Run coverage analysis
cd direct-go/tools/coverage && go run . -format markdown -output ../../COVERAGE.md

# Run example bot
cd daab-go-examples/ping && go run main.go
```

### Module Dependencies

daab-go depends on direct-go via local replace directive in `daab-go/go.mod`:
```go
replace github.com/f4ah6o/direct-go-sdk/direct-go => ../direct-go
```

Changes to direct-go are immediately visible to daab-go.

## Key Files

- `AGENTS.md` - Complete repository documentation (read this first!)
- `direct-go/COVERAGE.md` - Porting coverage status (88% complete)
- `direct-go/go.mod` - Direct SDK module definition
- `daab-go/go.mod` - Bot framework module definition
- `.github/workflows/ci.yaml` - CI/CD configuration

## Development Workflow

1. Read `AGENTS.md` for detailed context
2. Choose the module to work with (`direct-go` or `daab-go`)
3. Make changes in the appropriate module
4. Run tests: `go test ./...`
5. Update coverage report if adding RPC methods
6. Commit changes with clear messages

## Testing

```bash
# Run all tests
go test ./...

# Run with coverage
cd direct-go && go test -cover

# Run with race detection
cd direct-go && go test -race

# Verbose output
go test -v
```

Current test coverage: ~24% (18 tests across 3 files)

## Important Notes

- **Do not modify** `direct-go/direct-js-source/` or `daab-go/daab-source/` - these are managed by GitHub Actions
- API Version: `1.128`
- Default Endpoint: `wss://api.direct4b.com/albero-app-server/api`
- Porting progress: 72/82 RPC methods (~88%)

## Getting Help

- Read `AGENTS.md` for comprehensive documentation
- Check `direct-go/COVERAGE.md` for implementation status
- Review example projects in `daab-go-examples/`
- Examine test files for API usage examples

## Repository Status

This repository is actively maintained and being ported from JavaScript implementations:
- direct-go ports from [lisb/direct-js](https://github.com/lisb/direct-js)
- daab-go ports from [lisb/daab](https://github.com/lisb/daab)

See `AGENTS.md` for detailed porting progress and missing methods.
