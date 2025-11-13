package main

import (
	"fmt"
	"sync"
	"time"
)

// using channels, we can send and receive data between goroutines
// two types of channels, buffered and unbuffered

// unbuffered channel has size of 0
// buffered channel we can specify the size

func main() {
	// make(chan type, size)
	ch := make(chan string) // unbuffered channel
	//ch := make(chan int, 10) // buffered channel

	wg := new(sync.WaitGroup)
	wg.Go(func() {
		time.Sleep(2 * time.Second)
		str := helloV3()
		ch <- str // send signal to channel
	})

	fmt.Println("waiting to receive data from channel")
	// receive would block the goroutine until the value is not sent
	str := <-ch // receive data from channel
	fmt.Println(str, "received from channel")
	//fmt.Println(<-ch)

}

func helloV3() string {
	return "hello"
}
