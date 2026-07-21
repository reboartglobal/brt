#!/bin/bash

# build-wasmer.sh - Build script untuk Wasmer deployment

set -e

echo "🚀 Building RGN Website for Wasmer..."

# 1. Build Hugo
echo "📦 Building Hugo site..."
hugo --gc --minify

# 2. Build Go to WebAssembly
echo "🔨 Building Go to WebAssembly..."
GOOS=wasip1 GOARCH=wasm CGO_ENABLED=0 go build -o target/wasm32-wasi/release/rgn-server.wasm main_wasm.go

# 3. Copy static files ke target
echo "📁 Copying static files..."
mkdir -p target/static
cp -r public/* target/static/

# 4. Copy config
echo "📋 Copying Wasmer config..."
cp wasmer.toml target/

# 5. Create package for Wasmer
echo "📦 Creating Wasmer package..."
cd target
tar -czf rgn-website.wasmer.tar.gz .

echo "✅ Build complete! Package created: target/rgn-website.wasmer.tar.gz"
echo ""
echo "📝 To deploy to Wasmer:"
echo "  wasmer deploy"
echo ""
echo "📝 To run locally:"
echo "  wasmer run target/wasm32-wasi/release/rgn-server.wasm --net --env PORT=8080"