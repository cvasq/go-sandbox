// 1. Accepts serialized data via STDIN
// 2. Unmarshals data to empty struct
// 3. Reads struct fields

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	proto "github.com/golang/protobuf/proto"
)

func readStdin() []byte {
	var stdinBuffer []byte
	stdinBuffer, _ = ioutil.ReadAll(os.Stdin)
	return stdinBuffer
}

func main() {

	// Collect serialized data from STDIN
	data := readStdin()

	// Initialize empty struct
	person := &Person{}

	// Unmarshal data
	err := proto.Unmarshal(data, person)
	if err != nil {
		log.Fatal("Unmarshalling Error:", err)
	}

	// Read struct fields
	age := person.GetAge()
	name := person.GetName()

	fmt.Printf("New User\nname: %v\nage: %v\n", name, age)
}
