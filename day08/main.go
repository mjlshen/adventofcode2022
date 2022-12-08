package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", visibleTrees("input.txt", countVisibleTrees))
	fmt.Printf("Part 2: %d\n", visibleTrees("input.txt", mostScenicTree))
}

func visibleTrees(path string, scoring func([][]int) int) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make([]int, len(scanner.Text()))
		for i, t := range scanner.Text() {
			row[i] = int(t - '0')
		}
		grid = append(grid, row)
	}

	return scoring(grid)
}

func countVisibleTrees(grid [][]int) int {
	// All trees around the edge are visible
	count := len(grid)*2 + len(grid[0])*2 - 4

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			// Sweep left
			visible := true
			for k := i - 1; k >= 0; k-- {
				if grid[i][j] <= grid[k][j] {
					visible = false
					break
				}
			}
			if visible {
				count++
				continue
			}

			// Sweep right
			visible = true
			for k := i + 1; k < len(grid); k++ {
				if grid[i][j] <= grid[k][j] {
					visible = false
					break
				}
			}

			if visible {
				count++
				continue
			}

			// Sweep up
			visible = true
			for k := j - 1; k >= 0; k-- {
				if grid[i][j] <= grid[i][k] {
					visible = false
					break
				}
			}
			if visible {
				count++
				continue
			}

			// Sweep down
			visible = true
			for k := j + 1; k < len(grid[i]); k++ {
				if grid[i][j] <= grid[i][k] {
					visible = false
					break
				}
			}

			if visible {
				count++
				continue
			}
		}
	}

	return count
}

func mostScenicTree(grid [][]int) int {
	maxScore := 0

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			// Sweep left
			left := 0
			for k := i - 1; k >= 0; k-- {
				left++
				if grid[i][j] <= grid[k][j] {
					break
				}
			}

			// Sweep right
			right := 0
			for k := i + 1; k < len(grid); k++ {
				right++
				if grid[i][j] <= grid[k][j] {
					break
				}
			}

			// Sweep up
			up := 0
			for k := j - 1; k >= 0; k-- {
				up++
				if grid[i][j] <= grid[i][k] {
					break
				}
			}

			// Sweep down
			down := 0
			for k := j + 1; k < len(grid[i]); k++ {
				down++
				if grid[i][j] <= grid[i][k] {
					break
				}
			}

			score := left * right * up * down
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}
