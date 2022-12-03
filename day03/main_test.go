package main

import "testing"

func TestRucksackReorganization(t *testing.T) {
	tests := []struct {
		name     string
		part     func(string) int
		path     string
		expected int
	}{
		{
			name:     "Part 1",
			part:     part1,
			path:     "test.txt",
			expected: 157,
		},
		{
			name:     "Part 2",
			part:     part2,
			path:     "test.txt",
			expected: 70,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.part(test.path)
			if test.expected != actual {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}

func TestScoreRucksack(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		{
			name:     "vJrwpWtwJgWrhcsFMMfFFhFp",
			expected: 16,
		},
		{
			name:     "CrZsJsPPZsGzwwsLwLmpwMDw",
			expected: 19,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := scoreRucksack(test.name)
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}

func TestScoreRune(t *testing.T) {
	tests := []struct {
		name     string
		r        rune
		expected int
	}{
		{
			name:     "a",
			r:        'a',
			expected: 1,
		},
		{
			name:     "L",
			r:        'L',
			expected: 38,
		},
		{
			name:     "Z",
			r:        'Z',
			expected: 52,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := scoreRune(test.r)
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
