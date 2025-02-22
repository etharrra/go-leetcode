package solutions

func GetIndex(nums []int, target int) []int {
	res := []int{}

	for i_index, i := range nums {
		if i <= target {
			for j_index, j := range nums {
				if i == j {
					continue
				}
				if len(res) > 2 {
					break
				}
				if i+j == target && len(res) < 2 {
					res = append(res, i_index, j_index)
				}
			}
		}
	}

	return res
}
