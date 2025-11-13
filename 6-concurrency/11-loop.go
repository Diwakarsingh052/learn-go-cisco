package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	// buffered channel
	// deadlock in this program
	// we are receiving 5 values from the channel, but sending 12
	// buffer can only hold 1 value at any given time,
	// which means the sender is blocked on 7th value and deadlock
	
	ch := make(chan int, 1)

	for i := 1; i <= 6; i++ {
		wg.Go(func() {
			ch <- i
			ch <- i * i
		})
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	wg.Wait()
}
