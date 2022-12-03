package main

import (
	"testing"
)

func TestRpsTournament(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		strategy func(string, string) (string, string)
		expected int
	}{
		{
			name:     "Part 1",
			path:     "test.txt",
			strategy: part1,
			expected: 15,
		},
		{
			name:     "Part 2",
			path:     "test.txt",
			strategy: part2,
			expected: 12,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := rpsTournament(test.path, test.strategy)
			if test.expected != actual {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
