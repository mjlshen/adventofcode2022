package main

import (
	"testing"
)

func TestRegolithReservoir(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		floor    bool
		expected int
	}{
		{
			name:     "Part 1",
			path:     "test.txt",
			floor:    false,
			expected: 24,
		},
		{
			name:     "Part 2",
			path:     "test.txt",
			floor:    true,
			expected: 93,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := regolithReservoir(test.path, test.floor)
			if test.expected != actual {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
