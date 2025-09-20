package main

import "fmt"

// Use interfaces to provide flexibility to structures and how methods are implemented
type Shape interface {
	getArea() float64
}

// In order for a structure to implement an interface, it must use all functions within an interface
// This allows users to abstract public and private functions. If the function is present in the interface, it would be considered public.
type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) getArea() float64 {
	return r.width * r.height
}

func main() {
	var s Shape = Rectangle{width: 10.0, height: 20.0}
	fmt.Printf("Area: %.2f\n", s.getArea())
}
