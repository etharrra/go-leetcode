package solutions

import "github.com/etharrra/go-leetcode/utils"

func IsSymmetric(root *utils.TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	left := root.Left
	right := root.Right
	var checkSymmetric func(p *utils.TreeNode, q *utils.TreeNode) bool

	checkSymmetric = func(p *utils.TreeNode, q *utils.TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if p == nil || q == nil || p.Val != q.Val {
			return false
		}

		return checkSymmetric(p.Left, q.Right) && checkSymmetric(p.Right, q.Left)
	}
	return checkSymmetric(left, right)
}
