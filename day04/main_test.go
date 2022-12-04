package main

import "testing"

func TestCampCleanup(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		criteria func(int, int, int, int) bool
		expected int
	}{
		{
			name:     "Part 1",
			path:     "test.txt",
			criteria: contains,
			expected: 2,
		},
		{
			name:     "Part 2",
			path:     "test.txt",
			criteria: overlaps,
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := campCleanup(test.path, test.criteria)
			if test.expected != actual {
				t.Errorf("expected: %d, got: %d", test.expected, actual)
			}
		})
	}
}
