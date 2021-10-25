package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ViewData struct {
	Title string
	Users []User
}
type User struct {
	Name string
	Age  int
}

func main() {
	data := ViewData{
		Title: "Users List",
		Users: []User{
			{Name: "Tom", Age: 21},
			{Name: "Kate", Age: 23},
			{Name: "Alice", Age: 25},
		},
	}

	tmpl, err := template.ParseFiles("lection06/code/3_html_template/users.tmpl")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = tmpl.Execute(w, data)
	})

	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
