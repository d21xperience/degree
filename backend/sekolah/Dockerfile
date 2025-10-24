# Base image
# FROM golang:1.22.3 AS builder
FROM golang:1.23.4-alpine
# Set environment variables
# ENV GO111MODULE=on
# ENV GOPROXY=https://proxy.golang.org,direct

# Set working directory
WORKDIR /app

# Copy go.mod, go.sum, dan dependensi lokal
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy seluruh source code
COPY . .

# Expose port gRPC
EXPOSE 50052

# # Jalankan aplikasi dengan Air untuk hot reload
# CMD ["air"]