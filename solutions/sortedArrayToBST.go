package solutions

import "github.com/etharrra/go-leetcode/utils"

func SortedArrayToBST(nums []int) *utils.TreeNode {
	if len(nums) == 1 {
		return NewTreeNode(nums[0])
	}

	var createNode func(node *utils.TreeNode, left_arr []int, right_arr []int)

	mid := len(nums) / 2
	head := NewTreeNode(nums[mid])
	tmp_head := head
	left_arr := nums[:mid]
	right_arr := nums[mid+1:]

	createNode = func(node *utils.TreeNode, left_arr, right_arr []int) {
		if len(left_arr) == 0 && len(right_arr) == 0 {
			return
		}

		if len(left_arr) > 0 {
			mid = len(left_arr) / 2
			node.Left = NewTreeNode(left_arr[mid])
			createNode(node.Left, left_arr[:mid], left_arr[mid+1:])
		}

		if len(right_arr) > 0 {
			mid = len(right_arr) / 2
			node.Right = NewTreeNode(right_arr[len(right_arr)/2])
			createNode(node.Right, right_arr[:mid], right_arr[mid+1:])
		}
	}

	createNode(tmp_head, left_arr, right_arr)

	return head
}

func NewTreeNode(val int) *utils.TreeNode {
	return &utils.TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
