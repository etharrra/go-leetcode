package solutions

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s_map := make(map[byte]int, 0)
	t_map := make(map[byte]int, 0)

	for i := 0; i < len(s); i++ {
		s_map[s[i]] = s_map[s[i]] + 1
		t_map[t[i]] = t_map[t[i]] + 1

	}

	for k, _ := range s_map {
		if s_map[k] != t_map[k] {
			return false
		}
	}

	return true
}
