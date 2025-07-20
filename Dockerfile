# ---------- Build Stage ----------
FROM golang:1.20-alpine AS builder

# Set working directory inside the container
WORKDIR /build

# Install git and other dependencies (optional but common)
RUN apk add --no-cache git

# Copy go.mod and go.sum separately to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary (adjust path if needed)
RUN go build -o crm.shopdev.com ./cmd/server

# ---------- Run Stage ----------
FROM scratch

# Copy config files (adjust path if needed)
COPY --from=builder /build/config /config

# Copy built binary from the builder stage
COPY --from=builder /build/crm.shopdev.com /

# Command to run your binary
ENTRYPOINT ["/crm.shopdev.com", "config/local.yaml"]
