package solutions

import "github.com/etharrra/go-leetcode/utils"

func PostorderTraversal(root *utils.TreeNode) []int {
	res := []int{}
	postorder(root, &res)
	return res
}

func postorder(node *utils.TreeNode, res *[]int) {
	if node == nil {
		return
	}
	postorder(node.Left, res)
	postorder(node.Right, res)
	*res = append(*res, node.Val)
}
