package main

import "syscall/js"

func main() {
	println("Hello from GO")
	js.Global().Set("add", js.FuncOf(add))
	select {}
}

func add(this js.Value, inputs []js.Value) any {
	return inputs[0].Int() + inputs[1].Int()
}
