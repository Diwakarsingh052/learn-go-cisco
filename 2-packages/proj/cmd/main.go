package main

import (
	"proj/db"
	"proj/users"
)

// go run . // to compile and run all the package main files
func main() {
	//calc.Add(1, 2)
	//setup()
	db.SetConn("mysql")
	users.Get()
	users.Store()

}
