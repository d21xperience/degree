# Multi-stage build untuk production
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build optimized binary - GUNAKAN NAMA YANG SESUAI
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-w -s' -o auth-service .

# Production stage
FROM alpine:3.19

WORKDIR /root/

# Install runtime dependencies
RUN apk --no-cache add ca-certificates

# Copy binary from builder
COPY --from=builder /app/auth-service .

# Expose ports SESUAI DOCKER-COMPOSE
#gRPC port
EXPOSE 50051  
#HTTP port
EXPOSE 8182   

# CMD instruction
CMD ["./auth-service"]