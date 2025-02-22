package solutions

import "github.com/etharrra/go-leetcode/utils"

func MinDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		lenSize := len(queue)
		for i := 0; i < lenSize; i++ {
			current := queue[0]
			queue = queue[1:]
			if current.Left == nil && current.Right == nil {
				return depth
			}
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		depth++
	}

	return depth
}
