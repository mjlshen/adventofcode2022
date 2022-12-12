package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", hillClimbing("input.txt", false))
	fmt.Printf("Part 2: %d\n", hillClimbing("input.txt", true))
}

type coord struct {
	x, y int
}

func hillClimbing(path string, allPossibleStarts bool) int32 {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	parsed := strings.Split(string(input), "\n")
	heightmap := make([][]int32, len(parsed))
	for i := range heightmap {
		heightmap[i] = make([]int32, len(parsed[0]))
	}
	var start, end coord
	var possibleStarts []coord

	for i, p := range parsed {
		for j, char := range p {
			var elevation int32
			if char == 'E' {
				end = coord{i, j}
				elevation = 'z' - 'a'
			} else if char == 'S' {
				start = coord{i, j}
				elevation = 'a' - 'a'
			} else {
				elevation = char - 'a'
			}

			if elevation == 0 {
				possibleStarts = append(possibleStarts, coord{i, j})
			}

			heightmap[i][j] = elevation
		}
	}

	if !allPossibleStarts {
		return bfs(heightmap, end, start)
	}

	min := int32(math.MaxInt32)
	for _, s := range possibleStarts {
		if actual := bfs(heightmap, end, s); actual < min && actual != -1 {
			min = actual
		}
	}
	return min
}

func bfs(heightmap [][]int32, start coord, target coord) int32 {
	queue := []coord{start}
	dist := make([][]int32, len(heightmap))
	for i := range dist {
		dist[i] = make([]int32, len(heightmap[0]))
	}

	for i := range dist {
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	dist[start.x][start.y] = 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == target {
			return dist[target.x][target.y]
		}

		directions := []coord{
			{1, 0},
			{-1, 0},
			{0, 1},
			{0, -1},
		}

		for _, direction := range directions {
			next := coord{direction.x + curr.x, direction.y + curr.y}

			// Out of bounds
			if next.x < 0 || next.x == len(heightmap) || next.y < 0 || next.y == len(heightmap[0]) {
				continue
			}

			// Already visited
			if dist[next.x][next.y] != -1 {
				continue
			}

			// Only allowed to step one elevation down at a time
			if heightmap[curr.x][curr.y]-heightmap[next.x][next.y] > 1 {
				continue
			}

			dist[next.x][next.y] = dist[curr.x][curr.y] + 1
			queue = append(queue, next)
		}
	}

	return -1
}
