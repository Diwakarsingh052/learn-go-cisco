package main

import "fmt"

func main() {
	// default value of pointer is nil
	//var p *int // pointer to int

	var x int = 10
	var p *int = &x
	updateValue(p) // & is the address operator
	fmt.Println(x)
}

func updateValue(p *int) {
	*p = 20

}
