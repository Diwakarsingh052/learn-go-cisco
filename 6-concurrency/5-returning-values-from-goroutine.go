package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	// we can't return values from goroutines
	//str := go helloV2()
	wg.Go(func() {
		str := helloV2()
		fmt.Println(str)
	})
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//
	//}()
	fmt.Println("more things in main")
	wg.Wait()
}

func helloV2() string {

	return "hello"
}
