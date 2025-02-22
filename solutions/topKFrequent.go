package solutions

import "fmt"

func TopKFrequent(nums []int, k int) []int {
	hash_map := make(map[int]int, 0)
	res := []int{}

	for _, num := range nums {
		hash_map[num]++
	}

	frequency := make([][]int, len(nums))
	for k, v := range hash_map {
		frequency[v-1] = append(frequency[v-1], k)
	}

	for i := len(frequency) - 1; i >= 0; i-- {
		if len(frequency[i]) > 0 {
			res = append(res, frequency[i]...)
		}
	}

	fmt.Println(frequency)
	fmt.Println(res)

	return res[:k]
}
