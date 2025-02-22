package solutions

import "strings"

func LengthOfLastWord(s string) int {
	arr := strings.Split(s, " ")
	fil := []string{}
	for _, v := range arr {
		if v != "" {
			fil = append(fil, v)
		}
	}
	return len(fil[len(fil)-1])
}
