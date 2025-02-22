package solutions

import "fmt"

func MissingNumber(nums []int) int {
	sum, numsSum := 0, 0
	for i := 0; i <= len(nums); i++ {
		sum += i
		if i < len(nums) {
			numsSum += nums[i]
		}
	}
	fmt.Println(sum, numsSum)

	return sum - numsSum
}
