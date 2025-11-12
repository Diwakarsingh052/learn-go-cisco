package db

// this package is responsible for database connection and its management

// Conn is exported Global variable
// if any function changes its value then it would be changed for all

// exported global variables should be avoided to store connection states

//var Conn = "mysql"

var conn string

func SetConn(c string) {
	conn = c
}

func GetConn() string {
	return conn
}
