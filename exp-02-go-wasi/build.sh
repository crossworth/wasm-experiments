#!/usr/bin/env sh

GOOS=wasip1 GOARCH=wasm go build -o module.wasm ./main.go