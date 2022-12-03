package main

import "testing"

func TestMaxCalories(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		topK     int
		expected int
	}{
		{
			name:     "Part 1",
			path:     "test.txt",
			topK:     1,
			expected: 24000,
		},
		{
			name:     "Part 2",
			path:     "test.txt",
			topK:     3,
			expected: 45000,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := maxCalories(test.path, test.topK)
			if test.expected != actual {
				t.Errorf("expected: %d, got: %d", test.expected, actual)
			}
		})
	}

}
