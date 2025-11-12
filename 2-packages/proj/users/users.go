package users

import (
	"fmt"
	"proj/db"
)

func Get() {
	//db.Conn = "postgres"
	db.SetConn("postgres")
	//fmt.Println("getting users from", db.Conn)
	fmt.Println("getting users from", db.GetConn())
}

func Store() {
	//fmt.Println("storing users in ", db.Conn)
	fmt.Println("storing users in ", db.GetConn())
}
