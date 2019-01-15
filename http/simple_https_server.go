// A simple HTTPS web server created with the net/http package
// Returns the message specified in the URL query string
//
// Test with curl
// curl -k https://localhost:8080/?message=Hello

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const listeningPort = "8080"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	message := r.URL.Query().Get("message")

	if message == "" {
		message = "No message was provided"
	}
	response := map[string]string{"Message": message}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", rootHandler)

	log.Println("Listening for HTTPS connections on port:", listeningPort)

	err := http.ListenAndServeTLS(":"+listeningPort, "selfsigned-x509.crt", "selfsigned.key.pem", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
