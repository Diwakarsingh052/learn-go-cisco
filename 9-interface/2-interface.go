package main

import "fmt"

type Api interface {
	call()
	// Bigger the interface weaker the abstraction // Rob Pike
	//abc() // all methods of interface must be implemented to implement the interface
}
type api struct {
	url string
}

func (a api) call() {
	fmt.Println("calling api in production")
}

type mockApi struct {
	url string
}

func (a mockApi) call() {
	fmt.Println("calling api mock")
}

func CallApi(api Api) {
	api.call()
}

func main() {
	CallApi(api{"https://api.example.com"})
	CallApi(mockApi{"https://mock.example.com"})
}
