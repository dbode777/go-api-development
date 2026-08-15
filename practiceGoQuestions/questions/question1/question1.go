package question1

/*
In Go, write a function named twoLargest that accepts a slice of integers and returns a slice containing the
two largest integers from the input slice.

The returned slice should have the larger integer as the first element and the second largest integer as the second element.
The input slice will always contain at least two integers.

If the largest and the second largest number in the input slice are the same, the output slice should contain two copies of this number.
*/

import (
	"errors"
	"slices"
	"sort"
)

func twoLargest(arr []int) ([]int, error) {
	if len(arr) < 2 {
		return []int{}, errors.New("Array must contain at least two elements")
	}
	sort.IntSlice(arr).Sort()
	slices.Reverse(arr)
	return arr[:2], nil
}
