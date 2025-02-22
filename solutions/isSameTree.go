package solutions

import (
	"fmt"
	"slices"

	"github.com/etharrra/go-leetcode/utils"
)

func IsSameTree(p *utils.TreeNode, q *utils.TreeNode) bool {
	p_arr := levelorderTraversal(p)
	q_arr := levelorderTraversal(q)
	fmt.Println(p_arr)
	fmt.Println(q_arr)

	return slices.Compare(p_arr, q_arr) == 0
}
