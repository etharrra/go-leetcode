package solutions

import "github.com/etharrra/go-leetcode/utils"

func PreorderTraversal(root *utils.TreeNode) []int {
	res := []int{}
	preorder(root, &res)
	return res
}

func preorder(node *utils.TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	preorder(node.Left, res)
	preorder(node.Right, res)
}
