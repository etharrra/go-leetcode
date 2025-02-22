package solutions

import "github.com/etharrra/go-leetcode/utils"

func GetIntersectionNode(headA, headB *utils.ListNode) *utils.ListNode {
	visitedNode := make(map[*utils.ListNode]bool)
	currentA := headA
	currentB := headB

	for currentA != nil {
		visitedNode[currentA] = true
		currentA = currentA.Next
	}

	for currentB != nil {
		if _, foundB := visitedNode[currentB]; foundB {
			return currentB
		}
		currentB = currentB.Next
	}

	return nil
}
