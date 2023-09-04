package main

import (
	"os"
)

func main() {
	_, _ = os.Stdout.WriteString("Hello from Wasm (stdOut)")
	_, _ = os.Stderr.WriteString("Hello from Wasm (stdErr)")
}
