#!/usr/bin/env sh

GOOS=js GOARCH=wasm go build -o module.wasm ./main.go
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./wasm_exec.js