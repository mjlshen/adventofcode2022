package main

import (
	"testing"
)

func TestBoilingBoulders(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		part     func(map[coord]struct{}) int
		expected int
	}{
		{
			name:     "Part 1",
			path:     "test.txt",
			part:     part1,
			expected: 64,
		},
		{
			name:     "Part 2",
			path:     "test.txt",
			part:     part2,
			expected: 58,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := boilingBoulders(test.path, test.part)
			if test.expected != actual {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
