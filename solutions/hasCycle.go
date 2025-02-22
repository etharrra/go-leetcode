package solutions

import "github.com/etharrra/go-leetcode/utils"

func HasCycle(head *utils.ListNode) bool {
	visited_nodes := make(map[*utils.ListNode]bool)
	current := head

	for current != nil {
		if visited_nodes[current] {
			return true
		}
		visited_nodes[current] = true
		current = current.Next
	}
	return false
}
