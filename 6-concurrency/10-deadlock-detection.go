package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int)

	// goroutine leak
	// goroutine which is blocked forever and its memory is never released
	wg.Go(func() {
		ch <- 10
	})

	// go would not report deadlock if other goroutine is routine is running and not in blocked state
	wg.Go(func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("sleeping for 1 second")
		}

	})

	wg.Wait()

}
