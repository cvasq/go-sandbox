package main

import (
	"fmt"
	"time"
)

func helloWorld() string {
	currentTime := time.Now()
	return fmt.Sprintf("Hello World! The current time is: %v", currentTime.Format("3:04 PM 2006-01-02"))
}

func main() {
	fmt.Println(helloWorld())
}
