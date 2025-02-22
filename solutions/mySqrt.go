package solutions

func MySqrt(x int) int {
	if x <= 3 && x > 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	s := 2
	for i := s; true; i++ {
		if i*i > x {
			s = i - 1
			break
		}
	}
	return s
}
