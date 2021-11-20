package main

/*
#include <stdlib.h>

typedef struct {
    int32_t a;
    int32_t b;

    int32_t r;
} Foo;

void sum(Foo *req) {
    req->r = req->a + req->b;
}
*/
import "C"

import "fmt"

func main() {
	req := C.Foo{
		a: 12,
		b: 30,
	}
	C.sum(&req)
	fmt.Println(req.r)
}
