package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
	Address
}

type Address struct {
	Street  string
	Zipcode string
}

func main() {

	var students []Student

	student1 := Student{
		Name: "Robert Tappan",
		Age:  37,
		Address: Address{
			Street:  "9 Main Street",
			Zipcode: "22303",
		},
	}

	student2 := Student{
		Name: "Alexandra Elbakyan",
		Age:  30,
		Address: Address{
			Street:  "53 Industrial Park",
			Zipcode: "21520",
		},
	}
	students = append(students, student1)
	students = append(students, student2)

	// Loop through slice of Students
	// Street and Zipcode are promoted fields
	for _, s := range students {
		fmt.Println("Name:", s.Name)
		fmt.Println("Age:", s.Age)
		fmt.Println("Street:", s.Street)
		fmt.Println("Zipcode:", s.Zipcode, "\n")
	}
}
