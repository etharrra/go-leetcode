package solutions

import "github.com/etharrra/go-leetcode/utils"

func MaxDepth(root *utils.TreeNode) int {
	if root == nil {
		return 1
	}
	depth := 0

	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		lenSize := len(queue)
		for i := 0; i < lenSize; i++ {
			node := queue[0]
			queue = queue[:1]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}
