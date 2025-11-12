package main

import "learn-go-cisco/3-slices/inspect"

func main() {
	i := []int{10, 20, 30, 40, 50}
	b := i[1:3] // index:len // 20,30
	// len = 2 , cap = 4

	// b is sharing the same backing array with a slice,
	// any updates in b will also result changes in A slice
	b[0] = 999

	inspect.Slice("i", i)
	inspect.Slice("b", b)

}

/*
		a:= 10,20,30,40
		b := i[2:5]
a ,(b)-> [10,20,(30,40,50),60] // backing array
// a and b slice shares the same backing array. it is not a copy

b[0] = 999 // b is sharing the same backing array with a slice,
so any updates in b will also result changes in A slice
	    a ,(b)-> [10,20,(999,40,50),60]

*/
