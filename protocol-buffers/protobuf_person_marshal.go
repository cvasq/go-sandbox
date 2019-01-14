// Marshals data from --name [string] and --age [int] using Protocol Buffers
package main

import (
	"flag"
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {

	// Configure flags
	name := flag.String("name", "", "The name (string)")
	age := flag.Int("age", 0, "Age (int)")

	flag.Parse()

	// Create struct literal
	person := &Person{
		Name: *name,
		Age:  int32(*age),
	}

	// Marshal data
	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("Marshalling Error:", err)
	}

	// uninterpreted bytes -> STDOUT
	fmt.Printf("%s", data)
}
