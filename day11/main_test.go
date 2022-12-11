package main

import (
	"testing"
)

func TestMonkeyInTheMiddle(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		rounds   int
		relief   func(worry int) int
		expected int
	}{
		{
			name:     "Part 1",
			path:     "test.txt",
			rounds:   20,
			relief:   part1,
			expected: 10605,
		},
		{
			name:     "Part 1",
			path:     "test.txt",
			rounds:   10000,
			relief:   part2,
			expected: 2713310158,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := monkeyInTheMiddle(test.path, test.rounds, test.relief)
			if test.expected != actual {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
