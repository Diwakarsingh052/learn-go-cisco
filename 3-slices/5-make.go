package main

import (
	"fmt"
	"learn-go-cisco/3-slices/inspect"
)

func main() {
	// make pre-allocates underlying array
	// append will reuse the same memory until the capacity is reached

	// if len is provided then slice would intialize with default values of the type
	// if len is 2, then slice would have 0,0 values
	// append would add element to the third location after that,
	
	s := make([]int, 0, 50000) // make(type,len,cap)
	fmt.Println(s)
	inspect.Slice("s", s)

	s = append(s, 10)
	inspect.Slice("s", s)

}
