package solutions

import "slices"

func Intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	res := make([]int, 0)
	for _, num := range nums2 {
		if slices.Contains(nums1, num) {
			i := slices.Index(nums1, num)
			nums1 = slices.Delete(nums1, i, i+1)
			res = append(res, num)
		}
	}
	return res
}
