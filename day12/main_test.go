package main

import (
	"testing"
)

func TestHillClimbing(t *testing.T) {
	tests := []struct {
		name              string
		path              string
		allPossibleStarts bool
		expected          int32
	}{
		{
			name:              "Part 1",
			path:              "test.txt",
			allPossibleStarts: false,
			expected:          31,
		},
		{
			name:              "Part 2",
			path:              "test.txt",
			allPossibleStarts: true,
			expected:          29,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := hillClimbing(test.path, test.allPossibleStarts)
			if test.expected != actual {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
