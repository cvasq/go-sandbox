// Using the negroni middleware package for logging

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

const listeningPort = "8080"

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "You requested: "+r.URL.Path)
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", rootHandler)

	negroniHandler := negroni.New()
	negroniHandler.Use(negroni.NewLogger())

	negroniHandler.UseHandler(m)

	log.Println("Listening for HTTP connections on port:", listeningPort)
	err := http.ListenAndServe(":"+listeningPort, negroniHandler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
