package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", monkeyInTheMiddle("input.txt", 20, part1))
	fmt.Printf("Part 2: %d\n", monkeyInTheMiddle("input.txt", 10000, part2))
}

const monkeFmt = `Monkey %d:
  Starting items: %s
  Operation: new = old %s %d
  Test: divisible by %d
    If true: throw to monkey %d
    If false: throw to monkey %d`

type monke struct {
	items     []int
	operation func(int) int
	test      func(int) int
}

var (
	part1 = func(worry int) int { return worry / 3 }
	part2 = func(worry int) int { return worry }
)

func monkeyInTheMiddle(path string, rounds int, relief func(worry int) int) int {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	parsed := strings.Split(string(input), "\n\n")
	monkes := make([]monke, len(parsed))
	lcm := 1
	for _, p := range parsed {
		var is, operation string
		var i, operand, test, ifTrue, ifFalse int
		// Replace ", " with "," so that the starting items can all be parsed as one string
		// Replace "* old" with "^ 2" so that the operand can always be parsed as an integer
		if _, err := fmt.Sscanf(strings.NewReplacer(", ", ",", "* old", "^ 2").Replace(p), monkeFmt,
			&i, &is, &operation, &operand, &test, &ifTrue, &ifFalse); err != nil {
			panic(err)
		}

		monkes[i].items = parseItems(is)
		monkes[i].operation = monkeyOperation(operation, operand)
		monkes[i].test = func(worry int) int {
			if worry%test == 0 {
				return ifTrue
			}
			return ifFalse
		}
		// All the divisible tests are prime, so just multiply them to get the LCM
		lcm *= test
	}

	return inspection(monkes, lcm, rounds, relief)
}

func inspection(monkes []monke, lcm, rounds int, relief func(worry int) int) int {
	inspected := make([]int, len(monkes))
	for i := 0; i < rounds; i++ {
		for i, m := range monkes {
			for _, worry := range m.items {
				worry = m.operation(worry) % lcm
				worry = relief(worry)
				monkes[m.test(worry)].items = append(monkes[m.test(worry)].items, worry)
				inspected[i]++
			}
			monkes[i].items = make([]int, 0)
		}
	}

	sort.Sort(sort.IntSlice(inspected))
	return inspected[len(inspected)-1] * inspected[len(inspected)-2]
}

func parseItems(is string) []int {
	itemStrings := strings.Split(is, ",")
	items := make([]int, len(itemStrings))

	for i, s := range itemStrings {
		items[i], _ = strconv.Atoi(s)
	}

	return items
}

func monkeyOperation(op string, operand int) func(int) int {
	switch op {
	case "+":
		return func(out int) int { return out + operand }
	case "*":
		return func(out int) int { return out * operand }
	case "^":
		return func(out int) int { return out * out }
	}

	return func(int) int { return 0 }
}
