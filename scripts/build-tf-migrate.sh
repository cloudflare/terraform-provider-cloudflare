#!/bin/bash

set -e

# Build tf-migrate binary from the legacy_migrations branch
echo "Building tf-migrate binary from legacy_migrations branch..."

# Set up temporary directory for cloning
TMP_DIR=$(mktemp -d)
trap "rm -rf $TMP_DIR" EXIT

# Clone the repository
echo "Cloning tf-migrate repository..."
git clone --quiet https://github.com/cloudflare/tf-migrate.git "$TMP_DIR/tf-migrate"

# Navigate to the cloned directory
cd "$TMP_DIR/tf-migrate"

# Checkout the legacy_migrations branch
echo "Checking out legacy_migrations branch..."
git checkout legacy_migrations

# Build the binary
echo "Building tf-migrate binary..."
go build -o tf-migrate ./cmd/tf-migrate

# Copy the binary to the original working directory
OUTPUT_DIR="$OLDPWD"
cp tf-migrate "$OUTPUT_DIR/tf-migrate"

echo "Build complete! Binary available at: $OUTPUT_DIR/tf-migrate"