#!/bin/bash

# vercel-build.sh - Build script untuk Vercel

echo "🚀 Building RGN Website for Vercel..."

# Set environment
export HUGO_ENV=production

# Build Hugo
echo "📦 Building Hugo site..."
hugo --gc --minify

# Check if build was successful
if [ $? -eq 0 ]; then
    echo "✅ Hugo build successful!"
    echo "📁 Static files ready in ./public"
else
    echo "❌ Hugo build failed!"
    exit 1
fi

echo "🎉 Build complete!"