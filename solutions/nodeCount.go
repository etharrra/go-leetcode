package solutions

import (
	"fmt"

	"github.com/etharrra/go-leetcode/utils"
)

func NodeCount(node *utils.TreeNode) int {
	res := []int{}
	if node == nil {
		return len(res)
	}
	queue := []*utils.TreeNode{node}
	firstNull := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			res = append(res, 999)
			if firstNull == 0 {
				firstNull = len(res) - 1
			}
			continue
		}
		firstNull = 0

		res = append(res, node.Val)
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}

	fmt.Println("firstNull:", firstNull)
	fmt.Println("res:", res)
	res = res[:firstNull]
	fmt.Println("pure res:", res)
	fmt.Println("len res:", len(res))

	return len(res)
}
