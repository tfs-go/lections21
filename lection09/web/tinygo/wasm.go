//go:build wasm
// +build wasm

/**
Будет ругаться на:
syscall/js.finalizeRef not implemented

Это еще не исправлено:
https://github.com/tinygo-org/tinygo/issues/1140
*/

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	defer func() { <-make(chan bool) }()

	println("Hello!")

	glob := js.Global()
	doc := glob.Get("document")

	cb := func() {
		glob.Get("myButton").Call("addEventListener", "click", js.FuncOf(func(_ js.Value, _ []js.Value) interface{} {
			fmt.Println("clicked")
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
