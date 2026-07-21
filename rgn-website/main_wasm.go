// main_wasm.go - Go server yang dikompilasi ke WebAssembly
//go:build wasip1 || wasm

package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Serve static files from the embedded public directory
    fs := http.FileServer(http.Dir("./public"))
    
    http.Handle("/", fs)
    
    // Health check endpoint
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status":"ok","service":"rgn-website"}`))
    })

    // API endpoint example
    http.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{
            "name": "RGN Website",
            "version": "1.0.0",
            "tagline": "Unlimited Connection",
            "status": "running"
        }`))
    })

    fmt.Printf("🚀 RGN Website running on Wasmer at http://localhost:%s\n", port)
    fmt.Printf("📁 Serving static files from ./public\n")
    fmt.Printf("🔗 Health check: http://localhost:%s/health\n", port)
    fmt.Printf("🔗 API: http://localhost:%s/api/info\n", port)
    
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal("❌ Server failed:", err)
    }
}