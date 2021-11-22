//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

func main() {
	defer func() { <-make(chan bool) }()

	println("Hello!")

	glob := js.Global()
	doc := glob.Get("document")

	cb := func() {
		glob.Get("myButton").Call("addEventListener", "click", js.FuncOf(func(_ js.Value, _ []js.Value) interface{} {
			glob.Call("alert", "clicked!")
			go func() {
				println("go async")
			}()
			return nil
		}))
	}

	if doc.Get("readyState").String() != "complete" {
		doc.Call("addEventListener", "DOMContentLoaded", js.FuncOf(func(_ js.Value, _ []js.Value) interface{} {
			cb()
			return nil
		}))
	} else {
		cb()
	}
}
