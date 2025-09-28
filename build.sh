#!/bin/bash

set -e

echo "Building webview SetUserAgent demo..."

if [[ "$(uname)" != "Darwin" ]]; then
    echo "This is macOS-only. For Linux/Windows, see README."
    exit 1
fi

if ! command -v go &> /dev/null; then
    echo "Go not found. Please install Go first."
    exit 1
fi

go build -o webview-useragent-demo cmd/main.go

echo "Done! Run with: ./webview-useragent-demo"