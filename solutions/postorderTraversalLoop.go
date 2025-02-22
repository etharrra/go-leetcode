package solutions

import "github.com/etharrra/go-leetcode/utils"

func PostorderTraversalLoop(root *utils.TreeNode) []int {
	res := []int{}
	stack := []*utils.TreeNode{}
	var prev *utils.TreeNode // To track the previously visited node
	cur := root

	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		// Peek the top node in the stack
		cur = stack[len(stack)-1]

		// Check if the right subtree exists and has not been visited
		if cur.Right == nil || cur.Right == prev {
			res = append(res, cur.Val) // Visit the current node
			stack = stack[:len(stack)-1]
			prev = cur // Mark this node as visited
			cur = nil  // Set current to nil to prevent re-traversal
		} else {
			cur = cur.Right // Move to the right subtree
		}
	}
	return res
}
