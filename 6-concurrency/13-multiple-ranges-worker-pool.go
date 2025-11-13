package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 5)
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

	// running 2 workers to process the send values
	for i := 1; i <= 2; i++ {
		wg.Go(func() {
			for v := range ch {
				fmt.Println(v, "received by", i)
				time.Sleep(1 * time.Second)
			}
		})
	}

	wg.Wait()

}
