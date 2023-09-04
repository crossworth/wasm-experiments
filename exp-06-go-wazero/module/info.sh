#!/usr/bin/env sh

wasm2wat module.wasm -o module.wat
wasm-objdump module.wasm -x > details.txt