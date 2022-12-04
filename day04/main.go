package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", campCleanup("input.txt", contains))
	fmt.Printf("Part 2: %d\n", campCleanup("input.txt", overlaps))
}

// campCleanup returns the sum of the number of calories the topK elves are carrying
func campCleanup(path string, criteria func(int, int, int, int) bool) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reassignments := 0

	// Each elf's items are separated by an empty line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var min1, max1, min2, max2 int
		if _, err := fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &min1, &max1, &min2, &max2); err != nil {
			panic(err)
		}
		if criteria(min1, max1, min2, max2) {
			reassignments++
		}
	}

	return reassignments
}

// contains returns true if either of the ranges [start1, finish1] and [start2, finish2]
// are entirely contained in the other
func contains(start1, finish1, start2, finish2 int) bool {
	if start1 >= start2 && finish1 <= finish2 {
		return true
	}

	if start2 >= start1 && finish2 <= finish1 {
		return true
	}

	return false
}

// overlaps returns true if either of the ranges [start1, finish1] and [start2, finish2] overlap at all
func overlaps(start1, finish1, start2, finish2 int) bool {
	if (finish1 >= start2 && finish1 <= finish2) || (finish2 >= start1 && finish2 <= finish1) {
		return true
	}

	return false
}
