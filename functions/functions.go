package main

import "fmt"

func doubleNumber(num int) int {
	return num * 2
}

func callFunc(callable func(int) int) int {
	return callable(10)
}

func sums(nums ...int) (s int) {
	for _, num := range nums {
		s += num
	}
	return s
}

func main() {
	sum := sums(1, 2, 3, 4, 5)
	fmt.Println(sum)

	// Uses a slice to pass in a variable number of arguments to calculate the sums
	sum2 := sums([]int{1, 2, 3, 4, 5, 6}...)
	fmt.Println(sum2)

	value := callFunc(doubleNumber)
	fmt.Println(value)
}
