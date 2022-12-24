package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", unstableDiffusion("input.txt", 1))
	fmt.Printf("Part 2: %d\n", unstableDiffusion("input.txt", 2))
}

var directions = []coord{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

type coord struct {
	x, y int
}

type grove struct {
	minX, maxX, minY, maxY int
	elves                  map[coord]struct{}
	d                      int
	stable                 bool
}

func unstableDiffusion(path string, part int) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	y := 0
	elves := map[coord]struct{}{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for x, r := range scanner.Text() {
			switch r {
			case '.':
				continue
			case '#':
				elves[coord{x, y}] = struct{}{}
			}
		}
		y++
	}

	g := &grove{
		maxX:  math.MinInt,
		minX:  math.MaxInt,
		maxY:  math.MinInt,
		minY:  math.MaxInt,
		elves: elves,
		d:     0,
	}

	for elf := range g.elves {
		if elf.x > g.maxX {
			g.maxX = elf.x
		}

		if elf.x < g.minX {
			g.minX = elf.x
		}

		if elf.y > g.maxY {
			g.maxY = elf.y
		}

		if elf.y < g.minY {
			g.minY = elf.y
		}
	}

	switch part {
	case 1:
		return g.part1()
	default:
		return g.part2()
	}
}

func (g *grove) part1() int {
	for i := 0; i < 10; i++ {
		g.round()
	}

	for elf := range g.elves {
		if elf.x > g.maxX {
			g.maxX = elf.x
		}

		if elf.x < g.minX {
			g.minX = elf.x
		}

		if elf.y > g.maxY {
			g.maxY = elf.y
		}

		if elf.y < g.minY {
			g.minY = elf.y
		}
	}

	return (g.maxX-g.minX+1)*(g.maxY-g.minY+1) - len(g.elves)
}

func (g *grove) part2() int {
	for i := 1; ; i++ {
		g.round()
		if g.stable {
			return i
		}
	}
}

func (g *grove) round() {
	proposed := map[coord]map[coord]struct{}{}
	moved := map[coord]struct{}{}
	elves := g.elves

	for i := 0; i < len(directions); i++ {
		for elf := range elves {
			if g.noNeighbors(elf) {
				moved[elf] = struct{}{}
			}

			if _, ok := moved[elf]; !ok {
				pLoc := add(elf, directions[(g.d+i)%4])
				switch (g.d + i) % 4 {
				case 0, 1:
					if g.empty([]coord{add(pLoc, coord{1, 0}), pLoc, add(pLoc, coord{-1, 0})}...) {
						if _, ok := proposed[pLoc]; !ok {
							proposed[pLoc] = map[coord]struct{}{}
						}
						proposed[pLoc][elf] = struct{}{}
						moved[elf] = struct{}{}
						continue
					}
				case 2, 3:
					if g.empty([]coord{add(pLoc, coord{0, 1}), pLoc, add(pLoc, coord{0, -1})}...) {
						if _, ok := proposed[pLoc]; !ok {
							proposed[pLoc] = map[coord]struct{}{}
						}
						proposed[pLoc][elf] = struct{}{}
						moved[elf] = struct{}{}
						continue
					}
				}
			}
		}
	}

	stable := true
	for p, elves := range proposed {
		if len(elves) > 1 {
			continue
		}

		for elf := range elves {
			stable = false
			delete(g.elves, elf)
			g.elves[p] = struct{}{}
		}
	}
	g.stable = stable

	g.d++
}

func (g *grove) empty(cs ...coord) bool {
	for _, c := range cs {
		if _, ok := g.elves[c]; ok {
			return false
		}
	}

	return true
}

func (g *grove) noNeighbors(elf coord) bool {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if _, ok := g.elves[coord{elf.x + i, elf.y + j}]; ok {
				return false
			}
		}
	}
	return true
}

func add(a, b coord) coord {
	return coord{a.x + b.x, a.y + b.y}
}

func (g *grove) String() string {
	var res string
	for j := g.minY; j <= g.maxY; j++ {
		for i := g.minX; i <= g.maxX; i++ {
			if _, ok := g.elves[coord{i, j}]; ok {
				res += "#"
			} else {
				res += "."
			}
		}
		res += "\n"
	}

	return res
}
