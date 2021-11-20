.DEFAULT_GOAL := all

main.js: main.go
	gopherjs build -o main.js -m

.PHONY: all
all: main.js

.PHONY: clean
clean:
	@rm -f main.js*
