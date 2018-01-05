// web app que guarda informacion de ejercicios en banco de datos y genera pruebas en word.
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.html", "templates/partial/header.html", "templates/partials/footer.html", "templates/alternativas.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func next(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Funcion next")
	http.Redirect(w, r, "/test2", 301)
}

func main() {
	http.HandleFunc("/", index)
	// ruta dummy para probar pegarle con ajax
	http.HandleFunc("/siguiente", next)

	// Servir los assets estaticos
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.ListenAndServe(":8080", nil)
}
