#!/bin/bash

# SessionStart hook for direct-go-sdk repository
# This script runs automatically when a Claude Code web session starts

set -e

echo "=== Direct-Go-SDK Session Setup ==="
echo ""

# Check Go installation
echo "Checking Go installation..."
if ! command -v go &> /dev/null; then
    echo "❌ Error: Go is not installed"
    echo "   Please install Go 1.24 or later from https://go.dev/dl/"
    exit 1
fi

GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo "✓ Go version: $GO_VERSION"

# Check if Go version is 1.24 or later
MAJOR=$(echo "$GO_VERSION" | cut -d. -f1)
MINOR=$(echo "$GO_VERSION" | cut -d. -f2)
if [ "$MAJOR" -lt 1 ] || ([ "$MAJOR" -eq 1 ] && [ "$MINOR" -lt 24 ]); then
    echo "⚠️  Warning: Go 1.24+ recommended (current: $GO_VERSION)"
fi

echo ""

# Verify repository structure
echo "Verifying repository structure..."
REQUIRED_DIRS=("direct-go" "daab-go" "daab-go-examples")
for dir in "${REQUIRED_DIRS[@]}"; do
    if [ -d "$dir" ]; then
        echo "✓ Found: $dir/"
    else
        echo "❌ Missing: $dir/"
        exit 1
    fi
done

echo ""

# Check module files
echo "Checking Go modules..."
if [ -f "direct-go/go.mod" ]; then
    echo "✓ direct-go/go.mod exists"
else
    echo "❌ Missing: direct-go/go.mod"
    exit 1
fi

if [ -f "daab-go/go.mod" ]; then
    echo "✓ daab-go/go.mod exists"

    # Verify local replace directive
    if grep -q "replace github.com/f4ah6o/direct-go-sdk/direct-go => ../direct-go" "daab-go/go.mod"; then
        echo "✓ daab-go uses local direct-go (replace directive found)"
    else
        echo "⚠️  Warning: daab-go may not be using local direct-go"
    fi
else
    echo "❌ Missing: daab-go/go.mod"
    exit 1
fi

echo ""

# Download dependencies (silent mode)
echo "Downloading module dependencies..."
(cd direct-go && go mod download) &> /dev/null && echo "✓ direct-go dependencies ready" || echo "⚠️  direct-go dependency download had issues"
(cd daab-go && go mod download) &> /dev/null && echo "✓ daab-go dependencies ready" || echo "⚠️  daab-go dependency download had issues"

echo ""

# Verify dependencies
echo "Verifying module dependencies..."
(cd direct-go && go mod verify) &> /dev/null && echo "✓ direct-go modules verified" || echo "⚠️  direct-go module verification failed"
(cd daab-go && go mod verify) &> /dev/null && echo "✓ daab-go modules verified" || echo "⚠️  daab-go module verification failed"

echo ""

# Check for important documentation
echo "Documentation files:"
if [ -f "AGENTS.md" ]; then
    echo "✓ AGENTS.md (primary documentation)"
else
    echo "⚠️  Missing: AGENTS.md"
fi

if [ -f "CLAUDE.md" ]; then
    echo "✓ CLAUDE.md (quick reference)"
else
    echo "⚠️  Missing: CLAUDE.md"
fi

if [ -f "direct-go/COVERAGE.md" ]; then
    echo "✓ COVERAGE.md (porting progress)"
else
    echo "ℹ️  COVERAGE.md not found (run coverage tool to generate)"
fi

echo ""

# Report current branch
echo "Git branch: $(git branch --show-current 2>/dev/null || echo 'unknown')"
echo ""

# Success summary
echo "=== Setup Complete ==="
echo "✓ Environment ready for development"
echo ""
echo "Quick tips:"
echo "  • Read AGENTS.md for complete documentation"
echo "  • Read CLAUDE.md for quick reference"
echo "  • Run 'cd direct-go && go test ./...' to test SDK"
echo "  • Run 'cd daab-go && go test ./...' to test framework"
echo "  • Run 'cd direct-go/tools/coverage && go run .' for porting status"
echo ""
