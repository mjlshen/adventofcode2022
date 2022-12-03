package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("input.txt"))
	fmt.Printf("Part 2: %d\n", part2("input.txt"))
}

func part1(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var rucksack string
		if _, err := fmt.Sscanf(scanner.Text(), "%s", &rucksack); err != nil {
			panic(err)
		}
		score += scoreRucksack(rucksack)
	}

	return score
}

func part2(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var r1, r2, r3 string
		if _, err := fmt.Sscanf(scanner.Text(), "%s", &r1); err != nil {
			panic(err)
		}
		scanner.Scan()
		if _, err := fmt.Sscanf(scanner.Text(), "%s", &r2); err != nil {
			panic(err)
		}
		scanner.Scan()
		if _, err := fmt.Sscanf(scanner.Text(), "%s", &r3); err != nil {
			panic(err)
		}
		score += scoreBadges(r1, r2, r3)
	}

	return score
}

// scoreBadges returns the score of the rune that appears in all three of r1, r2, and r3
func scoreBadges(r1, r2, r3 string) int {
	seen := make(map[rune]bool)
	for _, r := range r1 {
		seen[r] = true
	}

	seenAgain := make(map[rune]bool)
	for _, r := range r2 {
		if _, ok := seen[r]; ok {
			seenAgain[r] = true
		}
	}

	for _, r := range r3 {
		if _, ok := seenAgain[r]; ok {
			return scoreRune(r)
		}
	}
	return 0
}

// scoreRucksack scores the rune that appears in both the first and second half of the rucksack
func scoreRucksack(rucksack string) int {
	seen := make(map[rune]bool)
	first := rucksack[:len(rucksack)/2]
	second := rucksack[len(rucksack)/2:]
	for _, r := range first {
		seen[r] = true
	}

	for _, r := range second {
		if _, ok := seen[r]; ok {
			return scoreRune(r)
		}
	}

	return 0
}

// scoreRune assigns a value to each rune
// a-z = 1-26
// A-Z = 27-52
func scoreRune(r rune) int {
	// ASCII a - 97
	// ASCII A - 65
	if int(r) < int('a') {
		// A-Z
		return int(r) - int('A') + 1 + 26
	} else {
		// a-z
		return int(r) - int('a') + 1
	}
}
