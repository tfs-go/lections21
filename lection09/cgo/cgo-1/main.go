package main

/*
int sum(int a, int b) {
    return a + b;
}
*/
import "C"

import (
	"fmt"
)

func main() {
	fmt.Println(C.sum(40, 2))
}
