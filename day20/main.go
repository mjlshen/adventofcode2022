package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", grovePositioningSystem("input.txt", 1, 1))
	fmt.Printf("Part 2: %d\n", grovePositioningSystem("input.txt", decryptionKey, 10))
}

const decryptionKey = 811589153

type element struct {
	index int
	value int
}

func grovePositioningSystem(path string, decryptionKey int, times int) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var values []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var v int
		if _, err := fmt.Sscanf(scanner.Text(), "%d", &v); err != nil {
			panic(err)
		}

		values = append(values, v*decryptionKey)
	}

	return mix(values, times)
}

func mix(values []int, times int) int {
	positions := make(map[element]*ring.Ring, len(values))
	zero := element{}
	numbers := ring.New(len(values))

	for i, v := range values {
		if v == 0 {
			// Save the position of the element with value 0
			zero = element{i, 0}
		}

		positions[element{i, v}] = numbers
		numbers.Value = v
		numbers = numbers.Next()
	}

	for i := 0; i < times; i++ {
		for j, v := range values {
			prev := positions[element{j, v}].Prev()
			curr := prev.Unlink(1)

			shift := v
			currLen := len(values) - 1

			if (shift > currLen/2) || (shift < -currLen/2) {
				shift %= currLen
				switch {
				case shift > currLen/2:
					shift -= currLen
				case shift < -currLen/2:
					shift += currLen
				}
			}

			prev.Move(shift).Link(curr)
		}
	}

	return ans(positions[zero])
}

func ans(r *ring.Ring) int {
	ans := 0
	for i := 0; i < 3; i++ {
		r = r.Move(1000)
		ans += r.Value.(int)
	}

	return ans
}
