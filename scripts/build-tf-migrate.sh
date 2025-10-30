#!/bin/bash

set -e

# Install tf-migrate binary to the current directory
echo "Installing tf-migrate binary..."
GOBIN=$(pwd) go install github.com/cloudflare/tf-migrate/cmd/tf-migrate@latest

echo "Build complete! Binary available at: $(pwd)/tf-migrate"