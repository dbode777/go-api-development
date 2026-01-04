package question1

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

func TestTwoLargest1(test *testing.T) {
	expected := []int{5, 4}
	result, error := twoLargest([]int{1, 2, 4, 3, 5})
	validation(result, expected, error, test)
}

func TestTwoLargest2(test *testing.T) {
	expected := []int{9, 8}
	result, error := twoLargest([]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0})
	validation(result, expected, error, test)
}

func TestTwoLargest3(test *testing.T) {
	expected := []int{101, 101}
	result, error := twoLargest([]int{100, 101, 100, 100, 101})
	validation(result, expected, error, test)
}

func TestTwoLargest4(test *testing.T) {
	expected := []int{5, 5}
	result, error := twoLargest([]int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1, 1, 2, 3, 4, 5, 5, 4, 3, 2, 1})
	validation(result, expected, error, test)
}

func TestTwoLargest5(test *testing.T) {
	expected := []int{3, 2}
	result, error := twoLargest([]int{2, 3})
	validation(result, expected, error, test)
}

func TestTwoLargest6(test *testing.T) {
	expected := []int{95, 89}
	result, error := twoLargest([]int{42, 13, 57, 81, 29, 66, 38, 95, 72, 20, 89, 10, 47, 63, 55, 12, 86, 51, 77, 32})
	validation(result, expected, error, test)
}

func TestTwoLargest7(test *testing.T) {
	expected := []int{98, 90}
	result, error := twoLargest([]int{15, 61, 25, 37, 52, 10, 47, 32, 73, 41, 88, 12, 64, 56, 29, 83, 98, 70, 90, 21})
	validation(result, expected, error, test)
}

func TestTwoLargestEmptyArray(test *testing.T) {
	expected := []int{}
	result, error := twoLargest([]int{})
	validation(result, expected, error, test)
}
