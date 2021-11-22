package main

import "C"

import (
	"fmt"
)

//export HelloWorld
func HelloWorld(i int32, msg string) {
	fmt.Printf("Hello from go: i:%d msg:%s\n", i, msg)
}

//export HelloWorld2
func HelloWorld2(i int32, msg *C.char) {
	fmt.Printf("Hello from go: i:%d msg:%s\n", i, C.GoString(msg))
}

func main() {}
