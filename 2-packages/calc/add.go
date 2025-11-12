package calc

import "fmt"

// Add is an 'exported function'
// first letter is capital,
// making first letter capital makes the function public
func Add(a int, b int) {
	fmt.Println(a + b)
}
