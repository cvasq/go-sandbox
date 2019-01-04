package main

import (
	"fmt"

	"github.com/rs/xid"
)

type Student struct {
	Name   string
	Age    int
	ID     string
	Points int
}

// newStudent constructs and returns a pointer to a new Student
func newStudent(name string, age int, points int) *Student {
	id := generateStudentID()
	s := &Student{name, age, id, points}
	return s
}

// Generate a unique ID to identify the student
func generateStudentID() string {
	id := xid.New()
	return fmt.Sprintf("%s", id.String())
}

func main() {

	s := newStudent("Brian Peppers", 16, 95)

	fmt.Println("Adding new student...\n")
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
	fmt.Println("ID:", s.ID)
	fmt.Println("Points:", s.Points)
}
