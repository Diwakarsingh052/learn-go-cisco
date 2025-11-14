package main

import "fmt"

type Person struct {
	name    string // fields of struct
	age     int
	address string
}

// money is a new type
type money int

func main() {
	// p var would be initialized with default values of the fields
	var p Person
	fmt.Printf("%#v\n", p)

	p1 := Person{
		name: "abc",
		age:  30,
		// not mandatory to provide values to all the fields
		//address: "abc",
	}

	// %+v prints the fields with their values
	fmt.Printf("%+v\n", p1)
	fmt.Println(p1.age)
}
