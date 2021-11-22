package main

/*
#include <stdio.h>
#include <stdlib.h>

#pragma pack(push,1)
typedef struct {
    int8_t a;
    int32_t b;
} Bar;
#pragma pack(pop)

typedef struct {
    int8_t a;
    int32_t b;
} Foo;

void test() {
    printf("Bar size: %lu\n", sizeof(Bar));
    printf("Foo size: %lu\n", sizeof(Foo));
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	C.test()

	fmt.Printf("C.Foo size: %d\n",
		unsafe.Sizeof(C.Foo{}),
	)
}
