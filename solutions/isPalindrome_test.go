package solutions

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "single character",
			input:    "a",
			expected: true,
		},
		{
			name:     "simple palindrome",
			input:    "aba",
			expected: true,
		},
		{
			name:     "palindrome with spaces",
			input:    "A man a plan a canal Panama",
			expected: true,
		},
		{
			name:     "palindrome with punctuation",
			input:    "Race a car!",
			expected: false,
		},
		{
			name:     "palindrome with numbers",
			input:    "12321",
			expected: true,
		},
		{
			name:     "mixed case palindrome",
			input:    "RaCeCaR",
			expected: true,
		},
		{
			name:     "non palindrome",
			input:    "hello",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.input); got != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
