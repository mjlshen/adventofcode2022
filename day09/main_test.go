package main

import (
	"testing"
)

func TestRopeBridge(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		numTails int
		expected int
	}{
		{
			name:     "Part 1",
			path:     "test.txt",
			numTails: 1,
			expected: 13,
		},
		{
			name:     "Part 2",
			path:     "test.txt",
			numTails: 9,
			expected: 1,
		},
		{
			name:     "Part 2",
			path:     "test2.txt",
			numTails: 9,
			expected: 36,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := ropeBridge(test.path, test.numTails)
			if test.expected != actual {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
