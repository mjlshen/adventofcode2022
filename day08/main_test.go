package main

import (
	"testing"
)

func TestRpsTournament(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		scoring  func([][]int) int
		expected int
	}{
		{
			name:     "Part 1",
			path:     "test.txt",
			scoring:  countVisibleTrees,
			expected: 21,
		},
		{
			name:     "Part 2",
			path:     "test.txt",
			scoring:  mostScenicTree,
			expected: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := visibleTrees(test.path, test.scoring)
			if test.expected != actual {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
