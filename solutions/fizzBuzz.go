package solutions

import (
	"fmt"
	"strconv"
)

func FizzBuzz(n int32) {
	s := ""
	for i := 1; i <= int(n); i++ {
		if i%3 == 0 {
			s = s + "Fizz"
		}
		if i%5 == 0 {
			s = s + "Buzz"
		}
		if s == "" {
			s = strconv.Itoa(i)
		}
		fmt.Print(s, "\n")
		s = ""
	}
}
