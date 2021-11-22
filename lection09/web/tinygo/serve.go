//go:build !wasm
// +build !wasm

package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	argv struct {
		Addr string
		Help bool
	}
)

func main() {
	flag.StringVar(&argv.Addr, `listen`, `:8000`, `Address to listen`)
	flag.BoolVar(&argv.Help, `h`, false, `Show this help`)
	flag.Parse()

	if argv.Help {
		flag.PrintDefaults()
		return
	}

	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	log.Printf("Listening on %s...", argv.Addr)
	err := http.ListenAndServe(argv.Addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
