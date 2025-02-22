package utils

import "fmt"

// TreeNode struct definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Helper function to create a binary tree from a slice of integers
// Use 99999 as a placeholder for nil nodes
func CreateTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		// Assign left child
		if i < len(values) && values[i] != 99999 {
			node.Left = &TreeNode{Val: values[i]}
			queue = append(queue, node.Left)
		}
		i++

		// Assign right child
		if i < len(values) && values[i] != 99999 {
			node.Right = &TreeNode{Val: values[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// Helper function to print a tree in level order (breadth-first traversal)
func PrintTree(root *TreeNode) {
	if root == nil {
		fmt.Println("Tree is empty")
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			fmt.Print("nil ")
			continue
		}

		fmt.Printf("%d ", node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	fmt.Println()
}
