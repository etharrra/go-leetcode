package solutions

func MajorityElement(nums []int) int {
	hash := make(map[int]int)
	majorityCount := len(nums) / 2

	for _, n := range nums {
		hash[n]++
		if hash[n] > majorityCount {
			return n
		}
	}

	return 0
}
