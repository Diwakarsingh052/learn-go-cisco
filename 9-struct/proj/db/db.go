package db

import (
	"errors"
	"fmt"
)

// var Conn string
// Conf struct is exported, it can be used outside the package
// dbConn is not exported, which means it can only be used inside the package
// and it can't  be changed by other packages

type Conf struct {
	dbConn string
}

func NewConf(conn string) (Conf, error) {
	if conn == "" {
		return Conf{}, errors.New("db connection string is empty")
	}
	return Conf{dbConn: conn}, nil
}

func (c *Conf) AddUsers() {
	fmt.Println("add users", c.dbConn)
}
