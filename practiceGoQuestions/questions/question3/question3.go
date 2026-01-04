package question3

/*
In Go, write a function named topAuthors that accepts a slice of Book structs and returns a map[string]int where the keys
are the author names and the values are the sum of the sales of all their books.
The function should only include authors in the result map if the total sales of all their books are at least 10000.
*/

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Sales  int
}

func topAuthors(books []Book) map[string]int {
	authorSales := make(map[string]int)
	for _, book := range books {
		authorSales[book.Author] += book.Sales
	}

	result := make(map[string]int)
	for author, sales := range authorSales {
		if sales >= 10000 {
			result[author] = sales
		} else {
			fmt.Printf("Skipping author %s as total sales is less than 10000\n", author)
		}
	}
	return result
}
