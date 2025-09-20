package main

import (
	"fmt"
)

// you can define generic structures with loosely defined types
type GenericPersonStruct[T any] struct {
	Name T
	Age  T
}

type GenericSlice[T any] []T

/*
Generics can be used to flexibly enforce type constraints on functions for variables that can be of multiple types.

Use the | operator to specify multiple types. Use a , to specify multiple constraint types.
*/
func add[T int | float64, U int](a T, b T) T {
	return a + b
}

/* You can use a generic type constraint like comparable or any if the type can be flexibly defined*/
func getValues[Key comparable, Value any](mp map[Key]Value) []Value {
	values := []Value{}
	for _, v := range mp {
		values = append(values, v)
	}

	return values
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1.0, 2.0))

	mp := map[string]uint{"a": 1, "b": 300, "c": 400}
	// Maps are randomly ordered, so the answer can be different each time
	fmt.Println(getValues(mp))
}
