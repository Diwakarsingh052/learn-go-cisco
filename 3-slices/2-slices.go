package main

import "fmt"

// https://pkg.go.dev/fmt#pkg-overview

func main() {
	// slice default value is nil
	// nil means no memory allocated
	var s []int
	fmt.Printf("%#v\n", s)
	//var x = []int{} // empty slice, memory is initialized, but no value is stored

	a := []int{10, 20, 30}
	a[0] = 100 // update operation
	//a[3] = 200 // runtime error, if length is not present
	fmt.Println(a)

}
