package solutions

func SingleNumber(nums []int) int {
	hash := make(map[int]int)

	for _, n := range nums {
		hash[n]++
	}

	for k, v := range hash {
		if v == 1 {
			return k
		}
	}
	return 0
}
