package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"log"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed module/module.wasm
var moduleWasm []byte

func main() {
	var (
		ctx = context.Background()
		r   = wazero.NewRuntime(ctx)
	)
	defer r.Close(ctx)
	wasi_snapshot_preview1.MustInstantiate(ctx, r)
	var (
		stdOut bytes.Buffer
		stdErr bytes.Buffer
		config = wazero.NewModuleConfig().WithStdout(&stdOut).WithStderr(&stdErr)
	)
	mod, err := r.InstantiateWithConfig(ctx, moduleWasm, config)
	if err != nil {
		log.Panicf("could not instantiate module: %v", err)
	}
	defer mod.Close(ctx)
	fmt.Println("name", mod.Name())                                               // name
	fmt.Println("exportedFunctionDefinitions", mod.ExportedFunctionDefinitions()) // exportedFunctionDefinitions map[_start:0x140004f4280]
	fmt.Println("stdOut", stdOut.String())                                        // stdOut Hello from Wasm (stdOut)
	fmt.Println("stdErr", stdErr.String())                                        // stdErr Hello from Wasm (stdErr)
}
