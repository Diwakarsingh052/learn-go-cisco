package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("hello")
	go add(1, 2)

	//anonymous function
	// function without name
	// anonymous function can be called directly when it is created
	go func(i int) {
		//time.Sleep(2 * time.Second)
		fmt.Println("anonymous function called with ", i, "")
	}(10) // () is calling the anonymous function
	fmt.Println("end of main")

	// time.Sleep would block the main goroutine
	// other goroutines can execute in the meantime
	// time.Sleep is not recommended to use
	time.Sleep(1 * time.Second)
}

func add(a, b int) {
	fmt.Println("addition", a+b)
}
