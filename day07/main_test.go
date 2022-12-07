package main

import (
	"testing"
)

func TestClearingSpace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		part     func(map[string]int) int
		expected int
	}{
		{
			name:     "Part 1",
			input:    "test.txt",
			part:     part1,
			expected: 95437,
		},
		{
			name:     "Part 2",
			input:    "test.txt",
			part:     part2,
			expected: 24933642,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := clearingSpace(test.input, test.part)
			if test.expected != actual {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
