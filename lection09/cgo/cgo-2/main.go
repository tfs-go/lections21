package main

/*
#include <stdio.h>
#include <stdlib.h>

void hello(char* s) {
  printf("From C: %s\n", s);
}
*/
import "C"

import "unsafe"

func main() {
	cs := C.CString("Hello, Tinkoff edu!")
	defer C.free(unsafe.Pointer(cs))

	C.hello(cs)
}
