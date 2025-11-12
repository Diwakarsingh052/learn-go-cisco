package main

import "fmt"

func main() {
	//var firstName string // camelCase for variable naming
	// every types have a default value
	var a int // int default is 0
	var b string = "hello"
	var c = true
	_, _ = c, b
	//c = 10 // we cant change the type // go is a statically typed language
	fmt.Println(a, b)

	// go compiler would infer the type automatically from the right side value
	d, ok := 100, true // shorthand// create and assign
	_, _ = d, ok

	//a = "1" // we cant change the type
	var x int
	var f float64 = float64(x)
	_ = f
}
