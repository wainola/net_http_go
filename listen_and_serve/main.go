// usando listen and serve para generar un servidor.
package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Any code you want en esta funcion hermano!")
}

func main() {
	/*
		Notar la particularidad del ejercicio:
		hotdog es un topo entero, pero implementa el metodo ServeHTTP.
		Al llevar a cabo esta implementacion del metodo, hotdog se convierte en un handler que puede ser pasado
		al http.ListenAndServe...

		Por lo que hotdog es tanto un tipo entero, como un tipo handler

		Recordar en Handler es una interfaz, por lo que todo aquel que implemente los metodos de la interfaz, toma el tipo de esta interfaz
	*/

	var d hotdog
	http.ListenAndServe(":8080", d)
}
