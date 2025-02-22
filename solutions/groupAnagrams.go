package solutions

func GroupAnagrams(strs []string) [][]string {
	res := [][]string{}
	hash := make(map[string][]string, len(strs))

	for _, str := range strs {
		// Use a frequency array to count character occurrences
		freq := make([]int, 26)
		for _, ch := range str {
			freq[ch-'a']++
		}

		// Build a unique key from the frequency array
		key := make([]byte, 0, 52)
		for i, count := range freq {
			if count > 0 {
				key = append(key, byte(i+'a'), byte(count+'0'))
			}
		}

		// Add the string to the appropriate group in the hash map
		hash[string(key)] = append(hash[string(key)], str)
	}

	// Collect results from the hash map
	for _, group := range hash {
		res = append(res, group)
	}

	return res
}
