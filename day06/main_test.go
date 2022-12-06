package main

import "testing"

func TestTune(t *testing.T) {
	tests := []struct {
		input    string
		marker   int
		expected int
	}{
		{
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			marker:   4,
			expected: 7,
		},
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			marker:   4,
			expected: 5,
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			marker:   4,
			expected: 6,
		},
		{
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			marker:   4,
			expected: 10,
		},
		{
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			marker:   4,
			expected: 11,
		},
		{
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			marker:   14,
			expected: 19,
		},
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			marker:   14,
			expected: 23,
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			marker:   14,
			expected: 23,
		},
		{
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			marker:   14,
			expected: 29,
		},
		{
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			marker:   14,
			expected: 26,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := tune(test.input, test.marker)
			if test.expected != actual {
				t.Errorf("expected: %d, got: %d", test.expected, actual)
			}
		})
	}
}
