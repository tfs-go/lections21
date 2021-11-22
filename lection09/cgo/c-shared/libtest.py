#!/bin/env python2

from ctypes import *

lib = cdll.LoadLibrary("./mylib.so")

class GoString(Structure):
    _fields_ = [("p", c_char_p), ("n", c_longlong)]

lib.HelloWorld.argtypes = [c_int, GoString]

msg = GoString("Hello from Python", 17)
lib.HelloWorld(42, msg)
