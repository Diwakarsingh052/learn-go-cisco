package main

import "fmt"

func main() {
	// in Go everything is pass by value
	// in case of pointers we copy the address from one partner to another
	var p *int       // nil
	updateValueV2(p) // copying nil to p1 pointer
	//fmt.Println(*p)  // this will panic
}

func updateValueV2(p1 *int) {
	var x int = 10
	p1 = &x  // p1 is local to updateValueV2, any change via p1 will not reflect in p in main
	*p1 = 20 // x is updated
	fmt.Println(x)
}
