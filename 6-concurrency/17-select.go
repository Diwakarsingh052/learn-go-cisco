package main

import (
	"fmt"
	"sync"
	"time"
)

// in this program select usage can be replaced with ranges
// we can run individual goroutines for each kind of request [(get, post, put)]
// each goroutine receives a value from the one channel it is assigned to
func main() {
	get := make(chan string)
	post := make(chan string)
	put := make(chan string)

	wg := new(sync.WaitGroup)
	wg.Go(func() {
		time.Sleep(3 * time.Second)
		get <- "get"
	})

	wg.Go(func() {
		time.Sleep(1 * time.Second)
		post <- "post"
	})

	wg.Go(func() {
		put <- "put"
	})
	//fmt.Println(<-get)
	//fmt.Println(<-post)
	//fmt.Println(<-put)

	// without the loop select evaluates only once
	// to run select multiple times we need a loop
	for i := 0; i < 3; i++ {
		select {
		//whichever case is not blocking exec that first
		//whichever case is ready first, exec that.
		// possible cases are chan recv , send , default
		case v := <-get:
			fmt.Println(v)
		case v := <-post:
			fmt.Println(v)
		case v := <-put:
			fmt.Println(v)
		}
	}
	wg.Wait()

}
