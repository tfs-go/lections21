.DEFAULT_GOAL := all

main.wasm: wasm.go wasm_exec.js
	GOOS=js GOARCH=wasm go build -o main.wasm

wasm_exec.js:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js .

.PHONY: all
all: main.wasm serve

.PHONY: serve
serve: serve.go
	go run .

.PHONY: clean
clean:
	@rm -f main.wasm
