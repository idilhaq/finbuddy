# ---------- Build stage ----------
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build fully static binary for Linux AMD64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o finbuddy-api ./cmd/api

# ---------- Production runtime ----------
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/finbuddy-api .

EXPOSE 3000

ENTRYPOINT ["./finbuddy-api"]

# ---------- Development runtime ----------
FROM golang:1.24-alpine AS dev

WORKDIR /app

RUN apk add --no-cache git curl

# Install air hot reload tool (binary installed in GOPATH/bin)
RUN go install github.com/air-verse/air@latest


COPY . .

# Run air for live reload during development
CMD ["air"]
