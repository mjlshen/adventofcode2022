package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	fmt.Printf("Part 1: %d\n", maxCalories("input.txt", 1))
	fmt.Printf("Part 2: %d\n", maxCalories("input.txt", 3))
}

// maxCalories returns the sum of the number of calories the topK elves are carrying
func maxCalories(path string, topK int) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var elves []int
	totalCalories := 0

	// Each elf's items are separated by an empty line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			elves = append(elves, totalCalories)
			totalCalories = 0
		} else {
			var calories int
			if _, err := fmt.Sscanf(scanner.Text(), "%d", &calories); err != nil {
				panic(err)
			}
			totalCalories += calories
		}
	}

	// The last elf doesn't have an empty line after it
	elves = append(elves, totalCalories)
	sort.Ints(elves)

	return sum(elves[len(elves)-topK:]...)
}

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}
