package main

import "net/http"

func main() {
	http.HandleFunc("/home", home)
	// starts the server on the specified port
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	//ResponseWriter
	// it is used to send the response back to the client
	// it also use to write headers and status code
	w.Write([]byte("Hello World"))

	//Request
	// It contains all the request data that our service received from the client
	// like body, headers, url, params etc.
}
