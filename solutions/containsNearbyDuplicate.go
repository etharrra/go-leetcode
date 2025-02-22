package solutions

import "fmt"

func ContainsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int]int)

	for i, num := range nums {
		if lastIndex, found := hash[num]; found && i-lastIndex <= k {
			return true
		}
		hash[num] = i
	}
	fmt.Println(hash)

	return false
}
