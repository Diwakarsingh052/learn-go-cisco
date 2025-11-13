package main

import "fmt"

// A send on an unbuffered channel can proceed if a receiver is ready.
// A send on a buffered channel can proceed if there is room in the buffer.

func main() {
	ch := make(chan int, 1)

	ch <- 10
	// receiving value from a buffered channel will remove the value from the buffer
	fmt.Println(<-ch)
	ch <- 20

}

// create four goroutines
// first two goroutine will be used to work with unbuffered channel
// third and fourth goroutine will be used to work with buffered channel

// in unbuffered channel goroutine
// sender sends 5 values with a delay of 1 second
// once all the values are sent, the goroutine prints all value sent msg

// receiver receives the values and prints them

// in buffered channel goroutine
// sender sends 5 values with a delay of 1 second
// once all the values are sent, the goroutine prints all value sent msg

// receiver receives the values and prints them
