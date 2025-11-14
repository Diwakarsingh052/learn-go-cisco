package main

import "fmt"

type User struct {
	name  string
	email string
}

func (u *User) UpdateEmail(email string) {
	u.email = email
}

type BookAuthor struct {
	// User is embedded
	// composing a new struct with an existing struct
	// User is a type here, field name would be based on the type name
	User     // anonymous field // field without a name
	BookName string
}

func (b *BookAuthor) UpdateBookName(name string) {
	b.BookName = name
}

func (b *BookAuthor) Print() {
	fmt.Printf("%+v\n", b)
}

type MovieAuthor struct {
	// u is a field name of type User
	u         User // not embedding
	MovieName string
}

func (m *MovieAuthor) Print() {
	fmt.Printf("%+v\n", m)
}

func main() {

	u := User{"John", "john@exy.com"}
	b := BookAuthor{
		User:     u,
		BookName: "Harry Potter",
	}
	m := MovieAuthor{
		u:         u,
		MovieName: "Inception",
	}
	// we can call the user methods directly on the bookauthor
	// because use is embedded
	b.UpdateEmail("john@gmail.com")
	b.Print()

	m.u.UpdateEmail("john@xyz.com")
	m.Print()

}
