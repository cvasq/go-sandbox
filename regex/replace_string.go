// Replacing a string with the regexp package

package main

import (
	"fmt"
	"regexp"
)

// Sample text
var fullText string = "This is a line with Foo"

func main() {

	// Word to match for replacement
	regex := regexp.MustCompile("Foo")

	// Replacement word
	replacement := "Bar"

	// Replace with ReplaceAllString function
	// Prints "This is a line with Bar"
	fmt.Println(regex.ReplaceAllString(fullText, replacement))
}
