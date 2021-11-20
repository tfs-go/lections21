package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	println("Hello!")

	glob := js.Global

	glob.Call("addEventListener", "DOMContentLoaded", func() {
		glob.Get("myButton").Call("addEventListener", "click", func() {
			glob.Call("alert", "clicked!")
			go func() {
				println("go async")
			}()
		})
	})
}
