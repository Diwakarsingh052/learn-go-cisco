package main

import (
	"log"
	"os"
)

func main() {
	log.New(os.Stdout, "prefix: ", log.LstdFlags)
	log.New(&os.File{}, "prefix: ", log.LstdFlags)
	log.New(os.Stderr, "prefix: ", log.LstdFlags)
}
