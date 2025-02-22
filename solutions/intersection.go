package solutions

func Intersection(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]bool, 0)
	set2 := make(map[int]bool, 0)

	for _, v := range nums1 {
		set1[v] = true
	}

	for _, v := range nums2 {
		if _, found := set1[v]; found {
			set2[v] = true
		}
	}
	res := make([]int, 0, len(set2))
	for k := range set2 {
		res = append(res, k)
	}
	// fmt.Println(set1, set2)
	return res
}
