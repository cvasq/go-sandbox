package main

import (
	"flag"
	"fmt"
)

func main() {

	// Configure string flag
	tgtHostPtr := flag.String("target-host", "", "The target host (string)")

	flag.Parse()

	if *tgtHostPtr == "" {
		fmt.Println("Target host is empty")
	} else {

		fmt.Printf("The target host is: %v\n", *tgtHostPtr)
	}

}
