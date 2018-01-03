// analizando el request
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// creamos un tipo cualquiera para implementar la interfaz ServeHTTP
type camilo string

// insertamos en struct los valores de request.
type DataReq struct {
	Method string
	Url    *url.URL
	Body   io.ReadCloser
	// retorna un map
	Header http.Header
}

func (n camilo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d := DataReq{
		Method: r.Method,
		Url:    r.URL,
		Body:   r.Body,
		Header: r.Header,
	}

	fmt.Println(d.Method)
	fmt.Println(d.Url)
	fmt.Println(d.Body)
	fmt.Println(d.Header)
}

func main() {
	var c camilo
	http.ListenAndServe(":8080", c)
}
