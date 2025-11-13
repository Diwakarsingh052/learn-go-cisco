package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	//WaitGroup struct would be initialized with default values for the fields
	// & would give a pointer to the struct
	// make sure to use always a pointer to the WaitGroup
	// don't make copies of the data while passing it to the goroutine functions
	//wg := &sync.WaitGroup{}

	// new does exactly the same thing as above
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done() // done decrement the counter by one
		// sleep blocks the current goroutine
		//time.Sleep(2 * time.Second)
		n := rand.Int()
		if n%2 == 0 {
			fmt.Println("hello even")
			return
		}
		fmt.Println("odd number")

	}()

	wg.Add(1)
	go addV2(wg)

	// wait for all the goroutines to finish
	// wg.Wait blocks a goroutine until the counter is zero
	wg.Wait()
	fmt.Println("end of main")

}

func addV2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("addition")

}
