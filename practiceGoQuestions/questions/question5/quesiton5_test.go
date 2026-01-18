package question5

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func validation(result []Contact, expected []Contact, err error, test *testing.T) {
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

func TestSortContacts(t *testing.T) {
	expected := []Contact{
		{
			Name:   "Jim Doe",
			Email:  "jimdoe@example.com",
			Age:    28,
			Groups: []string{"Book Club", "Moderators"},
		},
		{
			Name:   "Alice Johnson",
			Email:  "alicejohnson@example.com",
			Age:    30,
			Groups: []string{"Friends", "Work"},
		},
		{
			Name:   "John Doe",
			Email:  "johndoe@example.com",
			Age:    30,
			Groups: []string{"Friends", "Work"},
		},
		{
			Name:   "Jane Doe",
			Email:  "janedoe@example.com",
			Age:    25,
			Groups: []string{"Work"},
		},
	}

	contacts := []Contact{
		{
			Name:   "John Doe",
			Email:  "johndoe@example.com",
			Age:    30,
			Groups: []string{"Friends", "Work"},
		},
		{
			Name:   "Jane Doe",
			Email:  "janedoe@example.com",
			Age:    25,
			Groups: []string{"Work"},
		},
		{
			Name:   "Jim Doe",
			Email:  "jimdoe@example.com",
			Age:    28,
			Groups: []string{"Book Club", "Moderators"},
		},
		{
			Name:   "Alice Johnson",
			Email:  "alicejohnson@example.com",
			Age:    30,
			Groups: []string{"Friends", "Work"},
		},
	}

	result := sortContacts(contacts)
	validation(result, expected, nil, t)
}

func TestSortContactsHandleEmptyContacts(t *testing.T) {
	expected := []Contact{}

	contacts := []Contact{}

	result := sortContacts(contacts)
	validation(result, expected, nil, t)
}
