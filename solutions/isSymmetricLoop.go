package solutions

import "github.com/etharrra/go-leetcode/utils"

func IsSymmetricLoop(root *utils.TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}

	resBool := true
	queueLeft := []*utils.TreeNode{root.Left}
	queueRight := []*utils.TreeNode{root.Right}

	for len(queueLeft) > 0 && len(queueRight) > 0 {
		leftNode := queueLeft[0]
		rightNode := queueRight[0]

		queueLeft = queueLeft[1:]
		queueRight = queueRight[1:]

		if leftNode == nil && rightNode == nil {
			continue
		}

		if leftNode == nil || rightNode == nil || leftNode.Val != rightNode.Val {
			resBool = false
			break
		}

		queueLeft = append(queueLeft, leftNode.Left)
		queueLeft = append(queueLeft, leftNode.Right)

		queueRight = append(queueRight, rightNode.Right)
		queueRight = append(queueRight, rightNode.Left)
	}

	return resBool
}
