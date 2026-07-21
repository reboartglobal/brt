// wasi-go.mod - Module Go untuk WebAssembly System Interface
module github.com/username/rgn-website

go 1.21

require (
    github.com/gohugoio/hugo v0.125.0
)

// Untuk WASM/WASI build
// go:build wasip1 || wasm

// Dependencies untuk WASM
require (
    github.com/bytecodealliance/wasm-tools-go v0.1.0
    github.com/wasmerio/wasmer-go v1.0.4
)