// ejercicios para reforzar rutas y el uso de mux
package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.URL)
	fmt.Println("Funcion Index")
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Llamando a la func Perro o dog")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Println("La url ingresada es: ...")
	fmt.Println(r.URL)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
