#!/bin/bash

# vercel-deploy.sh - Deploy script untuk Vercel

set -e

echo "🚀 Deploying RGN Website to Vercel..."

# Check if Vercel CLI is installed
if ! command -v vercel &> /dev/null; then
    echo "❌ Vercel CLI not found. Installing..."
    npm install -g vercel
fi

# Check if logged in
if ! vercel whoami &> /dev/null; then
    echo "❌ Not logged in to Vercel. Please login:"
    echo "  vercel login"
    exit 1
fi

# Build
echo "📦 Building..."
npm run build

# Deploy to production
echo "🚀 Deploying to production..."
vercel --prod

echo "✅ Deployment successful!"
echo "🌐 Your site is live on Vercel!"