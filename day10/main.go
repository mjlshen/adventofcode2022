package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", cathodeRayTube("input.txt"))
}

type clock struct {
	x      int
	cycle  int
	future map[int]int
	score  int
	crt    [240]bool
}

func (c *clock) increment() {
	c.cycle++

	if c.cycle == 20 || (c.cycle-20)%40 == 0 {
		c.score += c.cycle * c.x
	}

	if c.cycle-1 < 240 {
		if abs((c.cycle-1)%40-c.x) <= 1 {
			c.crt[c.cycle-1] = true
		} else {
			c.crt[c.cycle-1] = false
		}
	}

	if v, ok := c.future[c.cycle]; ok {
		c.x += v
		delete(c.future, c.cycle)
	}
}

func (c *clock) draw() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			if c.crt[i*40+j] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func cathodeRayTube(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	c := &clock{
		x:      1,
		cycle:  0,
		future: map[int]int{},
		score:  0,
		crt:    [240]bool{},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var instruction string
		var value int

		if scanner.Text() == "noop" {
			c.increment()
			continue
		}

		if _, err := fmt.Sscanf(scanner.Text(), "%s %d", &instruction, &value); err != nil {
			panic(err)
		}

		c.future[c.cycle+2] = value

		c.increment()
		c.increment()
	}

	for i := 0; i < 2; i++ {
		c.increment()
	}

	fmt.Println("Part 2:")
	c.draw()
	return c.score
}
