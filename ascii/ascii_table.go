package main

import (
	"github.com/olekukonko/tablewriter"
	"github.com/telyn/row"
	"os"
)

type Person struct {
	Name   string
	Age    int
	City   string
	Gender string
}

func printPeopleTable() {

	header := []string{"Name", "Age", "City", "Gender"}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)

	var peopleRows []Person

	peopleRows = append(peopleRows, Person{"Serj Tankian", 51, "Los Angeles", "Male"})
	peopleRows = append(peopleRows, Person{"Daron Malakian", 43, "Canyon Country", "Male"})
	peopleRows = append(peopleRows, Person{"Shavarsh Odadjian", 44, "Los Angeles", "Male"})
	peopleRows = append(peopleRows, Person{"John Dolmayan", 45, "Los Angeles", "Male"})

	for _, v := range peopleRows {

		values, err := row.From(v, header)
		if err != nil {
			panic(err)
		}
		table.Append(values)
	}

	table.Render()
}

func main() {
	printPeopleTable()
}
