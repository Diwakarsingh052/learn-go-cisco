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
		x := SlowFuncV3()
		// if ctx.Done happens, then no receiver is present
		// we quit this goroutine
		select {
		case ch <- x:
			fmt.Println("value sent")
		case <-ctx.Done():
			fmt.Println("can't send the value, receiver lost")
		}
	})

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

func SlowFuncV3() int {
	time.Sleep(2 * time.Second)
	fmt.Println("slow function executed")
	return rand.Int()

}
