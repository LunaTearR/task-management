# ====== Stage 1: Build ======
FROM golang:1.23.11-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go app
RUN go build -o main .

# ====== Stage 2: Run ======
FROM debian:bookworm-slim

WORKDIR /app

# Update and install security patches
RUN apt-get update && apt-get upgrade -y && rm -rf /var/lib/apt/lists/*

# Copy binary from builder
COPY --from=builder /app/main .
COPY config.yaml .
# Expose port (เช่น 3000 ตาม config.yaml)
EXPOSE 3001

# Run app
CMD ["./main", "start"]
