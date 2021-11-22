.DEFAULT_GOAL := all

main.wasm: wasm.go wasm_exec.js
	tinygo build -target wasm -o main.wasm -tags wasm --no-debug wasm.go

wasm_exec.js:
	# https://tinygo.org/docs/guides/webassembly/#building
	cp /usr/local/lib/tinygo/targets/wasm_exec.js .

.PHONY: all
all: main.wasm serve

.PHONY: serve
serve: serve.go
	go run .

.PHONY: clean
clean:
	@rm -f main.wasm
