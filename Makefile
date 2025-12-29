.PHONY: test test-direct-go test-daab-go lint lint-direct-go lint-daab-go build clean help fmt vet

# Default target
all: test

# Run all tests
test:
	@echo "Running all tests..."
	@$(MAKE) test-direct-go
	@$(MAKE) test-daab-go

# Run tests for direct-go
test-direct-go:
	@echo "Testing direct-go..."
	@cd direct-go && go test -v -race -cover ./...

# Run tests for daab-go
test-daab-go:
	@echo "Testing daab-go..."
	@cd daab-go && go test -v -race -cover ./...

# Run linters
lint:
	@echo "Running linters..."
	@$(MAKE) lint-direct-go
	@$(MAKE) lint-daab-go

# Lint direct-go
lint-direct-go:
	@echo "Linting direct-go..."
	@cd direct-go && go vet ./...
	@if [ -n "$$(gofmt -l direct-go)" ]; then \
		echo "The following files need formatting:"; \
		gofmt -l direct-go; \
		exit 1; \
	fi

# Lint daab-go
lint-daab-go:
	@echo "Linting daab-go..."
	@cd daab-go && go vet ./...
	@if [ -n "$$(gofmt -l daab-go)" ]; then \
		echo "The following files need formatting:"; \
		gofmt -l daab-go; \
		exit 1; \
	fi

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./direct-go/...
	@go fmt ./daab-go/...

# Run go vet
vet:
	@echo "Running go vet..."
	@cd direct-go && go vet ./...
	@cd daab-go && go vet ./...

# Build daabgo CLI
build:
	@echo "Building daabgo..."
	@cd daab-go && go build -o bin/daabgo cmd/daabgo/main.go

# Build for multiple platforms
build-all:
	@echo "Building daabgo for multiple platforms..."
	@mkdir -p daab-go/bin
	@echo "  -> linux/amd64"
	@cd daab-go && GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/daabgo-linux-amd64 cmd/daabgo/main.go
	@echo "  -> linux/arm64"
	@cd daab-go && GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o bin/daabgo-linux-arm64 cmd/daabgo/main.go
	@echo "  -> darwin/amd64"
	@cd daab-go && GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/daabgo-darwin-amd64 cmd/daabgo/main.go
	@echo "  -> darwin/arm64"
	@cd daab-go && GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o bin/daabgo-darwin-arm64 cmd/daabgo/main.go
	@echo "  -> windows/amd64"
	@cd daab-go && GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/daabgo-windows-amd64.exe cmd/daabgo/main.go
	@echo "Done!"

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	@cd direct-go && go test -bench=. -benchmem ./...
	@cd daab-go && go test -bench=. -benchmem ./...

# Run go mod tidy
tidy:
	@echo "Tidying go.mod files..."
	@cd direct-go && go mod tidy
	@cd daab-go && go mod tidy

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf daab-go/bin
	@cd direct-go && go clean ./...
	@cd daab-go && go clean ./...

# Run security scan
security:
	@echo "Running security scan..."
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@cd direct-go && govulncheck ./...
	@cd daab-go && govulncheck ./...

# Install daabgo
install:
	@echo "Installing daabgo..."
	@cd daab-go && go install ./cmd/daabgo

# Show help
help:
	@echo "Available targets:"
	@echo "  all            - Run all tests (default)"
	@echo "  test           - Run all tests"
	@echo "  test-direct-go - Run direct-go tests"
	@echo "  test-daab-go   - Run daab-go tests"
	@echo "  lint           - Run all linters"
	@echo "  lint-direct-go - Lint direct-go"
	@echo "  lint-daab-go   - Lint daab-go"
	@echo "  fmt            - Format code"
	@echo "  vet            - Run go vet"
	@echo "  build          - Build daabgo CLI"
	@echo "  build-all      - Build daabgo for all platforms"
	@echo "  bench          - Run benchmarks"
	@echo "  tidy           - Run go mod tidy"
	@echo "  clean          - Clean build artifacts"
	@echo "  security       - Run security scan"
	@echo "  install        - Install daabgo to GOPATH/bin"
	@echo "  help           - Show this help message"
