package solutions

import "github.com/etharrra/go-leetcode/utils"

func InorderTraversal(root *utils.TreeNode) []int {
	res := []int{}
	inorder(root, &res)
	return res
}

func inorder(node *utils.TreeNode, res *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, res)
	*res = append(*res, node.Val)
	inorder(node.Right, res)
}
