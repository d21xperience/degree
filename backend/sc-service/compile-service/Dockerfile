# ===== Stage 1: builder (dengan Go lengkap) =====
FROM golang:1.23.10-alpine3.22 AS builder

WORKDIR /app
COPY . .
EXPOSE 8089
# Tambahkan toolchain jika perlu
RUN apk add --no-cache

# Build binary
RUN go build -o server .

# ===== Stage 2: final container (hanya solc + binary) =====
FROM ethereum/solc:0.8.30-alpine

# Override entrypoint bawaan solc
ENTRYPOINT []
WORKDIR /app

# Copy binary dari stage builder
COPY --from=builder /app/server .

# Pastikan folder uploads tersedia
RUN mkdir -p /app/sources

# Jalankan binary
CMD ["./server"]

