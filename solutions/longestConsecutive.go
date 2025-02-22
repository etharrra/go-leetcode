package solutions

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	hash := make(map[int]bool)
	for _, num := range nums {
		hash[num] = true
	}

	longestStreak := 0

	for num := range hash {
		if !hash[num-1] {
			currentNum := num
			currentStreak := 1

			for hash[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}
