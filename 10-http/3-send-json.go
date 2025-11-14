package main

import (
	"encoding/json"
	"net/http"
)

// Fields must be exported for json to work
// ` ` back ticks are used to define raw strings
// ` ` don't do any processing on the data
type user struct {
	// use field tags to specify the field name or ignore fields
	FirstName string `json:"first_name"`
	Password  string `json:"-"` // - ignore the field in the output
	Email     string `json:"email"`
}

func main() {

	// register one route to send json

	// start http server

}

func SendJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//
	u := user{
		FirstName: "abc",
		Password:  "xyz",
		Email:     "abc@gmail.com",
	}

	// NewEncoder will encode the struct to json and write the response to the client

	// converting the struct to json
	// send the json response to the client
	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
