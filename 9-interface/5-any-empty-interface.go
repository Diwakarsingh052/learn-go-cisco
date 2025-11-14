package main

import "fmt"

func main() {
	//var i interface{}
	var i any = 10
	i = "hello"

	// type assertion
	// checking if the interface is of type int
	// always use ok variant, it avoids panics
	x, ok := i.(int)
	if ok { // ok == true
		fmt.Println(x, "interface converted to int")
	}

	i = "hello"
	i = true

	fmt.Println(i)
}
