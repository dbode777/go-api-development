package question2

import (
	"fmt"
	"slices"
	"testing"
)

func validation(result []int, expected []int, err error, test *testing.T) {
	if err != nil && len(result) == 0 {
		fmt.Printf("Test failed as expected with error: %s", err)
	} else if err != nil && result == nil {
		test.Errorf("Error occurred: %v", err)
	}

	if !slices.Equal(result, expected) {
		test.Fatalf("Expected result %v, did not equal %v", expected, result)
	}

	fmt.Printf("Test passed: %v, equaled %v\n", result, expected)
}

func TestOddSumMaxPair1(test *testing.T) {
	expected := []int{4, 5}
	result := oddSumMaxPair([]int{1, 2, 4, 3, 5})
	validation(result, expected, nil, test)
}

func TestOddSumMaxPair2(test *testing.T) {
	expected := []int{4, 9}
	result := oddSumMaxPair([]int{7, 3, 9, 4, 2, 1, 5})
	validation(result, expected, nil, test)
}

func TestOddSumMaxPair3(test *testing.T) {
	expected := []int{8, 11}
	result := oddSumMaxPair([]int{11, 8, 7, 5, 3, 1})
	validation(result, expected, nil, test)
}

func TestOddSumMaxPairAllOdd(test *testing.T) {
	expected := []int{}
	result := oddSumMaxPair([]int{1, 1, 1, 1, 1})
	validation(result, expected, nil, test)
}

func TestOddSumMaxPairAllEven(test *testing.T) {
	expected := []int{}
	result := oddSumMaxPair([]int{2, 4, 6, 8, 10})
	validation(result, expected, nil, test)
}

func TestOddSumMaxPairHandlesNegatives(test *testing.T) {
	expected := []int{4, 3}
	result := oddSumMaxPair([]int{-5, 3, -7, 2, -1, 4, -6})
	validation(result, expected, nil, test)
}

func TestOddSumMaxPairCanIncludeZero(test *testing.T) {
	expected := []int{0, 1}
	result := oddSumMaxPair([]int{0, -1, 1, 1, -2})
	validation(result, expected, nil, test)
}

func TestOddSumMaxPairCanIncludeNegatives(test *testing.T) {
	expected := []int{-2, -1}
	result := oddSumMaxPair([]int{-3, -2, -1, -4, -5})
	validation(result, expected, nil, test)
}
