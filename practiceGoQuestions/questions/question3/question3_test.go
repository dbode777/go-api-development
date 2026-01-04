package question3

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func validation(result map[string]int, expected map[string]int, err error, test *testing.T) {
	if err != nil && len(result) == 0 {
		fmt.Printf("Test failed as expected with error: %s", err)
	} else if err != nil && result == nil {
		test.Errorf("Error occurred: %v", err)
	}

	if diff := cmp.Diff(result, expected); diff != "" {
		test.Fatalf("Diff: %s\n", diff)
	}

	fmt.Printf("Test passed: %v, equaled %v\n", result, expected)
}

func TestTopAuthors1(test *testing.T) {
	expected := map[string]int{
		"John Doe":         12500,
		"Michael Crichton": 20000,
		"Stephen King":     33000,
	}

	books := []Book{
		{
			Title:  "The Great Escape",
			Author: "John Doe",
			Sales:  4500,
		}, {
			Title:  "The Lost World",
			Author: "Michael Crichton",
			Sales:  20000,
		}, {
			Title:  "The Final Countdown",
			Author: "John Doe",
			Sales:  8000,
		}, {
			Title:  "The Tale of Two Cities",
			Author: "Charles Dickens",
			Sales:  5000,
		}, {
			Title:  "Under the Dome",
			Author: "Stephen King",
			Sales:  25000,
		}, {
			Title:  "Pet Sematary",
			Author: "Stephen King",
			Sales:  8000,
		}}
	result := topAuthors(books)
	validation(result, expected, nil, test)
}

func TestTopAuthors2(test *testing.T) {
	expected := map[string]int{
		"Author 1": 13000,
		"Author 3": 13000,
		"Author 4": 15000,
	}

	books := []Book{
		{
			Title:  "Book 1",
			Author: "Author 1",
			Sales:  8000,
		}, {
			Title:  "Book 2",
			Author: "Author 2",
			Sales:  6000,
		}, {
			Title:  "Book 3",
			Author: "Author 3",
			Sales:  12000,
		}, {
			Title:  "Book 4",
			Author: "Author 1",
			Sales:  5000,
		}, {
			Title:  "Book 5",
			Author: "Author 3",
			Sales:  1000,
		}, {
			Title:  "Book 6",
			Author: "Author 4",
			Sales:  15000,
		}}
	result := topAuthors(books)
	validation(result, expected, nil, test)
}

func TestTopAuthors3(test *testing.T) {
	expected := map[string]int{
		"Author 3": 11000,
	}

	books := []Book{
		{
			Title:  "Book 1",
			Author: "Author 1",
			Sales:  3000,
		}, {
			Title:  "Book 2",
			Author: "Author 3",
			Sales:  4000,
		}, {
			Title:  "Book 3",
			Author: "Author 1",
			Sales:  2000,
		}, {
			Title:  "Book 4",
			Author: "Author 3",
			Sales:  7000,
		}}
	result := topAuthors(books)
	validation(result, expected, nil, test)
}

func TestTopAuthorsShouldReturnEmptyIfNoAuthorsMeetCriteria(test *testing.T) {
	expected := map[string]int{}

	books := []Book{
		{
			Title:  "Book 1",
			Author: "Author 1",
			Sales:  5000,
		}, {
			Title:  "Book 2",
			Author: "Author 2",
			Sales:  2000,
		}}
	result := topAuthors(books)
	validation(result, expected, nil, test)
}

func TestTopAuthorsHandlesEmptyInputs(test *testing.T) {
	expected := map[string]int{}

	books := []Book{}
	result := topAuthors(books)
	validation(result, expected, nil, test)
}
