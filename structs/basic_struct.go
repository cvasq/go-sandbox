package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person1 := Person{"Carlos Vasquez", 26}

	fmt.Println(person1.Name)
	fmt.Println(person1.Age)
}
