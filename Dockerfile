# Build stage
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files
COPY direct-go/go.mod direct-go/go.sum direct-go/
COPY daab-go/go.mod daab-go/go.sum daab-go/

# Download dependencies
RUN cd direct-go && go mod download
RUN cd daab-go && go mod download

# Copy source code
COPY direct-go/ direct-go/
COPY daab-go/ daab-go/

# Build daabgo
RUN cd daab-go && CGO_ENABLED=0 go build -ldflags="-s -w" -o /usr/local/bin/daabgo ./cmd/daabgo

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy the binary from builder
COPY --from=builder /usr/local/bin/daabgo /usr/local/bin/daabgo

# Create a non-root user
RUN addgroup -g 1000 daabgo && \
    adduser -D -u 1000 -G daabgo daabgo
USER daabgo

# Set the entrypoint
ENTRYPOINT ["/usr/local/bin/daabgo"]

# Default command shows help
CMD ["--help"]
