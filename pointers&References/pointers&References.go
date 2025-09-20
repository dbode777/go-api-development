package main

import "fmt"

func changeValue(num *int) {
	*num = 20
}

type Book struct {
	author string
	title  string
}

func (b *Book) setTitle(title string) {
	b.title = title
}

func main() {
	a := 10
	b := &a         // The & is used to get the memory address
	fmt.Println(a)  // will be 10
	fmt.Println(*b) // will be 10 as it will print the value at the memory address
	*b = 20         // The * is a derefence operator
	fmt.Println(a)  // will be 20 as the memory address now points to 20

	// Use pointers to change values of variables within separate functions
	changeValue(&a)
	fmt.Println(a)

	// This is also how you would update fields in a structure, i.e. similar to a setter in Java/Python
	book := Book{author: "Jimothy Timbers", title: "Odyssey"}
	book.setTitle("The Illiad")
	fmt.Println(book)
}
