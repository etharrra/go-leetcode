package solutions

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "single word",
			input:    "Hello",
			expected: 5,
		},
		{
			name:     "two words",
			input:    "Hello World",
			expected: 5,
		},
		{
			name:     "trailing spaces",
			input:    "Hello World  ",
			expected: 5,
		},
		{
			name:     "leading spaces",
			input:    "   Hello World",
			expected: 5,
		},
		{
			name:     "multiple spaces between words",
			input:    "Hello      World",
			expected: 5,
		},
		{
			name:     "single letter word",
			input:    "a",
			expected: 1,
		},
		{
			name:     "multiple words",
			input:    "Today is a nice day",
			expected: 3,
		},
		{
			name:     "sentence with punctuation",
			input:    "Hello World!",
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLastWord(tt.input); got != tt.expected {
				t.Errorf("LengthOfLastWord(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
