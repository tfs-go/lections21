#include <stdio.h>

#include "../mylib.h"

int main() {
    GoInt i = 42;
    GoString msg = {"C Caller", 8};
    // HelloWorld(i int32, msg string)
    HelloWorld(i, msg);

    HelloWorld2(100500, "C caller#2");
}
