package question5

/*
In Go, write a function named sortContacts that accepts a slice of Contact structs and returns a sorted slice of Contact structs.
The contacts should be sorted based on the following criteria:

- Contacts are primarily sorted by the number of groups they belong to, in descending
order.

- If two contacts belong to the same number of groups, they should be sorted by their age,
in ascending order.

- If two contacts belong to the same number of groups and are of the same age, they
should be sorted by their name, in lexicographic (alphabetical) order.

The function signature in Go is: func sortContacts(contacts []Contact) []Contact
*/

import (
	"sort"
)

type Contact struct {
	Name   string
	Email  string
	Age    int
	Groups []string // list of groups the the contact belongs to
}

// sortContacts sorts a slice of Contact structs by the number of groups (descending),
// then by age (ascending), and finally by name (ascending/A-Z).
// It takes a slice of Contact structs as input and returns a new sorted slice
// without modifying the original.
func sortContacts(contacts []Contact) []Contact {
	sort.Slice(contacts, func(i, j int) bool {
		c1, c2 := contacts[i], contacts[j]
		if len(c1.Groups) != len(c2.Groups) {
			return len(c1.Groups) > len(c2.Groups)
		}
		if c1.Age != c2.Age {
			return c1.Age < c2.Age
		}
		return c1.Name < c2.Name
	})

	return contacts
}
