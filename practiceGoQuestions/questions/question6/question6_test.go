package question6

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func validation(result []int, expected []int, err error, test *testing.T) {
	if err != nil && len(result) == 0 {
		fmt.Printf("Test failed as expected with error: %s", err)
	} else if err != nil && result == nil {
		test.Errorf("Error occurred: %v", err)
	}

	if diff := cmp.Diff(result, expected); diff != "" {
		test.Fatalf("Diff: %s\n", diff)
	}

	fmt.Printf("Test passed: %v, equaled %v\n", result, expected)
}

func TestFlattenBST1(t *testing.T) {
	expected := []int{10, 15, 20, 30}
	root := &TreeNode{
		Value: 20,
		Left: &TreeNode{
			Value: 10,
			Left: &TreeNode{
				Value: 5,
			},
			Right: &TreeNode{
				Value: 15,
			},
		},
		Right: &TreeNode{
			Value: 30,
			Right: &TreeNode{
				Value: 35,
			},
		},
	}

	min := 10
	max := 30

	result := flattenBST(root, min, max)
	validation(result, expected, nil, t)
}

func TestFlattenBST2(t *testing.T) {
	expected := []int{30, 40, 45, 50, 60, 70}
	root := &TreeNode{
		Value: 50,
		Left: &TreeNode{
			Value: 30,
			Left: &TreeNode{
				Value: 20,
				Left: &TreeNode{
					Value: 10,
				},
			},
			Right: &TreeNode{
				Value: 40,
				Right: &TreeNode{
					Value: 45,
				},
			},
		},
		Right: &TreeNode{
			Value: 70,
			Left: &TreeNode{
				Value: 60,
			},
			Right: &TreeNode{
				Value: 80,
			},
		},
	}

	min := 30
	max := 70

	result := flattenBST(root, min, max)
	validation(result, expected, nil, t)
}

func TestFlattenBST3(t *testing.T) {
	expected := []int{20, 25, 30, 40, 50, 51, 55, 60, 70}
	root := &TreeNode{
		Value: 50,
		Left: &TreeNode{
			Value: 30,
			Left: &TreeNode{
				Value: 20,
				Left: &TreeNode{
					Value: 10,
				},
				Right: &TreeNode{
					Value: 25,
				},
			},
			Right: &TreeNode{
				Value: 40,
			},
		},
		Right: &TreeNode{
			Value: 70,
			Left: &TreeNode{
				Value: 55,
				Left: &TreeNode{
					Value: 51,
				},
				Right: &TreeNode{
					Value: 60,
				},
			},
			Right: &TreeNode{
				Value: 80,
				Left: &TreeNode{
					Value: 75,
				},
				Right: &TreeNode{
					Value: 90,
				},
			},
		},
	}

	min := 20
	max := 70

	result := flattenBST(root, min, max)
	validation(result, expected, nil, t)
}
