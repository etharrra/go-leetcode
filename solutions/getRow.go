package solutions

func GetRow(rowIndex int) []int {
	arr := make([]int, rowIndex+1) // Initialize with capacity i
	tmp := arr

	for i := 0; i < rowIndex+1; i++ {
		if i == 0 {
			arr = []int{1}
		} else if i == 1 {
			arr = []int{1, 1}
		} else {
			arr = []int{}
			for j := 0; j <= len(tmp); j++ {
				if j == 0 || j == len(tmp) {
					arr = append(arr, 1)
				} else {
					arr = append(arr, tmp[j-1]+tmp[j])
				}
			}
		}

		// fmt.Println("tmp", tmp)
		// fmt.Println(arr)
		tmp = arr
	}
	// fmt.Println("=========")
	return arr
}
