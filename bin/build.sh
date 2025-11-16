#!/bin/bash

# Make script exit on first error
set -e

echo "=== Electerm Sync Server Build Script ==="
echo

# Create bin directory if it doesn't exist
mkdir -p bin

# Determine OS and Architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Map architecture names
case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    aarch64|arm64)
        ARCH="arm64"
        ;;
    *)
        echo "Architecture $ARCH might not be supported"
        ;;
esac

# Set output name based on OS
if [ "$OS" = "darwin" ]; then
    OUTPUT="output/electerm-sync-server-go"
elif [ "$OS" = "linux" ]; then
    OUTPUT="output/electerm-sync-server-go"
else
    OUTPUT="output/electerm-sync-server-go"
fi

echo "Building for $OS/$ARCH..."
GIN_MODE=release GOOS=$OS GOARCH=$ARCH go build -o $OUTPUT src/main.go
echo "Build complete: $OUTPUT"

# Make the binary executable
chmod +x $OUTPUT

echo
echo "=== Build Success ==="
echo "To run the server:"
echo "1. Make sure you have configured your .env file"
echo "2. Run the following command:"
echo "   GIN_MODE=release $OUTPUT"
echo
echo "The server should start and show:"
echo "server running at http://HOST:PORT"
echo "(HOST and PORT are specified in your .env file)"