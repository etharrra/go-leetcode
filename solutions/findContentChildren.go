package solutions

import (
	"fmt"
	"slices"
)

func FindContentChildren(g []int, s []int) int {
	fmt.Println(g, s, "---")
	count := 0
	slices.Sort(s)
	slices.Sort(g)

	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			count++
			i++
		}
		j++
	}
	return count
}
