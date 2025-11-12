package main

import "fmt"

// https://go.dev/ref/spec#Appending_and_copying_slices
// If the capacity of s is not large enough to fit the additional values,
// append allocates a new, sufficiently large underlying array
// that fits both the existing slice elements and the additional values.
// Otherwise, append re-uses the underlying array.

func main() {
	// len, cap
	// slices are references
	// slices are three word data structure
	// 1- > len, 2-> cap, 3-> pointer to array

	// len is the number of elements in the slice
	// cap is the number of elements your array can hold
	a := []int{10, 20, 30}
	fmt.Println(len(a), cap(a), &a[0])
	a = append(a, 40)
	fmt.Println(len(a), cap(a), &a[0])

	a = append(a, 50)
	fmt.Println(len(a), cap(a), &a[0])
	fmt.Println(a)

}
