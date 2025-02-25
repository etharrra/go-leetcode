package solutions

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{0, 1, 0, 3, 12}, expected: []int{1, 3, 12, 0, 0}},
		{input: []int{0, 0, 1}, expected: []int{1, 0, 0}},
		{input: []int{1, 0, 1}, expected: []int{1, 1, 0}},
		{input: []int{0, 0, 0}, expected: []int{0, 0, 0}},
		{input: []int{1, 2, 3}, expected: []int{1, 2, 3}},
	}

	for _, test := range tests {
		MoveZeroes(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("MoveZeroes(%v) = %v; expected %v", test.input, test.input, test.expected)
		}
	}
}
