package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {

	person1 := Person{"Carlos V", 26, "Male"}

	fmt.Println("Name:", person1.Name)
	fmt.Println("Age:", person1.Age)
	fmt.Println("Gender:", person1.Gender)

}
