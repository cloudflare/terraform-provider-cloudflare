#!/bin/bash

set -e

# Clone the repository
if [ -d "tf-migrate" ]; then
    echo "tf-migrate directory already exists. Pulling latest changes..."
    cd tf-migrate
    git pull
else
    echo "Cloning tf-migrate repository..."
    git clone https://github.com/cloudflare/tf-migrate.git tf-migrate
    cd tf-migrate
fi

# Build the binary
echo "Building tf-migrate binary..."
go build -o tf-migrate

echo "Build complete! Binary available at: $(pwd)/tf-migrate"