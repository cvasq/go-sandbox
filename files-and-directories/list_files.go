package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var downloadDir = "/home/toor/Downloads/"

func main() {
	downloads, err := ioutil.ReadDir(downloadDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range downloads {
		fmt.Println(file.Name())
	}
}
