# STAGE 1: Builder
FROM golang:alpine as builder

WORKDIR /build

# Files are copied from ./backend/api/
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary
RUN go build -o main .

# STAGE 2: Runner
FROM alpine:latest

RUN adduser -S -D -H -h /app appuser
USER appuser

WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/main /app/
# Copy .env from the local directory
COPY .env /app/.env

EXPOSE 3000

CMD ["./main"]