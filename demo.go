package main

import "fmt" // have to use this if you want to execute the program

func main() {
	fmt.Println("hello world")

	fmt.Printf("%d - %s", 'A', "A")
	fmt.Printf("%d - %s", '€', "€")

	arr := [3]int{5}
	fmt.Println("\n", arr)

	s := make([]int, 3)
	s[0] = 1
	s[2] = 3
	fmt.Println(s)

	s1 := []int{1, 2, 3, 4, 5}
	s1 = append(s1[:2], s1[3:]...)
	fmt.Println(s1)
}

// Use go build demo.go in the terminal to create an executable file
// To run the executable file, use ./demo in the terminal
