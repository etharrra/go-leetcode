package main

import (
	"fmt"

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
