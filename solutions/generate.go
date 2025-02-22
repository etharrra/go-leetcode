package solutions

func Generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	res := [][]int{{1}, {1, 1}}
	for i := 3; i <= numRows; i++ {
		arr := make([]int, 0, i) // Initialize with capacity i
		for j := 0; j < i; j++ {
			if j == 0 || j+1 == i {
				arr = append(arr, 1)
			} else {
				prev := res[i-2]
				arr = append(arr, prev[j-1]+prev[j])
			}
		}
		res = append(res, arr)
	}
	return res
}
