package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	wg.Go(func() {
		x := SlowFuncV2()
		ch <- x
	})

	// select would do deadlock if ctx.Done case evaluates before ch <- x
	// after ctx.Done case we would move further in the code
	// and after coming out of select, receiver would not be present to receive the value
	// and later on when goroutine sends the value it would be blocked
	// because it can't send the value without a receiver(property of unbuffered channel)
	select {
	case v := <-ch:
		fmt.Println("value received", v)
	case <-ctx.Done():
		fmt.Println("context done")
		fmt.Println(ctx.Err())

	}
	fmt.Println("moved on from the select no receiver to receive the value")
	wg.Wait()
}

func SlowFuncV2() int {
	time.Sleep(2 * time.Second)
	fmt.Println("slow function executed")
	return rand.Int()

}
