package main

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

func SendJSON() {
	w.Header().Set("Content-Type", "application/json")

	// create a variable of type user

	// convert the struct to json, json.Marshal()

	// send the json using w.Write()
}
