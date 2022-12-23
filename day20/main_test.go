package main

import (
	"testing"
)

func TestGrovePositioningSystem(t *testing.T) {
	tests := []struct {
		name          string
		path          string
		decryptionKey int
		times         int
		expected      int
	}{
		{
			name:          "Part 1",
			path:          "test.txt",
			decryptionKey: 1,
			times:         1,
			expected:      3,
		},
		{
			name:          "Part 2",
			path:          "test.txt",
			decryptionKey: decryptionKey,
			times:         10,
			expected:      1623178306,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := grovePositioningSystem(test.path, test.decryptionKey, test.times)
			if test.expected != actual {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
