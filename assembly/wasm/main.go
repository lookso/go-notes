package main

import (
	"syscall/js"
)
// GOOS=js GOARCH=wasm go build -o main.wasm
// GOOS=js GOARCH=wasm go build -o main.wasm wasm.go
func main() {
	js.Global().Set("printMessage", js.FuncOf(printMessage))
	select {}
}

func printMessage(this js.Value, inputs []js.Value) interface{} {
	message := inputs[0].String()
	js.Global().Get("console").Call("log", message)
	js.Global().Call("alert", message)
	return nil
}