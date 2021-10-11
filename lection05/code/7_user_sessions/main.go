package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	cookieAuth           = "auth"
	userID     cookieVal = "ID"
)

type cookieVal string

func main() {
	root := chi.NewRouter()
	root.Use(middleware.Logger)
	root.Post("/login", Login)

	r := chi.NewRouter()
	r.Use(Auth)
	r.Get("/hello", GetHello)

	root.Mount("/api", r)

	log.Fatal(http.ListenAndServe(":5000", root))
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value(userID).(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write([]byte("hello, " + id))
}

type U struct {
	Login string
}

func Login(w http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var u U
	err = json.Unmarshal(d, &u)
	if err != nil || u.Login == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c := &http.Cookie{
		Name:  cookieAuth,
		Value: u.Login,
		Path:  "/",
	}

	http.SetCookie(w, c)
}

func Auth(handler http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(cookieAuth)
		switch err {
		case nil:
		case http.ErrNoCookie:
			w.WriteHeader(http.StatusUnauthorized)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if c.Value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		idCtx := context.WithValue(r.Context(), userID, cookieVal(c.Value))

		handler.ServeHTTP(w, r.WithContext(idCtx))
	}

	return http.HandlerFunc(fn)
}
