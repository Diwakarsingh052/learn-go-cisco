package main

import "fmt"

// we can't create types on predefined types
// methods can be created on any custom type
type moneyV2 int

func main() {
	var i moneyV2
	i.update()
	fmt.Println(i)
}

func (m *moneyV2) update() {
	var i moneyV2 = 10
	*m = i

}
