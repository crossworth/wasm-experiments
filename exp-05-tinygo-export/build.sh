#!/usr/bin/env sh

tinygo build -target wasm -no-debug -o module.wasm ./main.go
cp $(tinygo env TINYGOROOT)/targets/wasm_exec.js ./wasm_exec.js