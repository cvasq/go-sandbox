package main

import (
	"flag"
	"fmt"
)

func main() {

	// Configure options
	dryRunPtr := flag.Bool("dry-run", false, "a bool")
	flag.Parse()
	fmt.Printf("Dry Run Flag Enabled: %v\n", *dryRunPtr)

}
