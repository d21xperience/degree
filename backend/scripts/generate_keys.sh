#!/usr/bin/env bash
set -e

# Create directories
mkdir -p auth_service/keys

echo "Generating RSA 2048 keypair..."

# Generate private key (hanya di auth_service)
openssl genrsa -out auth_service/keys/private.pem 2048

# Generate public key untuk auth_service
openssl rsa -in auth_service/keys/private.pem -pubout -out auth_service/keys/public.pem

echo "Done! Keys have been generated and distributed:"
echo "- Private key: auth_service/keys/private.pem"
echo "- Public keys:"
echo "  • auth_service/keys/public.pem"