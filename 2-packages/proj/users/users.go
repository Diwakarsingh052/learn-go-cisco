package users

import (
	"fmt"
	"proj/db"
)

func Get() {
	db.Conn = "postgres"
	fmt.Println("getting users from", db.Conn)
}

func Store() {
	fmt.Println("storing users in ", db.Conn)
}
