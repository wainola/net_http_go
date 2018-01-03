// trabajando y ruteando con multiplexer
package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/hola.html", "templates/chao.html", "templates/index.html"))
}

func HolaMundo(w http.ResponseWriter, r *http.Request) {
	// ejecutar template.
	tpl.ExecuteTemplate(w, "hola.html", nil)

}

func ChaoMundo(w http.ResponseWriter, r *http.Request) {
	// ejecutar template
	tpl.ExecuteTemplate(w, "chao.html", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	// ejecutar template
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/hola", HolaMundo)
	http.HandleFunc("/chao", ChaoMundo)

	http.ListenAndServe(":8080", nil)
}
