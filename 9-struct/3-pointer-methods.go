package main

import "fmt"

type person struct {
	name string
}

func (p *person) updateName(name string) {
	// p is a pointer to the struct
	// any changes done by p would reflect in the original struct
	p.name = name
}

func main() {
	p := person{"John"}
	// without pointer struct data would be copied
	// any changes by a method would not reflect in the original struct

	//if receiver is a pointer then variable address is passed to the receiver
	p.updateName("Jane")
	fmt.Println(p.name)
}
