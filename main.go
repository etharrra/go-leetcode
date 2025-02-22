package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/etharrra/go-leetcode/solutions"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	s = "race a car"
	s = " "
	nums := []int{}
	fmt.Println(solutions.IsPalindrome(s))
	fmt.Println(solutions.LongestConsecutive(nums))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 1
	}
	depth := 0

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		lenSize := len(queue)
		for i := 0; i < lenSize; i++ {
			node := queue[0]
			queue = queue[:1]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}

func _maxDepth(root *TreeNode) int {
	if root == nil {
		return 1
	}
	res := nodeCount(root)

	depth := 0
	for i := 0; i < res; i++ {
		depth++
		pow := math.Pow(2, float64(i))
		powPlus := math.Pow(2, float64(i+1))
		if res >= int(pow) && res < int(powPlus) {
			break
		}
	}
	return depth
}

func nodeCount(node *TreeNode) int {
	res := []int{}
	if node == nil {
		return len(res)
	}
	queue := []*TreeNode{node}
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

func isSymmetricLoop(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}

	resBool := true
	queueLeft := []*TreeNode{root.Left}
	queueRight := []*TreeNode{root.Right}

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

func isSymmetric(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	left := root.Left
	right := root.Right
	var checkSymmetric func(p *TreeNode, q *TreeNode) bool

	checkSymmetric = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if p == nil || q == nil || p.Val != q.Val {
			return false
		}

		return checkSymmetric(p.Left, q.Right) && checkSymmetric(p.Right, q.Left)
	}
	return checkSymmetric(left, right)
}

func ModifyString(str string) string {
	re := regexp.MustCompile("[0-9]")
	str = strings.TrimSpace(str)
	str = re.ReplaceAllString(str, "")
	str_arr := strings.Split(str, "")
	slices.Reverse(str_arr)
	str = strings.Join(str_arr, "")

	return str
}

func RemainderSorting(strArr []string) []string {
	sort.Slice(strArr, func(i, j int) bool {
		a := strArr[i]
		b := strArr[j]
		modA := len(a) % 3
		modB := len(b) % 3
		if modA == modB {
			return a < b
		}
		return modA < modB
	})
	return strArr
}

func fizzBuzz(n int32) {
	s := ""
	for i := 1; i <= int(n); i++ {
		if i%3 == 0 {
			s = s + "Fizz"
		}
		if i%5 == 0 {
			s = s + "Buzz"
		}
		if s == "" {
			s = strconv.Itoa(i)
		}
		fmt.Print(s, "\n")
		s = ""
	}
}

func postorderTraversalLoop(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	var prev *TreeNode // To track the previously visited node
	cur := root

	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		// Peek the top node in the stack
		cur = stack[len(stack)-1]

		// Check if the right subtree exists and has not been visited
		if cur.Right == nil || cur.Right == prev {
			res = append(res, cur.Val) // Visit the current node
			stack = stack[:len(stack)-1]
			prev = cur // Mark this node as visited
			cur = nil  // Set current to nil to prevent re-traversal
		} else {
			cur = cur.Right // Move to the right subtree
		}
	}
	return res
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	postorder(root, &res)
	return res
}

func postorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	postorder(node.Left, res)
	postorder(node.Right, res)
	*res = append(*res, node.Val)
}

func preorderTraversalLoop(root *TreeNode) []int {
	res := []int{}
	stack := []TreeNode{}
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

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	preorder(root, &res)
	return res
}

func preorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	preorder(node.Left, res)
	preorder(node.Right, res)
}

func inorderTraversalLoop(root *TreeNode) []int {
	res := []int{}
	stack := []TreeNode{}
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

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	inorder(root, &res)
	return res
}

func inorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, res)
	*res = append(*res, node.Val)
	inorder(node.Right, res)
}

// Helper function to create a binary tree from a slice of integers
// Use 99999 as a placeholder for nil nodes
func createTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		// Assign left child
		if i < len(values) && values[i] != 99999 {
			node.Left = &TreeNode{Val: values[i]}
			queue = append(queue, node.Left)
		}
		i++

		// Assign right child
		if i < len(values) && values[i] != 99999 {
			node.Right = &TreeNode{Val: values[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// Helper function to print a tree in level order (breadth-first traversal)
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("Tree is empty")
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			fmt.Print("nil ")
			continue
		}

		fmt.Printf("%d ", node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	fmt.Println()
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n > 0 {
		j := 0
		for i := m; i < len(nums1); i++ {
			nums1[i] = nums2[j]
			j++
		}
		// slices.Sort(nums1)
		sort.Ints(nums1)
	}
}

// TreeNode struct definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
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

func nextUniqueNode(head *ListNode, val int) *ListNode {
	for head != nil {
		if head.Val != val {
			return head
		}
		head = head.Next
	}
	return nil
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	current := head
	for _, val := range vals[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func climbStairs(n int) int {
	f := 0
	f0 := 1
	f1 := 1

	if n == 1 {
		return f1
	}

	for i := 1; i < n; i++ {
		f = f0 + f1
		f0 = f1
		f1 = f
	}

	return f
}

func mySqrt(x int) int {
	if x <= 3 && x > 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	s := 2
	for i := s; true; i++ {
		if i*i > x {
			s = i - 1
			break
		}
	}
	return s
}

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	res := make([]string, len(a))

	var carry, sum byte
	for i := 0; i < len(a); i++ {
		bl := (len(a) - 1) - i
		sl := (len(b) - 1) - i
		bi := a[bl] - '0'
		si := byte(0)
		if sl > -1 {
			si = b[sl] - '0'
		}
		sum = carry + bi + si
		// fmt.Println("bi:", bi, "si:", si, "sum:", sum)
		switch sum {
		case 3:
			res[bl] = "1"
			carry = 1
		case 2:
			res[bl] = "0"
			carry = 1
		case 1:
			res[bl] = "1"
			carry = 0
		case 0:
			res[bl] = "0"
			carry = 0
		}
	}
	// fmt.Println(res, carry)
	if carry > 0 {
		return "1" + strings.Join(res, "")
	}
	return strings.Join(res, "")
}

func addBinarySlow(a string, b string) string {
	sum := binaryToDecimal(a) + binaryToDecimal(b)
	if sum > 1 {
		return fmt.Sprintf("%b", sum)
	}
	return strconv.Itoa(sum)
}

func binaryToDecimal(s string) int {
	var decimal float64
	bi := len(s) - 1
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "1" {
			decimal += math.Pow(2, float64(bi-i))
		}
	}
	fmt.Println(decimal)
	return int(decimal)
}

func decimalToBinary(n int) string {
	power := 0
	for i := 0; true; i++ {
		if math.Pow(2, float64(power)) > float64(n) {
			power = i - 1
			break
		}
		power++
	}
	var sb strings.Builder
	for i := power; i > -1; i-- {
		if n >= int(math.Pow(2, float64(i))) {
			n = n - int(math.Pow(2, float64(i)))
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
	}
	return sb.String()
}

func lengthOfLastWord(s string) int {
	arr := strings.Split(s, " ")
	fil := []string{}
	for _, v := range arr {
		if v != "" {
			fil = append(fil, v)
		}
	}
	return len(fil[len(fil)-1])
}

func plusOne(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	return append([]int{1}, digits...)
}
