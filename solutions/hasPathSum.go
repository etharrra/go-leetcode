package solutions

import "github.com/etharrra/go-leetcode/utils"

func HasPathSum(root *utils.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := []struct {
		node *utils.TreeNode
		sum  int
	}{{root, root.Val}}

	for len(stack) > 0 {
		// Pop the top element from the stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := current.node
		currentSum := current.sum

		// Check if it's a leaf node and the path sum matches the target
		if node.Left == nil && node.Right == nil && currentSum == targetSum {
			return true
		}

		// Push the right and left child nodes onto the stack with updated sums
		if node.Right != nil {
			stack = append(stack, struct {
				node *utils.TreeNode
				sum  int
			}{node.Right, currentSum + node.Right.Val})
		}
		if node.Left != nil {
			stack = append(stack, struct {
				node *utils.TreeNode
				sum  int
			}{node.Left, currentSum + node.Left.Val})
		}
	}

	return false
}
