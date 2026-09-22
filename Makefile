# Validate every module listed in .github/ci-modules.json (the same manifest CI uses).
MODULES := $(shell jq -r '.modules[].dir' .github/ci-modules.json)

.PHONY: all test vet fmt fmt-fix build build-all doccheck bench tidy clean security install lint help

all: vet fmt doccheck test build

# Run all tests in every module
test:
	@for d in $(MODULES); do echo "== go test $$d =="; (cd $$d && go test ./...) || exit 1; done

# Run tests for direct-go only
test-direct-go:
	@cd direct-go && go test -v -race -cover ./...

# Run tests for daab-go only
test-daab-go:
	@cd daab-go && go test -v -race -cover ./...

# go vet across every module
vet:
	@for d in $(MODULES); do echo "== go vet $$d =="; (cd $$d && go vet ./...) || exit 1; done

# gofmt check on all tracked files (same check CI runs)
fmt:
	@unformatted=$$(gofmt -l $$(git ls-files '*.go')); \
	if [ -n "$$unformatted" ]; then \
		echo "The following files need formatting:"; \
		echo "$$unformatted"; \
		exit 1; \
	fi

# Format code in place
fmt-fix:
	@gofmt -w $$(git ls-files '*.go')

# go vet + gofmt check
lint: vet fmt

# Compile-check the Go snippets in README files (same check CI runs)
doccheck:
	go run ./tools/doccheck

# Build every module
build:
	@for d in $(MODULES); do echo "== go build $$d =="; (cd $$d && go build ./...) || exit 1; done

# Build daabgo CLI
build-daabgo:
	@cd daab-go && go build -o bin/daabgo cmd/daabgo/main.go

# Build daabgo for multiple platforms
build-all:
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
	@cd direct-go && go test -bench=. -benchmem ./...
	@cd daab-go && go test -bench=. -benchmem ./...

# Regenerate go.mod/go.sum for every module (e.g. after adding a module to the manifest)
tidy:
	@for d in $(MODULES); do echo "== go mod tidy $$d =="; (cd $$d && go mod tidy) || exit 1; done

# Clean build artifacts
clean:
	@rm -rf daab-go/bin docs/gen
	@cd direct-go && go clean ./...
	@cd daab-go && go clean ./...

# Run govulncheck (pinned to match CI; v1.8.0+ requires Go >= 1.26)
security:
	@go install golang.org/x/vuln/cmd/govulncheck@v1.7.0
	@for d in $(MODULES); do echo "== govulncheck $$d =="; (cd $$d && govulncheck ./...) || exit 1; done

# Install daabgo
install:
	@cd daab-go && go install ./cmd/daabgo

# Show help
help:
	@echo "Available targets:"
	@echo "  all            - Run vet + fmt + doccheck + test + build"
	@echo "  test           - Run go test ./... in every module"
	@echo "  test-direct-go - Run direct-go tests only"
	@echo "  test-daab-go   - Run daab-go tests only"
	@echo "  vet            - Run go vet ./... in every module"
	@echo "  fmt            - Check gofmt on tracked files (CI parity)"
	@echo "  fmt-fix        - Format tracked files in place"
	@echo "  lint           - Run vet + fmt checks"
	@echo "  doccheck       - Compile-check README Go snippets"
	@echo "  build          - Build every module"
	@echo "  build-daabgo   - Build daabgo CLI"
	@echo "  build-all      - Build daabgo for all platforms"
	@echo "  bench          - Run benchmarks"
	@echo "  tidy           - Run go mod tidy in every module"
	@echo "  clean          - Clean build artifacts"
	@echo "  security       - Run govulncheck (pinned) in every module"
	@echo "  install        - Install daabgo to GOPATH/bin"
	@echo "  help           - Show this help message"
