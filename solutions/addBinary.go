package solutions

import "strings"

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	res := make([]string, len(a))

	var carry, sum byte
	for i := 0; i < len(a); i++ {
		bl := (len(a) - 1) - i
		sl := (len(b) - 1) - i
		bi := a[bl] - '0'
		si := byte(0)
		if sl > -1 {
			si = b[sl] - '0'
		}
		sum = carry + bi + si
		// fmt.Println("bi:", bi, "si:", si, "sum:", sum)
		switch sum {
		case 3:
			res[bl] = "1"
			carry = 1
		case 2:
			res[bl] = "0"
			carry = 1
		case 1:
			res[bl] = "1"
			carry = 0
		case 0:
			res[bl] = "0"
			carry = 0
		}
	}
	// fmt.Println(res, carry)
	if carry > 0 {
		return "1" + strings.Join(res, "")
	}
	return strings.Join(res, "")
}
