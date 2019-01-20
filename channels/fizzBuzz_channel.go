// FizzBuzz done with a string channel and go routine

package main

import (
	"fmt"
)

func fizzBuzz(n int) chan string {

	out := make(chan string, n)
	go func() {
		for i := 1; i <= n; i++ {

			if i%15 == 0 {
				out <- fmt.Sprintf("%v - %v", i, "FizzBuzz")
				continue
			}
			if i%3 == 0 {
				out <- fmt.Sprintf("%v - %v", i, "Fizz")
				continue
			}

			if i%5 == 0 {
				out <- fmt.Sprintf("%v - %v", i, "Buzz")
				continue
			}

			out <- fmt.Sprintf("%v", i)
		}
		close(out)
	}()

	return out
}

func main() {

	for v := range fizzBuzz(100) {
		fmt.Println(v)
	}
}
