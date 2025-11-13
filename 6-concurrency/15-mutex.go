package main

import (
	"fmt"
	"sync"
)

// shared resources
/*
1. Global Variables
2. Pointer variables
3. Structs fields where methods are pointers
4. Maps,
*/

// data race situations
//	- at least one concurrent write and n number of reads
//	- n number of concurrent writes
// 	- n number of concurrent writes and n number of concurrent reads
// 	Note - Data race doesn't happen if there are only concurrent reads

var x int

func main() {
	m := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	for i := 1; i <= 4; i++ {
		wg.Go(func() {
			UpdateX(i, m)
		})
	}
	wg.Wait()
	fmt.Println(x)
}

func UpdateX(val int, m *sync.Mutex) {
	// critical section
	// this is the place where we access the shared resource

	// when a goroutine acquires a lock,
	//another goroutine can't access the critical section
	// until the lock is not released
	m.Lock()
	defer m.Unlock() // release the lock when the function returns
	x = val

}
