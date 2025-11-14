package main

import (
	"proj/db"
)

func main() {
	c, err := db.NewConf("mysql")
	if err != nil {
		panic(err)
	}
	c.AddUsers()
	//log.New()
	//os.OpenFile()
}
