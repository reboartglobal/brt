#!/bin/bash

# deploy-wasmer.sh - Deploy script untuk Wasmer

set -e

echo "🚀 Deploying RGN Website to Wasmer..."

# Check if wasmer is installed
if ! command -v wasmer &> /dev/null; then
    echo "❌ Wasmer not found. Please install Wasmer first:"
    echo "  curl https://get.wasmer.io -sSfL | sh"
    exit 1
fi

# Check if logged in
if ! wasmer whoami &> /dev/null; then
    echo "❌ Not logged in to Wasmer. Please login:"
    echo "  wasmer login"
    exit 1
fi

# Build
echo "📦 Building..."
bash build-wasmer.sh

# Deploy
echo "🚀 Deploying to Wasmer..."
cd target
wasmer deploy --package rgn-website.wasmer.tar.gz

echo "✅ Deployment successful!"
echo "🌐 Your site is live on Wasmer!"