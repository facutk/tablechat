# Build stage
FROM golang:1.25.2-alpine AS builder

WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate templ files (if needed)
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate

# Build the binary
RUN go build -o tablechat .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary and static files
COPY --from=builder /app/tablechat .
COPY --from=builder /app/static ./static

EXPOSE 3000

CMD ["./tablechat"]
