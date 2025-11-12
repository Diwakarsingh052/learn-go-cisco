package main

import "fmt"

func main() {
	var x int

	if x == 0 {
		// x is shadowed
		// x is local variable in this scope
		// this x is different from the global x
		x, msg := doSomething()
		fmt.Println(x, msg) //10, done
	}
	// x is still 0
	fmt.Println(x)
}

func doSomething() (int, string) {
	return 10, "done"
}
