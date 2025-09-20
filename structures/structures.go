package main

import "fmt"

type Sport struct {
	name     string
	position string
}
type Person struct {
	// Capitalize the property names to make them public, lowercase to make them private
	name     string
	age      uint
	f        func(int, int) int
	favSport Sport   // Embedding a struct within a struct
	sports   []Sport // Embedding a slice of a struct within a struct
}

// This is how to define a method that's usable by all instances of the structure
// Capitalize functions namnes to make them public, lowercase to make them private
func (p Person) getPropertiesString() string {
	return fmt.Sprintf("%s is %d years old", p.name, p.age)
}

func (p Person) getName() string {
	return p.name
}

func (p Person) setName(newName string) {
	p.name = newName
	fmt.Println(p.name)
}

func main() {
	/* The order these properties do not matter when the properties are specified.
	If they are not specified, order does matter.*/
	person := Person{
		name:     "Dalton",
		age:      30,
		favSport: Sport{name: "Basketball", position: "Point Guard"},
		sports: []Sport{
			{name: "Basketball", position: "Point Guard"},
			{name: "Football", position: "Quarterback"},
		},
	}
	fmt.Println(person)

	// Can specify a function as a property and define it for a single instance of the structure
	// Useful if you want to use the same function structure for multiple instances that perform different actions
	person.f = func(a, b int) int {
		return a + b
	}

	fmt.Println(person.f(1, 2))

	properties := person.getPropertiesString()
	fmt.Println(properties)

	fmt.Println(person.getName())
	person.setName("John") // Will print "John", doesn't change the name in the structure
}
