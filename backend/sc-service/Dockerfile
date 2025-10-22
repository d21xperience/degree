# Build stage (tetap sama)
# FROM golang:1.23.10-alpine3.22 AS builder
FROM golang:1.23.4-alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /sc-service

# Runtime stage
FROM debian:stable-slim
WORKDIR /app

COPY solc-static-linux /usr/bin/solc
RUN chmod +x /usr/bin/solc

COPY --from=builder /sc-service /sc-service
COPY .env .
EXPOSE 8184
EXPOSE 50054
ENTRYPOINT ["/sc-service"]