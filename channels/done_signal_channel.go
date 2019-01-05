// This program launches a goroutine which will close a 'done' channel after 5 seconds
// The 'done' channel is used as a signal for breaking out of a named for loop

package main

import (
	"fmt"
	"time"
)

func main() {

	// Initialize 'done' channel
	done := make(chan interface{})

	// goroutine: Close 'done' channel after 5 seconds
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	// Loop counter
	loopCounter := 0

loop:
	for {
		// Break out of named for loop upon receiving signal from 'done' channel
		select {
		case <-done:
			break loop
		default:
			fmt.Println("Hit default select case")
		}

		// Increase loop counter and sleep 1 second
		loopCounter++
		time.Sleep(1 * time.Second)
	}

	// Exited the for loop
	fmt.Printf("Looped %v times before receiving signal to stop.\n", loopCounter)
}
