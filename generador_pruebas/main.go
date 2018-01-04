// web app que guarda informacion de ejercicios en banco de datos y genera pruebas en word.
package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	http.HandleFunc("/", index)
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.ListenAndServe(":8080", nil)
}
