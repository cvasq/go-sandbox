package main

import (
	"fmt"
	"os"
)

func main() {

	// Print hostname
	hostname, _ := os.Hostname()
	fmt.Printf("Hostname: %v\n\n", hostname)

	// Print all env variables
	for _, envVariable := range os.Environ() {
		fmt.Println(envVariable)
	}

}
