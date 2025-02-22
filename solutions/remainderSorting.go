package solutions

import "sort"

func RemainderSorting(strArr []string) []string {
	sort.Slice(strArr, func(i, j int) bool {
		a := strArr[i]
		b := strArr[j]
		modA := len(a) % 3
		modB := len(b) % 3
		if modA == modB {
			return a < b
		}
		return modA < modB
	})
	return strArr
}
