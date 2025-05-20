# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go binary for Linux (static)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o finbuddy-api .

# Final stage: lightweight Alpine image
FROM alpine:latest

# Install certificates for HTTPS calls
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/finbuddy-api .

# Expose API port
EXPOSE 8080

# Run the Go binary
ENTRYPOINT ["./finbuddy-api"]

# Run the executable
CMD ["finbuddy"]

# Test stage
# Use the same base image as the builder stage
FROM golang:1.24.0 AS test

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test ./...