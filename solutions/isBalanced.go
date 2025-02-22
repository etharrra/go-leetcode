package solutions

import "github.com/etharrra/go-leetcode/utils"

func IsBalanced(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}

	diff := GetDepth(root.Left) - GetDepth(root.Right)
	if diff > -2 && diff < 2 {
		return IsBalanced(root.Left) && IsBalanced(root.Right)
	}
	return false
}
