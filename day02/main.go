package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", rpsTournament("test.txt", part1))
	fmt.Printf("Part 2: %d\n", rpsTournament("test.txt", part2))
}

func rpsTournament(path string, strategy func(string, string) (string, string)) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var first, second string
		if _, err := fmt.Sscanf(scanner.Text(), "%s %s", &first, &second); err != nil {
			panic(err)
		}

		score += rps(strategy(first, second))
	}

	return score
}

// rps plays Rock Paper Scissors and returns a score
// opponent: A - Rock, B - Paper, C - Scissors
// self: X - Rock, Y - Paper, Z - Scissors
func rps(opponent, self string) int {
	var score int

	switch self {
	case "X":
		score += 1
		switch opponent {
		case "C":
			score += 6
		case "A":
			score += 3
		}
	case "Y":
		score += 2
		switch opponent {
		case "A":
			score += 6
		case "B":
			score += 3
		}
	case "Z":
		score += 3
		switch opponent {
		case "B":
			score += 6
		case "C":
			score += 3
		}
	}

	return score
}

// part1 feeds-forward an input opponent move and self move
// opponent: A - Rock, B - Paper, C - Scissors
// self: X - Rock, Y - Paper, Z - Scissors
func part1(opponent, self string) (string, string) {
	return opponent, self
}

// part2 takes an input opponent move and a desired result
// opponent: A - Rock, B - Paper, C - Scissors
// result: X - Lose, Y - Draw, Z - Win
// self: X - Rock, Y - Paper, Z - Scissors
func part2(opponent, result string) (string, string) {
	possibilities := map[string]map[string]string{
		"A": {
			"X": "Z",
			"Y": "X",
			"Z": "Y",
		},
		"B": {
			"X": "X",
			"Y": "Y",
			"Z": "Z",
		},
		"C": {
			"X": "Y",
			"Y": "Z",
			"Z": "X",
		},
	}

	return opponent, possibilities[opponent][result]
}
