// Writing structs to a CSV file

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	csvFile, err := os.Create("exported_data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvFile.Close()

	person1 := Person{FirstName: "Carlos", LastName: "V", Age: 27}
	person2 := Person{FirstName: "Miguel", LastName: "V", Age: 24}

	personToSlice := func(p Person) []string {
		return []string{p.FirstName, p.LastName, strconv.Itoa(p.Age)}
	}

	var people [][]string

	people = append(people, personToSlice(person1))
	people = append(people, personToSlice(person2))

	writer := csv.NewWriter(csvFile)
	for _, person := range people {
		err := writer.Write(person)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	writer.Flush()
}
