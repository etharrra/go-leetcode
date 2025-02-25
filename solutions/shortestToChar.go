package solutions

import "fmt"

func ShortestToChar(s string, c byte) []int {
	fmt.Println(s, len(s))
	res := make([]int, len(s))

	// First pass: find the shortest distance to the previous occurrence of c
	prev := -1
	for i := range s {
		if s[i] == c {
			prev = i
		}
		if prev == -1 {
			res[i] = len(s) // Use a large number to indicate no occurrence found yet
		} else {
			res[i] = i - prev
		}
		fmt.Printf("First pass - index: %d, char: %c, res: %v\n", i, s[i], res)
	}

	// Second pass: find the shortest distance to the next occurrence of c
	next := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			next = i
		}
		if next != -1 {
			res[i] = smaller(res[i], next-i)
		}
		fmt.Printf("Second pass - index: %d, char: %c, res: %v\n", i, s[i], res)
	}

	return res
}

func smaller(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
