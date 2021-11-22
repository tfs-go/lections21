.DEFAULT_GOAL := all

mylib.a: mylib.go
	go build -o mylib.a -buildmode=c-archive mylib.go

libtest.o: mylib.a libtest/
	cd libtest && gcc -o ../libtest.o -pthread libtest.c ../mylib.a

.PHONY: all
all: libtest.o

.PHONY: clean
clean:
	@rm -f ./mylib.a ./mylib.h ./libtest.o
