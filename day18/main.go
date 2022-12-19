package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", boilingBoulders("input.txt", part1))
	fmt.Printf("Part 2: %d\n", boilingBoulders("input.txt", part2))
}

type coord struct {
	x, y, z int
}

func boilingBoulders(path string, part func(map[coord]struct{}) int) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	lavaScan := map[coord]struct{}{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var lava coord
		if _, err := fmt.Sscanf(scanner.Text(), "%d,%d,%d", &lava.x, &lava.y, &lava.z); err != nil {
			panic(err)
		}

		lavaScan[lava] = struct{}{}
	}

	return part(lavaScan)
}

func part2(lavaScan map[coord]struct{}) int {
	minBound := coord{x: math.MaxInt, y: math.MaxInt, z: math.MaxInt}
	maxBound := coord{x: math.MinInt, y: math.MinInt, z: math.MinInt}

	for lava := range lavaScan {
		minBound = coord{
			x: min(minBound.x, lava.x),
			y: min(minBound.y, lava.y),
			z: min(minBound.z, lava.z),
		}

		maxBound = coord{
			x: max(maxBound.x, lava.x),
			y: max(maxBound.y, lava.y),
			z: max(maxBound.z, lava.z),
		}
	}

	minBound = coord{minBound.x - 1, minBound.y - 1, minBound.z - 1}
	maxBound = coord{maxBound.x + 1, maxBound.y + 1, maxBound.z + 1}
	queue := []coord{minBound}
	seen := map[coord]struct{}{}
	ans := 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, neighbor := range curr.neighbors() {
			if _, ok := lavaScan[neighbor]; ok {
				ans++
			} else if _, ok := seen[neighbor]; !ok {
				if neighbor.x >= minBound.x && neighbor.y >= minBound.y && neighbor.z >= minBound.z &&
					neighbor.x <= maxBound.x && neighbor.y <= maxBound.y && neighbor.z <= maxBound.z {
					seen[neighbor] = struct{}{}
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return ans
}

func part1(lavaScan map[coord]struct{}) int {
	ans := 0

	for lava := range lavaScan {
		ans += 6
		for _, neighbor := range lava.neighbors() {
			if _, ok := lavaScan[neighbor]; ok {
				ans -= 1
			}
		}
	}

	return ans
}

func (c coord) neighbors() []coord {
	neighbors := make([]coord, 6)

	dirs := []coord{
		{1, 0, 0},
		{-1, 0, 0},
		{0, 1, 0},
		{0, -1, 0},
		{0, 0, 1},
		{0, 0, -1},
	}

	for i, dir := range dirs {
		neighbors[i] = coord{c.x + dir.x, c.y + dir.y, c.z + dir.z}
	}

	return neighbors
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
