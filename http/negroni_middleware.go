// Using the negroni middleware package for logging

package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "You requested: "+r.URL.Path)
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", rootHandler)

	negroniHandler := negroni.New()
	negroniHandler.Use(negroni.NewLogger())

	negroniHandler.UseHandler(m)

	http.ListenAndServe("localhost:3000", negroniHandler)
}
