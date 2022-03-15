package main

import (
	"log"
	"net/http"
	"text/template"
)

func Home(writer http.ResponseWriter, reader *http.Request) {
	templateHTML := template.Must(template.ParseFiles("main.html"))
	templateHTML.Execute(writer, nil)
}

func main() {
	log.Println("Server started on: http://localhost:8000")
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
}
