package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func main() {
	fmt.Printf("Part 1: %d\n", distressSignal("input2.txt", part1))
	fmt.Printf("Part 2: %d\n", distressSignal("input.txt", part2))
}

func distressSignal(path string, part func([][]any) int) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var packets [][]any
	var pair []any

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			packets = append(packets, pair)
			pair = []any{}
			continue
		}

		var p any
		if err := json.Unmarshal([]byte(line), &p); err != nil {
			panic(err)
		}
		pair = append(pair, p)
	}
	packets = append(packets, pair)

	return part(packets)
}

func part1(packets [][]any) int {
	ans := 0
	for i := range packets {
		if compare(packets[i][0], packets[i][1]) <= 0 {
			ans += i + 1
		}
	}

	return ans
}

func part2(packets [][]any) int {
	ans := 1

	var ordered []any
	for _, pair := range packets {
		ordered = append(ordered, pair...)
	}

	var dividers = []any{[]any{[]any{float64(2)}}, []any{[]any{float64(6)}}}
	ordered = append(ordered, dividers...)

	sort.Slice(ordered, func(i, j int) bool {
		return compare(ordered[i], ordered[j]) <= 0
	})

	for i := range ordered {
		packet, err := json.Marshal(ordered[i])
		if err != nil {
			panic(err)
		}

		if string(packet) == "[[2]]" || string(packet) == "[[6]]" {
			ans *= i + 1
		}
	}

	return ans
}

// compare is good if this returns a negative number. Positive numbers are treated as failures and 0 is no decision:
// The left packet integers are <= The right packet integers
// The left packet lists' integers are <= The right packet lists' integers in order
func compare(l, r any) int {
	var left, right []any

	// If they're both integers, just compare the values
	_, leftOk := l.(float64)
	_, rightOk := r.(float64)
	if leftOk && rightOk {
		return int(l.(float64) - r.(float64))
	}

	// Convert any plain integers to lists
	if leftOk {
		left = []any{l}
	} else {
		left = l.([]any)
	}
	if rightOk {
		right = []any{r}
	} else {
		right = r.([]any)
	}

	// If there are no elements left in the list, then ensure left is shorter than right
	if len(left) == 0 || len(right) == 0 {
		return len(left) - len(right)
	}

	if comparison := compare(left[0], right[0]); comparison != 0 {
		return comparison
	}

	left = left[1:]
	right = right[1:]

	return compare(left, right)
}
