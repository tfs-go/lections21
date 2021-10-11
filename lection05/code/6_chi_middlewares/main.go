package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/hello", GetHello)

	log.Fatal(http.ListenAndServe(":5000", r))
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello"))
}
