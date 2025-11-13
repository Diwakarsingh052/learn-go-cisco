package main

import "fmt"

// close the channel at the correct place
// and use waitgroups to wait for all goroutines to finish
// make sure no deadlock occurs
func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 4; i++ {
			go func() {
				ch <- i
			}()
		}
	}()

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

}
