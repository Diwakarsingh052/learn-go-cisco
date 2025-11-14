package main

import "fmt"

type FileDetails struct {
	name string
}

func (f *FileDetails) Print() {
	fmt.Println(f.name)
}
func (f *FileDetails) Update(name string) {
	f.name = name
}
func main() {

	// try to avoid declaring struct variables as pointers,
	// it can cause nil pointer exceptions, if you forget to initialize the pointer
	//var f *FileDetails // nil
	var f FileDetails // not nil
	f.Update("file.txt")
	f.Print()
	Print(&f)
}

func Print(f *FileDetails) {
	fmt.Println(f.name)
}
