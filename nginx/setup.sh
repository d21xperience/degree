#!/bin/bash

# Nama network yang digunakan
NETWORK_NAME="app-network"
COMPOSE_FILE="docker-compose.yml"

echo "📡 Memeriksa Docker network: $NETWORK_NAME..."

# Cek apakah network sudah ada
if docker network ls | grep -q "$NETWORK_NAME"; then
  echo "✅ Network '$NETWORK_NAME' sudah tersedia."
else
  echo "⚙️  Network '$NETWORK_NAME' belum tersedia. Membuat..."
  docker network create "$NETWORK_NAME"
  echo "✅ Network '$NETWORK_NAME' berhasil dibuat."
fi

echo "🔨 Membangun image dan menjalankan docker-compose..."
docker-compose -f "$COMPOSE_FILE" up --build -d

echo ""
echo "📦 Container yang sedang berjalan:"
docker ps --filter "name=nginx_gateway"

echo ""
echo "🌐 Status network '$NETWORK_NAME':"
docker network inspect "$NETWORK_NAME" | grep Name

echo ""
echo "🚀 Setup selesai. Kamu bisa akses:"
echo "   🔹 http://localhost"
echo "   🔹 https://localhost"
echo "   🔹 https://localhost/static/"
