package solutions

import (
	"math"
	"strings"
)

func DecimalToBinary(n int) string {
	power := 0
	for i := 0; true; i++ {
		if math.Pow(2, float64(power)) > float64(n) {
			power = i - 1
			break
		}
		power++
	}
	var sb strings.Builder
	for i := power; i > -1; i-- {
		if n >= int(math.Pow(2, float64(i))) {
			n = n - int(math.Pow(2, float64(i)))
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
	}
	return sb.String()
}
