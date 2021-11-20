#!/bin/env luajit

local ffi = require("ffi")
ffi.cdef[[
void HelloWorld2(int p0, const char* p1);
]]

local mylib = ffi.load("./mylib.so")

mylib.HelloWorld2(42, "Hello from Lua!")
