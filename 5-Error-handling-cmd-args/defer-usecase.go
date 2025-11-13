package main

import (
	"log"
	"os"
)

func main() {
	//f, err := openFile()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	f, err := os.Open("abc.txt")
	if err != nil {
		log.Println(err)
		return
	}
	// when the function returns, f.Close() will be called
	defer f.Close()

	// structure of program
	// first handles the error and then registers the defer statement

	//if someCondition {
	//	return
	//	// defer would be called
	//	// i don't have to write f.Close() here
	//}

	// some panic happens below
	// file would be closed automatically
}

func openFile() (*os.File, error) {
	f, err := os.Open("abc.txt")
	if err != nil {
		return nil, err
	}
	// here defer f.Close should not be registered
	// it will close the file immediately after the function returns
	// and the end use would not be able to work with the file
	//defer f.Close()
	return f, nil

}
