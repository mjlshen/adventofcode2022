package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: %s\n", supplyStacks("input.txt", moveCrates9000))
	fmt.Printf("Part 2: %s\n", supplyStacks("input.txt", moveCrates9001))
}

type stack []string

// supplyStacks returns the sum of the number of calories the topK elves are carrying
func supplyStacks(path string, crane func(int, int, int, []stack) []stack) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// TODO: Parse the input
	// test.txt
	//stacks := []stack{
	//	[]string{"Z", "N"},
	//	[]string{"M", "C", "D"},
	//	[]string{"P"},
	//}

	// input.txt
	stacks := []stack{
		[]string{"F", "T", "C", "L", "R", "P", "G", "Q"},
		[]string{"N", "Q", "H", "W", "R", "F", "S", "J"},
		[]string{"F", "B", "H", "W", "P", "M", "Q"},
		[]string{"V", "S", "T", "D", "F"},
		[]string{"Q", "L", "D", "W", "V", "F", "Z"},
		[]string{"Z", "C", "L", "S"},
		[]string{"Z", "B", "M", "V", "D", "F"},
		[]string{"T", "J", "B"},
		[]string{"Q", "N", "B", "G", "L", "S", "P", "H"},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var count, source, dest int
		if _, err := fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &count, &source, &dest); err != nil {
			// As a result of manually feeding in the input, ignore errors
			continue
		}

		stacks = crane(count, source, dest, stacks)
	}

	return ans(stacks)
}

// moveCrates9000 pops and pushes count crates one at a time from source to dest.
// source/dest are 1-indexed by identifier
func moveCrates9000(count, source, dest int, stacks []stack) []stack {
	// Take the last count elements from source, reverse that slice, and append it to the destination slice
	stacks[dest-1] = append(stacks[dest-1], reverseSlice(stacks[source-1][len(stacks[source-1])-count:])...)
	// Remove the last count elements from source
	stacks[source-1] = stacks[source-1][:len(stacks[source-1])-count]

	return stacks
}

// moveCrates9000 pops and pushes all count crates at once from source to dest.
// source/dest are 1-indexed by identifier
func moveCrates9001(count, source, dest int, stacks []stack) []stack {
	// Take the last count elements from source and append it to the destination slice
	stacks[dest-1] = append(stacks[dest-1], stacks[source-1][len(stacks[source-1])-count:]...)
	// Remove the last count elements from source
	stacks[source-1] = stacks[source-1][:len(stacks[source-1])-count]

	return stacks
}

func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

// ans formats the answer, which is the top element of each stack
func ans(stacks []stack) string {
	var a string
	for _, s := range stacks {
		if len(s) == 0 {
			continue
		}
		a += s[len(s)-1]
	}

	return a
}
