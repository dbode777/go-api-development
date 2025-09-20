package main

import (
	errorhandler "example/GO-API-TUTORIAL/errorHandler"
	structs "example/GO-API-TUTORIAL/structs"
	validator "example/GO-API-TUTORIAL/validator"

	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

var books = []structs.Book{
	{ID: "1", Title: "Book 1", Author: "Author 1", Quantity: 1, Price: 9.99},
	{ID: "2", Title: "Book 2", Author: "Author 2", Quantity: 5, Price: 14.99},
	{ID: "3", Title: "Book 3", Author: "Author 3", Quantity: 15, Price: 19.99},
}

func getBooks(c *gin.Context) {
	// This will return a JSON response object when the GET request is made.
	// Can also return files, html, etc.
	c.IndentedJSON(http.StatusOK, books)
}

func lookupBook(c *gin.Context, isCheckingOut bool) *structs.Book {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": errorhandler.ErrBookIdMissing})
		return nil
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": errorhandler.ErrBookNotFound, "error": err.Error()})
		return nil
	}

	if isCheckingOut && book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": errorhandler.ErrBookCheckedOut, "error": "error occurred when attempting to checkout a book"})
		return nil
	} else if isCheckingOut && book.Quantity > 0 {
		book.Quantity -= 1
		return book
	}

	// Returning book or adding a book to the stock
	book.Quantity += 1
	return book
}

func checkoutBook(c *gin.Context) {
	var book *structs.Book = lookupBook(c, true) // true if checking out

	fmt.Printf("%s by %s has successfully been checked out\n", book.Title, book.Author)
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	var book *structs.Book = lookupBook(c, false) // false if returning

	fmt.Printf("%s by %s has successfully been returned\n", book.Title, book.Author)
	c.IndentedJSON(http.StatusOK, book)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": errorhandler.ErrBookNotFound, "error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*structs.Book, error) {
	fmt.Println("Getting book by id: " + id)
	for index, book := range books {
		if book.ID == id {
			return &books[index], nil
		}
	}

	return nil, errorhandler.NewError(errorhandler.ErrBookNotFound)
}

func createBook(c *gin.Context) {
	var newBook structs.Book

	// Call BindJSON to bind the received JSON to the pointer of newBook
	if err := c.BindJSON(&newBook); err != nil {
		return // will automatically return an error response to the client via BindJSON
	}

	// Conduct validation on the newBook
	validator.ValidateStruct(newBook)

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook) // http.StatusCreated is 201 and means the object was created successfully
}

func updateBook(c *gin.Context) {
	var updateBook structs.Book

	if err := c.BindJSON(&updateBook); err != nil {
		return
	}

	validator.ValidateStruct(updateBook)

	for index, book := range books {
		if book.ID == updateBook.ID {
			books[index] = updateBook
			fmt.Printf("Book with id %s has been updated\n", updateBook.ID)
			c.IndentedJSON(http.StatusOK, updateBook)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": errorhandler.ErrBookNotFound, "error": "error occurred when attempting to update a book"})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")

	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			fmt.Printf("Book with id %s has been deleted\n", id)
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": errorhandler.ErrBookNotFound, "error": "error occurred when attempting to delete a book"})
}

// Use go run main.go in the terminal to run the program, and then in another terminal use curl localhost:8080/{endpoint} to make requests

func main() {
	for _, book := range books {
		validator.ValidateStruct(book)
	}
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PUT("/checkout", checkoutBook)
	router.PUT("/return", returnBook)
	router.PUT("/books", updateBook)
	router.DELETE("/books/:id", deleteBook)
	router.Run("localhost:8080")
}
