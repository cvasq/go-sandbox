## Using Go Channels

A channel is a mechanism for goroutines to synchronize execution and communicate by passing values.

Creating a new channel
```
// Create unbuffered bidirectional channel of type int
intChannel := make(chan int)

// Create buffered bidirectional channel with room for 10 ints
intChannel := make(chan int, 10)

// Create 'Send only' channel
intChannel := make(chan<- int)

// Write data to a channel
intChannel <- 5

// Read data off a channel
myNum := <- intChannel

// Close channel
close(intChannel)
```

#### Examples

* **done_signal_channel.go** - Using a channel to end a 'for' loop
* **ping_pong.go** - Communicating between 2 goroutines
* **fizzBuzz_channel.go** - The Fizz Buzz test done with a channel and go routine