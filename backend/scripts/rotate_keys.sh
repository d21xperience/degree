#!/usr/bin/env bash
set -euo pipefail

# Config
KEY_DIR="${1:-./keys}"         # default ./keys or pass custom dir as first arg
ARCHIVE_DIR="$KEY_DIR/archive"
TMP_DIR="$KEY_DIR/.tmp"
mkdir -p "$KEY_DIR" "$ARCHIVE_DIR" "$TMP_DIR"

# Helpers
base64url_encode() {
  # read stdin, output base64url (no padding)
  openssl base64 -A | tr '+/' '-_' | tr -d '='
}

hex_to_base64url() {
  # $1 = hex string (no spaces)
  printf "%s" "$1" | xxd -r -p | base64url_encode
}

# Generate timestamp
TS=$(date +"%Y%m%d_%H%M%S")

echo "[*] Generating keys in $KEY_DIR (timestamp: $TS)"

# Backup old keys if exist
if [[ -f "$KEY_DIR/private.pem" ]]; then
  echo "[*] Archiving old keys"
  mv "$KEY_DIR/private.pem" "$ARCHIVE_DIR/private_$TS.pem"
fi
if [[ -f "$KEY_DIR/public.pem" ]]; then
  mv "$KEY_DIR/public.pem" "$ARCHIVE_DIR/public_$TS.pem"
fi
if [[ -f "$KEY_DIR/jwks.json" ]]; then
  mv "$KEY_DIR/jwks.json" "$ARCHIVE_DIR/jwks_$TS.json"
fi

# 1) Generate RSA private key (2048 bits)
openssl genpkey -algorithm RSA -out "$KEY_DIR/private.pem" -pkeyopt rsa_keygen_bits:2048

# 2) Derive public key PEM
openssl rsa -in "$KEY_DIR/private.pem" -pubout -out "$KEY_DIR/public.pem"

# 3) Extract modulus hex (full) from public.pem
# Use openssl -text -noout and grab lines between "Modulus:" and "Exponent:"
MOD_HEX=$(openssl rsa -pubin -in "$KEY_DIR/public.pem" -text -noout 2>/dev/null | \
  awk '/Modulus:/{flag=1; next} /Exponent:/{flag=0} flag { gsub(/[: \t]/,""); printf "%s", $0 } END{print ""}')

if [[ -z "$MOD_HEX" ]]; then
  echo "ERROR: Failed to extract modulus hex from public.pem" >&2
  exit 1
fi

# 4) Extract exponent hex (from parentheses like 0x10001) or fallback to decimal conversion
EXP_PAREN=$(openssl rsa -pubin -in "$KEY_DIR/public.pem" -text -noout 2>/dev/null | awk -F'[()]' '/Exponent:/{print $2; exit}')
if [[ -n "$EXP_PAREN" && "$EXP_PAREN" =~ ^0x ]]; then
  # remove 0x prefix
  EXP_HEX="${EXP_PAREN#0x}"
else
  # fallback: parse decimal exponent and convert to hex
  EXP_DEC=$(openssl rsa -pubin -in "$KEY_DIR/public.pem" -text -noout 2>/dev/null | awk '/Exponent:/{print $2; exit}')
  if [[ -z "$EXP_DEC" ]]; then
    echo "ERROR: Failed to parse exponent" >&2
    exit 1
  fi
  # convert decimal to hex (portable using printf with bc if needed)
  # try bash/printf
  printf -v EXP_HEX "%x" "$EXP_DEC" || EXP_HEX=$(echo "obase=16; $EXP_DEC" | bc)
fi

# Normalize exponent hex to even length (xxd -r -p needs even length)
if (( ${#EXP_HEX} % 2 == 1 )); then
  EXP_HEX="0$EXP_HEX"
fi

# Ensure MOD_HEX length is even (should be)
if (( ${#MOD_HEX} % 2 == 1 )); then
  MOD_HEX="0$MOD_HEX"
fi

# 5) Convert to base64url
N_B64URL=$(hex_to_base64url "$MOD_HEX")
E_B64URL=$(hex_to_base64url "$EXP_HEX")

# 6) Compute kid as JWK Thumbprint per RFC7638:
# thumbprint input is the UTF-8 bytes of the JSON object with keys lexicographically ordered.
THUMBPRINT_JSON=$(printf '{"e":"%s","kty":"RSA","n":"%s"}' "$E_B64URL" "$N_B64URL")
KID=$(printf '%s' "$THUMBPRINT_JSON" | openssl dgst -sha256 -binary | base64url_encode)

# 7) Build JWKS JSON (atomic write)
JWKS_TMP="$TMP_DIR/jwks.$TS.json"
cat > "$JWKS_TMP" <<EOF
{
  "keys": [
    {
      "kid": "$KID",
      "kty": "RSA",
      "alg": "RS256",
      "use": "sig",
      "n": "$N_B64URL",
      "e": "$E_B64URL"
    }
  ]
}
EOF

# Atomic move
mv "$JWKS_TMP" "$KEY_DIR/jwks.json"
rm -rf "$TMP_DIR" || true

echo "[✓] Generated:"
echo "    - $KEY_DIR/private.pem"
echo "    - $KEY_DIR/public.pem"
echo "    - $KEY_DIR/jwks.json (kid: $KID)"
echo "[i] Archived previous keys (if any) under $ARCHIVE_DIR"
