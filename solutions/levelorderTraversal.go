package solutions

import "github.com/etharrra/go-leetcode/utils"

func levelorderTraversal(root *utils.TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	current := root
	queue := []*utils.TreeNode{current}

	for len(queue) > 0 {
		current = queue[0]
		queue = queue[1:]
		if current == nil {
			res = append(res, 99999)
			continue
		}

		res = append(res, current.Val)
		queue = append(queue, current.Left)
		queue = append(queue, current.Right)
	}

	return res
}
