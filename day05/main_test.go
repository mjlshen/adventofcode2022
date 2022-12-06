package main

import (
	"fmt"
	"testing"
)

func stacksEqual(a []stack, b []stack) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

func TestMoveCrates9000(t *testing.T) {
	tests := []struct {
		count    int
		source   int
		dest     int
		stacks   []stack
		expected []stack
	}{
		{
			count:  1,
			source: 2,
			dest:   1,
			stacks: []stack{
				[]string{"Z", "N"},
				[]string{"M", "C", "D"},
				[]string{"P"},
			},
			expected: []stack{
				[]string{"Z", "N", "D"},
				[]string{"M", "C"},
				[]string{"P"},
			},
		},
	}

	for _, test := range tests {
		actual := moveCrates9000(test.count, test.source, test.dest, test.stacks)
		if !stacksEqual(test.expected, actual) {
			t.Errorf("expected %v, got %v", test.expected, actual)
		}
	}
}

func TestMoveCrates9001(t *testing.T) {
	tests := []struct {
		count    int
		source   int
		dest     int
		stacks   []stack
		expected []stack
	}{
		{
			count:  3,
			source: 1,
			dest:   3,
			stacks: []stack{
				[]string{"Z", "N", "D"},
				[]string{"M", "C"},
				[]string{"P"},
			},
			expected: []stack{
				[]string{},
				[]string{"M", "C"},
				[]string{"P", "Z", "N", "D"},
			},
		},
	}

	for _, test := range tests {
		actual := moveCrates9001(test.count, test.source, test.dest, test.stacks)
		if !stacksEqual(test.expected, actual) {
			t.Errorf("expected %v, got %v", test.expected, actual)
		}
	}
}

func TestReverseSlice(t *testing.T) {
	tests := []struct {
		s        []string
		expected []string
	}{
		{
			s:        []string{"A", "B", "C"},
			expected: []string{"C", "B", "A"},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.s), func(t *testing.T) {
			actual := reverseSlice(test.s)
			if len(test.expected) != len(actual) {
				t.Errorf("expected: %v, got %v", test.expected, actual)
			}
			for i := 0; i < len(test.expected); i++ {
				if test.expected[i] != actual[i] {
					t.Errorf("expected: %v, got %v", test.expected, actual)
				}
			}
		})
	}
}
