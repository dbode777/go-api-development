package question6

/*
In Go, write a function named flattenBST that accepts a pointer to the
root node of a binary search tree (BST) and two integer values min and max.
The function should return a slice of integers representing a flattened,
inorder traversal of the binary search tree,
but only including the nodes whose values fall within the inclusive range min to max.

A tree is "flattened" by performing an inorder traversal,
which visits the left subtree, the root node, and then the right subtree.

The function should return a slice that contains the values of the nodes in the order they were visited.

The function signature in Go is: func flattenBST(root *TreeNode, min int, max int) []int
*/

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func flattenBST(root *TreeNode, min int, max int) []int {
	flattenedTree := []int{}
	if root == nil {
		fmt.Println("Tree is empty")
		return flattenedTree
	}

	if root.Value < min || root.Value > max {
		fmt.Printf("Value %d is outside of the range %d to %d\n", root.Value, min, max)
		return []int{}
	}

	flattenedTree = append(flattenedTree, flattenBST(root.Left, min, max)...)
	flattenedTree = append(flattenedTree, root.Value)
	flattenedTree = append(flattenedTree, flattenBST(root.Right, min, max)...)
	return flattenedTree
}
