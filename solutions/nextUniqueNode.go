package solutions

import "github.com/etharrra/go-leetcode/utils"

func NextUniqueNode(head *utils.ListNode, val int) *utils.ListNode {
	for head != nil {
		if head.Val != val {
			return head
		}
		head = head.Next
	}
	return nil
}
