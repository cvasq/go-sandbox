package main

import (
	"fmt"
	"io/ioutil"
)

// ascii_art.txt file has banner
func banner() string {
	b, err := ioutil.ReadFile("ascii_art.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main() {
	fmt.Printf(banner())
}
