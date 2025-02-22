package solutions

import "github.com/etharrra/go-leetcode/utils"

func isBalanced(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}

	diff := GetDepth(root.Left) - GetDepth(root.Right)
	if diff > -2 && diff < 2 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}
