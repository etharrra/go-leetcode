package solutions

func MoveZeroes(nums []int) {
	// Initialize nonZeroIndex to keep track of the position to place the next non-zero element
	nonZeroIndex := 0

	for i := range nums {
		if nums[i] != 0 {
			// Swap the current element with the element at nonZeroIndex
			nonZeroValue := nums[nonZeroIndex]
			nums[nonZeroIndex] = nums[i]
			nums[i] = nonZeroValue
			// Move the nonZeroIndex to the next position
			nonZeroIndex++
		}
	}
}
