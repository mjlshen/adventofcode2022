package main

import (
	"testing"
)

func TestUnstableDiffusion(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		part     int
		expected int
	}{
		{
			name:     "Part 1",
			path:     "test.txt",
			part:     1,
			expected: 110,
		},
		{
			name:     "Part 1",
			path:     "test.txt",
			part:     2,
			expected: 20,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := unstableDiffusion(test.path, test.part)
			if test.expected != actual {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}

func TestGrove_round(t *testing.T) {
	tests := []struct {
		name     string
		g        *grove
		expected map[coord]struct{}
	}{
		{
			name: "N",
			g: &grove{
				elves: map[coord]struct{}{
					{2, 1}: {},
					{2, 2}: {},
					{2, 4}: {},
					{3, 1}: {},
					{3, 4}: {},
				},
				d: 0,
			},
			expected: map[coord]struct{}{
				{2, 0}: {},
				{2, 2}: {},
				{2, 4}: {},
				{3, 0}: {},
				{3, 3}: {},
			},
		},
		{
			name: "S",
			g: &grove{
				elves: map[coord]struct{}{
					{2, 0}: {},
					{2, 2}: {},
					{2, 4}: {},
					{3, 0}: {},
					{3, 3}: {},
				},
				d: 1,
			},
			expected: map[coord]struct{}{
				{1, 2}: {},
				{2, 1}: {},
				{2, 5}: {},
				{3, 1}: {},
				{4, 3}: {},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.g.round()
			actual := test.g.elves
			if len(actual) != len(test.expected) {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
			for elf := range test.expected {
				if _, ok := actual[elf]; !ok {
					t.Errorf("expected %v, got %v", test.expected, actual)
				}
			}
		})
	}
}
