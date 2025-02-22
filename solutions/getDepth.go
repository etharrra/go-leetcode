package solutions

import "github.com/etharrra/go-leetcode/utils"

func GetDepth(node *utils.TreeNode) int {
	depth := 0
	if node == nil {
		return depth
	}

	queue := []*utils.TreeNode{node}
	for len(queue) > 0 {
		lenSize := len(queue)
		for i := 0; i < lenSize; i++ {
			current := queue[0]
			queue = queue[1:]
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