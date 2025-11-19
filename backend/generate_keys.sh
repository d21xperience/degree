#!/usr/bin/env bash
set -e

# Create directories
mkdir -p auth_service/config
mkdir -p sekolah/config
mkdir -p sc-service/config

echo "Generating RSA 2048 keypair..."

# Generate private key (hanya di auth_service)
openssl genrsa -out auth_service/config/private.pem 2048

# Generate public key untuk auth_service
openssl rsa -in auth_service/config/private.pem -pubout -out auth_service/config/public.pem

# Copy public key ke service lainnya
cp auth_service/config/public.pem sekolah/config/public.pem
cp auth_service/config/public.pem sc-service/config/public.pem

echo "Done! Keys have been generated and distributed:"
echo "- Private key: auth_service/config/private.pem"
echo "- Public keys:"
echo "  • auth_service/config/public.pem"
echo "  • sekolah-service/config/public.pem"
echo "  • smartcontract-service/config/public.pem"