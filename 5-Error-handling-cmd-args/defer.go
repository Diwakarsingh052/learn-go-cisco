package main

import "fmt"

func main() {
	//defer statements are executed after the function returns
	//1. return
	//2. panic
	//3. function exit or ends

	//defer maintains a stack
	// defer gives guarantee of execution, even if the function panics, it will be executed
	defer fmt.Println(1)
	defer fmt.Println(2)

	fmt.Println(3)

	fmt.Println(4)

}
