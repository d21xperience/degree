# Build stage (tetap sama)
FROM golang:1.23.10-alpine3.22 AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /sc-service

# Runtime stage
FROM debian:stable-slim
WORKDIR /app

# Install solc dengan error handling
RUN set -eux; \
    apt-get update || (rm -rf /var/lib/apt/lists/* && apt-get update); \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        wget; \
    wget -O /usr/bin/solc https://binaries.soliditylang.org/linux-amd64/solc-linux-amd64-v0.8.30+commit.6d205f9d; \
    chmod +x /usr/bin/solc; \
    apt-get purge -y wget; \
    apt-get autoremove -y; \
    apt-get clean; \
    rm -rf /var/lib/apt/lists/*; \
    solc --version

COPY --from=builder /sc-service /sc-service
COPY .env .
EXPOSE 8184
EXPOSE 50054
ENTRYPOINT ["/sc-service"]