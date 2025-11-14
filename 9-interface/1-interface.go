package main

import (
	"fmt"
	"io"
)

// Reader is an interface
// Interface is automatically implemented by the struct
// if the struct has all the methods of the interface

// Polymorphism means that a piece of code changes its behavior depending on the
// concrete data it’s operating on // Tom Kurtz, Basic inventor

// "Don’t design with interfaces, discover them". - Rob Pike

// with interface we can switch between multiple implementations from a single variable of interface

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

// // ReadData accepts any type that implements Reader interface
func ReadData(r io.Reader) {
	fmt.Printf("%T\n", r)
	r.Read(nil)

}

func main() {

	f := File{"test.txt"}
	i := IO{name: "io"}
	ReadData(f)
	ReadData(i)

}
