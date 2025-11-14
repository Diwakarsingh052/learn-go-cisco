package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/double", doubleHandler)
	http.ListenAndServe(":8080", nil)
}

// http:localhost:8080/double?v=2

func doubleHandler(w http.ResponseWriter, r *http.Request) {

	// get the query parameter
	text := r.URL.Query().Get("v")
	if text == "" {
		http.Error(w, "missing query parameter", http.StatusBadRequest)
		return
	}

	v, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "double of number:", v, ":", v*2)

}
