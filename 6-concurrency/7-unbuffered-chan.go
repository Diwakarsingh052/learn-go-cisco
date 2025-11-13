package main

import (
	"fmt"
	"sync"
)

//A send on an unbuffered channel can proceed if a receiver is ready.
//A send on a buffered channel can proceed if there is room in the buffer.
// receiver blocks until the sender sends the value

func main() {
	ch := make(chan int) // unbuffered channel
	wg := new(sync.WaitGroup)

	wg.Go(func() {
		// if this goroutine gets picked up first then it would block
		// because there is no receiver
		// and go scheduler would pick up main to execute
		// once in the main receiver is ready, the send would proceed
		ch <- 10
	})

	// receiver always blocks until the sender sends the value
	fmt.Println(<-ch)

	wg.Wait()

}
