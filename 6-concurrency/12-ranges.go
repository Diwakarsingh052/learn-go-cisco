package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 1)
	wg := new(sync.WaitGroup)

	wg.Go(func() {

		for i := 1; i <= 7; i++ {
			ch <- i
		}

		// closing the channel
		// once channel is closed, no more values can be sent to it
		// range can receive remaining values even after channel is closed
		// and range would stop after receiving all the values
		close(ch)
	})

	// range is a receive operation
	// range runs infinitely until the channel is closed
	for v := range ch {
		fmt.Println(v)
	}

	wg.Wait()

}
