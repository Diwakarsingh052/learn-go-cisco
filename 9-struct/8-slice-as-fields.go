package main

import "fmt"

type student struct {
	name  string
	marks []int
}

// marks is a variadic parameter
// it can take any number of arguments of type int
// variadic parameters are always last in the parameter list
// variadic parameter is optional

// use pointer receiver to modify the slice

func (s *student) AddMarks(marks ...int) {
	fmt.Printf("marks: %T\n", marks)

	// ... will unpack the slice
	// individual elements of the slice will be passed as arguments
	s.marks = append(s.marks, marks...)
	//s.marks = append(s.marks, 40, 50, 70, 80)
	//s.marks = append(s.marks, []int{90, 100}...)
}
func main() {
	s := student{name: "John"}
	s.AddMarks(10, 20, 30)
	fmt.Println(s)
}
