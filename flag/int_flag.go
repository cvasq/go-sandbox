package main

import (
	"flag"
	"fmt"
)

func main() {

	// Configure options
	quantityPtr := flag.Int("quantity", 1, "Quantity (int)")
	flag.Parse()
	fmt.Printf("The quantity is: %v\n", *quantityPtr)

}
