package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", ropeBridge("input.txt", 1))
	fmt.Printf("Part 2: %d\n", ropeBridge("input.txt", 9))
}

type coord struct {
	x, y int
}

func ropeBridge(path string, numTails int) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	head := &coord{0, 0}
	tails := make([]*coord, numTails)
	for i := 0; i < numTails; i++ {
		tails[i] = &coord{0, 0}
	}

	visited := map[coord]bool{
		coord{0, 0}: true,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var direction string
		var amount int
		if _, err := fmt.Sscanf(scanner.Text(), "%s %d", &direction, &amount); err != nil {
			panic(err)
		}

		for i := 0; i < amount; i++ {
			switch direction {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y++
			case "D":
				head.y--
			}

			prev := head
			for _, tail := range tails {
				tail.follow(*prev)
				prev = tail
			}
			visited[*tails[len(tails)-1]] = true
		}
	}

	return len(visited)
}

func (tail *coord) follow(head coord) {
	if abs(head.x-tail.x) <= 1 && abs(head.y-tail.y) <= 1 {
		return
	}

	// Two spaces horizontally away
	if abs(head.x-tail.x) == 2 {
		// If there's no vertical spacing
		if head.y == tail.y {
			if head.x > tail.x {
				tail.x++
			} else {
				tail.x--
			}
		} else {
			if head.x > tail.x {
				if head.y > tail.y {
					tail.x++
					tail.y++
				} else {
					tail.x++
					tail.y--
				}
			} else {
				if head.y > tail.y {
					tail.x--
					tail.y++
				} else {
					tail.x--
					tail.y--
				}
			}
		}
	}

	// Two spaces vertically away
	if abs(head.y-tail.y) == 2 {
		// If there's no horizontal spacing
		if head.x == tail.x {
			if head.y > tail.y {
				tail.y++
			} else {
				tail.y--
			}
		} else {
			if head.y > tail.y {
				if head.x > tail.x {
					tail.x++
					tail.y++
				} else {
					tail.x--
					tail.y++
				}
			} else {
				if head.x > tail.x {
					tail.x++
					tail.y--
				} else {
					tail.x--
					tail.y--
				}
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
