package main

import "fmt"

type Student struct {
	name string
	age  int
}

// methods: func (receiver) name() returnType{}
// receiver is the variable of the struct
func (s Student) printName() {
	fmt.Println(s.name)
}

func main() {
	s := Student{"John", 20}
	s.printName()

}
