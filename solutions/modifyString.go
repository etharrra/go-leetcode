package solutions

import (
	"regexp"
	"slices"
	"strings"
)

func ModifyString(str string) string {
	re := regexp.MustCompile("[0-9]")
	str = strings.TrimSpace(str)
	str = re.ReplaceAllString(str, "")
	str_arr := strings.Split(str, "")
	slices.Reverse(str_arr)
	str = strings.Join(str_arr, "")

	return str
}
