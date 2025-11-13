package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

// error variables should start with the word Err
// errors.Is , errors.As

var ErrInvalidArgs = errors.New("invalid arguments")

func main() {
	//fmt.Println(os.Args)
	//var err error // nil // nil error means no error
	err := greet()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("End of the main")
}

func greet() error {

	// getting rid of the program name from the args and constructing a new slice
	cmdArgs := os.Args[1:]
	if len(cmdArgs) != 3 {

		//return errors.New("invalid arguments")
		//return ErrInvalidArgs
		return fmt.Errorf("invalid arguments %v", cmdArgs)
	}
	name := cmdArgs[0]
	ageString := cmdArgs[1]
	marksString := cmdArgs[2]

	// Atoi converts string to int
	// whenever a func returns an error, next step must be error handling
	age, err := strconv.Atoi(ageString)
	if err != nil {
		//fmt.Println(err)
		//return // simply write return if no return values in the function
		//return err

	}

	marks, err := strconv.Atoi(marksString)
	if err != nil {
		return fmt.Errorf("error converting marks to int %v", err)
	}
	fmt.Println("Hello", name, "you are", age, "years old and you have", marks, "marks")

	return nil
}
