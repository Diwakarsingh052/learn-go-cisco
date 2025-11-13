package main

import "fmt"

func main() {
	calc(add)
	calc(func(a, b int) {
		fmt.Println(a-b, "from anonymous func")
	})
}

// to pass a function as an argument
// the signature of the function must match
func calc(f func(int, int)) {
	fmt.Println("calling calc func")
	// whenever f is called, it will execute the body of the function
	// that was passed as an argument
	f(10, 20)
}

func add(a, b int) { // datatype of a func includes args and return values
	fmt.Println(a + b)
}
