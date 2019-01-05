// Launches 2 goroutines, one for each player
// A 'table' channel is created for communicating the ball

package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func player(name string, table chan *Ball) {

	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(1 * time.Second)
		table <- ball
	}

}

func main() {

	// table channel is used to pass the ball
	table := make(chan *Ball)

	go player("ping", table)
	go player("pong", table)

	fmt.Println("Game starts in 1 second...")
	time.Sleep(1 * time.Second)

	table <- new(Ball)

	// Game ends after 4 seconds
	time.Sleep(4 * time.Second)
}
