package solutions

import "github.com/etharrra/go-leetcode/utils"

func PreorderTraversalLoop(root *utils.TreeNode) []int {
	res := []int{}
	stack := []utils.TreeNode{}
	cur := root

	for len(stack) != 0 || cur != nil {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, *cur)
			cur = cur.Left
		}
		// fmt.Println("stack:", stack)
		cur = &stack[len(stack)-1]
		cur = cur.Right
		stack = stack[:len(stack)-1]
	}

	return res
}
