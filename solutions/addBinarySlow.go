package solutions

import (
	"fmt"
	"math"
	"strconv"
)

func AddBinarySlow(a string, b string) string {
	sum := binaryToDecimal(a) + binaryToDecimal(b)
	if sum > 1 {
		return fmt.Sprintf("%b", sum)
	}
	return strconv.Itoa(sum)
}

func binaryToDecimal(s string) int {
	var decimal float64
	bi := len(s) - 1
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "1" {
			decimal += math.Pow(2, float64(bi-i))
		}
	}
	fmt.Println(decimal)
	return int(decimal)
}
