// Generates any desired number of random 'silly names'.
// By default 1 name is generated.
// The --quantity flag controls the number of names generated.

package main

import (
	"flag"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

func generateRandomNames(quantity int) []string {

	var sillyNames []string

	for i := 0; i < quantity; i++ {

		sillyNames = append(sillyNames, randomdata.SillyName())
	}

	return sillyNames

}

func main() {

	// Configure options
	quantityPtr := flag.Int("quantity", 1, "Quantity (int)")
	flag.Parse()

	for _, v := range generateRandomNames(*quantityPtr) {
		fmt.Println(v)
	}
}
