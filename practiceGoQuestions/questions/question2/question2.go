package question2

/*
In Go, write a function named oddSumMaxPair that accepts a slice of integers and returns a slice containing the pair of integers
that have the highest sum among all pairs with odd sums.
The input slice will always contain at least two integers.
The pairs should be considered as (numbers[i], numbers[j]) where j>i.
If multiple pairs have the same highest odd sum, return the pair that occurs first.

The returned slice should have the first element of the pair as the first element and the second element of the pair as the second element.
The sum of the pair should be an odd number. Remember to handle the edge case where there are no pairs with an odd sum.
In this case, the function should return a nil or empty slice.
*/

import (
	"slices"
	"sort"
)

func oddSumMaxPair(arr []int) []int {
	oddNumbers := []int{}
	evenNumbers := []int{}

	for _, num := range arr {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		} else {
			oddNumbers = append(oddNumbers, num)
		}
	}

	// Skip evalution of sorting if an odd sum is not possible
	if len(evenNumbers) == 0 || len(oddNumbers) == 0 {
		return []int{}
	}

	sort.IntSlice(evenNumbers).Sort()
	slices.Reverse(evenNumbers)
	sort.IntSlice(oddNumbers).Sort()
	slices.Reverse(oddNumbers)

	return []int{evenNumbers[0], oddNumbers[0]}
}
