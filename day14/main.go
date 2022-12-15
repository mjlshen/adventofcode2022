package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", regolithReservoir("input.txt", false))
	fmt.Printf("Part 2: %d\n", regolithReservoir("input.txt", true))
}

type coord struct {
	x, y int
}

func regolithReservoir(path string, floor bool) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	maxY := 0
	grid := make(map[coord]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		endpoints := strings.Split(scanner.Text(), " -> ")

		for i := 0; i < len(endpoints)-1; i++ {
			var head, tail coord
			if _, err := fmt.Sscanf(endpoints[i], "%d,%d", &head.x, &head.y); err != nil {
				panic(err)
			}
			if _, err := fmt.Sscanf(endpoints[i+1], "%d,%d", &tail.x, &tail.y); err != nil {
				panic(err)
			}
			if max(head.y, tail.y) > maxY {
				maxY = max(head.y, tail.y)
			}

			if head.x == tail.x {
				for i := min(head.y, tail.y); i <= max(head.y, tail.y); i++ {
					grid[coord{head.x, i}] = true
				}
			}

			if head.y == tail.y {
				for i := min(head.x, tail.x); i <= max(head.x, tail.x); i++ {
					grid[coord{i, head.y}] = true
				}
			}
		}
	}

	return simulate(grid, maxY, floor)
}

func simulate(grid map[coord]bool, maxY int, floor bool) int {
	ans := 0
	if floor {
		maxY += 2
	}

	for {
		sand := coord{500, 0}
		for sand.y <= maxY && !grid[sand] {
			if floor && sand.y == maxY-1 {
				grid[sand] = true
				ans++
				break
			}

			switch {
			case !grid[coord{sand.x, sand.y + 1}]:
				// Sand falls down if it can
				sand.y++
			case !grid[coord{sand.x - 1, sand.y + 1}]:
				// If it can't fall down it falls diagonally left
				sand.y++
				sand.x--
			case !grid[coord{sand.x + 1, sand.y + 1}]:
				// If it can't fall diagonally left it falls diagonally right
				sand.y++
				sand.x++
			default:
				// It's found a home
				grid[sand] = true
				ans++
			}
		}

		// The simulation should end if the final state of sand is also it's starting point
		// or if it's falling off the screen (past maxY)
		if sand.x == 500 && sand.y == 0 || sand.y > maxY {
			break
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
