// Reading a CSV file to a struct

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	firstName string `json:"first_name"`
	lastName  string `json:"last_name"`
	ssn       string `json:"address"`
}

func main() {
	file, err := os.Open("dummy_data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	var people []person

	reader := csv.NewReader(bufio.NewReader(file))

	for {

		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		people = append(people, person{
			firstName: line[0],
			lastName:  line[1],
			ssn:       line[2],
		},
		)
	}

	for _, p := range people {
		fmt.Printf("Name: %v %v - SSN: %v \n", p.firstName, p.lastName, p.ssn)
	}
}
