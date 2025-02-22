package solutions

import "github.com/etharrra/go-leetcode/utils"

func DeleteDuplicates(head *utils.ListNode) *utils.ListNode {
	headTemp := head
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			// head.Next = nextUniqueNode(head.Next, head.Val)
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return headTemp
}
