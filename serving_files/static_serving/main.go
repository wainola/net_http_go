// servimos sitios estaticos con go
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/login.html"))
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Page")
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func main() {
	http.HandleFunc("/login", login)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))

	http.ListenAndServe(":8080", nil)
}
