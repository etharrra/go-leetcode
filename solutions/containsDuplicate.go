package solutions

func ContainsDuplicate(nums []int) bool {
	indexMap := make(map[int]bool, 0)
	// fmt.Println(indexMap[0])
	for _, v := range nums {
		if indexMap[v] {
			return true
		}
		indexMap[v] = true
	}
	return false
}
