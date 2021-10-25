package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:     nil,
		AllowOriginFunc:    nil,
		AllowedMethods:     nil,
		AllowedHeaders:     nil,
		ExposedHeaders:     nil,
		AllowCredentials:   false,
		MaxAge:             0,
		OptionsPassthrough: false,
		Debug:              false,
	}))

	log.Fatal(http.ListenAndServe(":5000", r))
}
