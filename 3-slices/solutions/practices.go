package main

import "fmt"

func main() {
	x := make([]int, 0, 3)
	x = append(x, 10, 20, 30)
	//x := []int{10, 20, 30}
	//UpdateSlice(x, 1, 40)
	//fmt.Println(x)

	// if you want to append, avoid using slicing, it can add confusion
	x = UpdateAndAppend(x, 1, 50)
	fmt.Println(x)
}

func UpdateSlice(s []int, index, value int) {
	// if you want to update the value ,
	//then passing the reference of existing backing array is fine
	// we don't have to use copy
	s[index] = value
}

func UpdateAndAppend(s []int, index int, value int) []int {
	s[index] = value
	// make sure to return the updated slice if using append
	// or use double pointer
	s = append(s, 100)
	return s
}
