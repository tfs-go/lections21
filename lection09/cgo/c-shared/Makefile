.DEFAULT_GOAL := all

mylib.so: mylib.go
	go build -o mylib.so -buildmode=c-shared mylib.go

libtest.o: mylib.so libtest/libtest.c
	cd libtest && gcc -o ../libtest.o libtest.c -L.. -l:./mylib.so

.PHONY: all
all: libtest.o

.PHONY: clean
clean:
	@rm -f ./mylib.so ./mylib.h ./libtest.o
