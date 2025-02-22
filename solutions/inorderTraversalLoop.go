package solutions

import "github.com/etharrra/go-leetcode/utils"

func InorderTraversalLoop(root *utils.TreeNode) []int {
	res := []int{}
	stack := []utils.TreeNode{}
	cur := root

	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, *cur)
			cur = cur.Left
		}
		// fmt.Println("stack:", stack)
		cur = &stack[len(stack)-1]   // set pointer to last of stack
		res = append(res, cur.Val)   // append pointer Val to res
		stack = stack[:len(stack)-1] // remove last of stack
		cur = cur.Right              // set pointer to pointer Right
	}
	return res
}
