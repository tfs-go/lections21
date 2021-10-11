package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	s := http.Server{
		Addr:        ":5000",
		Handler:     nil,
		ReadTimeout: time.Second,
	}
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/advanced", AdvancedHandler)
	log.Fatal(s.ListenAndServe())
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello, World!"))
}

func AdvancedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("method: %s\n", r.Method)
	fmt.Printf("query values: %v\n", r.URL.Query())
	fmt.Printf("headers: %v\n", r.Header)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	fmt.Printf("body: %s\n", string(body))

	w.WriteHeader(http.StatusAccepted)
	w.Header().Add("User-Agent", "Mozilla/5.0 (X11; Linux i686; rv:2.0.1) Gecko/20100101 Firefox/4.0.1")
	_, _ = w.Write([]byte("My name is..."))
}
