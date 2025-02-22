package solutions

import "fmt"

func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	suffix := make([]int, n)
	res := make([]int, n)

	prefix[0], suffix[n-1] = 1, 1
	fmt.Println(nums)
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
		fmt.Println("pre", prefix)
	}

	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
		fmt.Println("suf", suffix)
	}

	for i := 0; i < n; i++ {
		res[i] = prefix[i] * suffix[i]
	}

	return res
}
