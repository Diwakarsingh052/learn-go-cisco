package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// http in go by default serves the request in a separate goroutine
	// it is concurrent by default
	//panic("panic in main")
	http.HandleFunc("/hello", homeV2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func homeV2(w http.ResponseWriter, r *http.Request) {
	// http layer would automatically recover from the panic if it happens within a request
	//panic("panic in homeV2")
	wg := new(sync.WaitGroup)
	wg.Go(func() {
		// when panic happens in a goroutine, it would be recovered automatically
		defer recoverPanic()
		// panic in a manual spawned goroutine will crash the server
		// it would not recover from the panic automatically
		panic("panic in homeV2")
	})
	w.Write([]byte("Hello World"))
}

func recoverPanic() {
	// msg is the panic message
	// if nil then there was no panic
	msg := recover()
	if msg != nil {
		fmt.Println("recovered from panic", msg)
		return
	}
}
