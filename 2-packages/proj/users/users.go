package users

import (
	"fmt"
	"proj/db"
)

func Get() {
	fmt.Println("getting users from", db.Conn)
}
