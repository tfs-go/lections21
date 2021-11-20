package main

/*
#cgo LDFLAGS: -lm
#include <math.h>
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.pow(2, 5))
}
