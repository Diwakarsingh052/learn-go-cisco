package calc

import "fmt"

// Add is an 'exported function'
// first letter is capital,
// making first letter capital makes the function public
func Add(a int, b int) {
	fmt.Println(a + b)
	// display(a + b)

}

// create a new file named as print.go
// create a function named display (unexported)
// instead of using fmt.Println in Add function, use your custom print function

// create a new package named as format
// create a exported function that appends ** to the string
//str = "s" + "**"
