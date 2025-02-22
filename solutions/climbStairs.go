package solutions

func ClimbStairs(n int) int {
	f := 0
	f0 := 1
	f1 := 1

	if n == 1 {
		return f1
	}

	for i := 1; i < n; i++ {
		f = f0 + f1
		f0 = f1
		f1 = f
	}

	return f
}
