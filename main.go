package main

import (
	"fmt"

	"github.com/etharrra/go-leetcode/solutions"
)

func main() {
	s := "loveleetcode"
	c := 'e'
	s = "aaba"
	c = 'b'
	fmt.Println(solutions.ShortestToChar(s, byte(c)))
}
