package main

import "testing"

func TestDistressSignal(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		part     func([][]any) int
		expected int
	}{
		{
			name:     "Part 1",
			path:     "test.txt",
			part:     part1,
			expected: 13,
		},
		{
			name:     "Part 2",
			path:     "test.txt",
			part:     part2,
			expected: 140,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := distressSignal(test.path, test.part)
			if test.expected != actual {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		name        string
		left, right any
		expected    int
	}{
		{
			name:     "Both ints",
			left:     []any{float64(5)},
			right:    []any{float64(7)},
			expected: -2,
		},
		{
			name:     "int and list",
			left:     []any{float64(1)},
			right:    float64(3),
			expected: -2,
		},
		{
			name:     "list of empty lists",
			left:     []any{[]any{[]any{}}},
			right:    []any{[]any{}},
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := compare(test.left, test.right)
			if test.expected != actual {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}
