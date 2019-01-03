package main

import (
	"fmt"
	"os"
)

func main() {

	os.Setenv("TEST", "1")

	fmt.Println("User:", os.Getenv("USER"))
	fmt.Println("Test:", os.Getenv("TEST"))
	fmt.Println("Bar (Not set):", os.Getenv("BAR"))
}
