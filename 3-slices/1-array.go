package main

import "fmt"

func main() {
	// array has fixed size
	var a [5]int = [5]int{1, 2, 3}
	a[0] = 10
	fmt.Println(a)
	for _, v := range a { // index, value
		fmt.Println(v)
	}
}
