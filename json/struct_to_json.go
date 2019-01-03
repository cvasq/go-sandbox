package main

import (
	"encoding/json"
	"fmt"
)

type Package struct {
	Width  int
	Height int
	Value  float64
	Weight int
}

func main() {

	p := Package{
		Width:  14,
		Height: 10,
		Value:  100.43,
		Weight: 4,
	}

	pBytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
	}

	jsonData := string(pBytes)
	fmt.Println(jsonData)
}
