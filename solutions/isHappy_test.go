package solutions

import "testing"

func TestIsHappy(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{19, true},
		{2, false},
		{7, true},
		{4, false},
	}

	for _, test := range tests {
		result := IsHappy(test.input)
		if result != test.expected {
			t.Errorf("IsHappy(%d) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
