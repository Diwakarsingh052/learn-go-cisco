package main

import "fmt"

// double pointers are useful when we want to update the original pointer
// usually nil pointers
func main() {
	var p *int // value of p is nil // address of p = x100
	fmt.Println(&p)
	updateNilPointer(&p)
	fmt.Println(*p)

	// examples from standard library where double pointer is used
	//errors.As()
	//json.Unmarshal()

}

func updateNilPointer(p1 **int) {
	x := 10 // assume address x90
	fmt.Println("& of x", &x)
	*p1 = &x  // accessing the pointer p from the main function
	**p1 = 20 // ** would give access to the concrete value

}
