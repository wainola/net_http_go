// Fundaciones en net/http
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

// generamos un tipo que implicitamente implementa la interfaz Handler
type nicolas string

// tipo con datoa de los formularios para ordenar el asunto
type FormData struct {
	Nombre string
	Boton  string
}

// metodo para type nicolas que implementa el metodo de la interfaz de Handler
func (n nicolas) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(req.Form)
	data := req.Form
	fData := FormData{Nombre: strings.Join(data["nombre"], ""), Boton: strings.Join(data["btn_submit"], "")}
	fmt.Println(fData)
	tpl.ExecuteTemplate(w, "index.html", fData)
}

func main() {
	var n nicolas
	http.ListenAndServe(":8080", n)
}
