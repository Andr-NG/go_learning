// STRUCTS
// Defining a structure
// Keywords: type and struct

package main

import "fmt"

type Author struct {
	Name string
	Age int
	IsActive bool
}

func main() {

	// Declaring a struct
	type Book struct {
		Title string
		Author string
		Subject string
		BookId int
		Read bool
	}

	// Initialising a struct
	book1 := Book {
		Title: "My_Book",
		Author: "Author_MY",
		Subject: "Novel",
		BookId: 123,
		Read: true,
	}

	// Declaring and initialising a pointer
	ptr := &book1
	
	// Accessing Read via the pointer
	fmt.Println("Book is read:", ptr.Read)

	// Accessing the rest of the fields via the struct itself
	fmt.Println("Title is", book1.Title)
	fmt.Println("Author is", book1.Author)
	fmt.Println("Subject is", book1.Subject)
	fmt.Println("BookID is", book1.BookId)

	// Initialising Author struct

	var Author1 Author
	Author1.Age = 45
	Author1.IsActive = false
	Author1.Name = "Raymond Holt"

	fmt.Println("Author's age is", Author1.Age)
	fmt.Println("Author's name is", Author1.Name)
	fmt.Println("Author's active is", Author1.IsActive)
}