// A basic http client
// Example from Security with Go (ISBN: 978-1788627917)

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatal("Error fetching URL: ", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	fmt.Printf("%s\n", body)
}
