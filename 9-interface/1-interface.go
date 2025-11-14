package main

import "fmt"

// Reader is an interface
// Interface is automatically implemented by the struct
// if the struct has all the methods of the interface

type Reader interface {
	Read(b []byte) (int, error)
}
type File struct {
	Name string
}

func (f File) Read(b []byte) (int, error) {
	fmt.Println("reading files and processing file", f.Name)
	return 0, nil
}

type IO struct {
	name string
}

func (i IO) Read(b []byte) (int, error) {
	fmt.Println("reading and processing io ", i.name)
	return 0, nil
}

//	func ReadData(f File, i IO) {
//		f.Read(nil)
//		i.Read(nil)
//	}

func ReadData(r Reader) {
}
func main() {

	f := File{"test.txt"}
	i := IO{name: "io"}
	ReadData(f)
	ReadData(i)

}
