## Using Go Channels

A channel is a mechanism for goroutines to synchronize execution and communicate by passing values.

Creating a new channel
```
// Unbuffered channel of type int
intChannel := make(chan int)

// Buffered channel with room for 10 ints
intChannel := make(chan int, 10)
```

#### Examples

* **done_signal_channel.go** - Using a channel to end a 'for' loop
* **ping_pong.go** - Communicating between 2 goroutines
