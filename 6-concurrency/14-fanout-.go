package main

import (
	"fmt"
	"sync"
)

// close the channel at the correct place
// and use waitgroups to wait for all goroutines to finish
// make sure no deadlock occurs
func main() {
	ch := make(chan int)

	wg := new(sync.WaitGroup)
	wgTask := new(sync.WaitGroup)

	wg.Go(func() {
		for i := 1; i <= 4; i++ {
			// fan out pattern
			// spinning n number of goroutines for n tasks
			wgTask.Go(func() {
				ch <- i
			})
		}
		// wait for all goroutines to finish
		wgTask.Wait()
		close(ch) // close the channel once all tasks are done

	})

	wg.Go(func() {
		for v := range ch {
			fmt.Println(v)
		}
	})
	wg.Wait()
}
