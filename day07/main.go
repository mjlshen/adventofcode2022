package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", clearingSpace("input.txt", part1))
	fmt.Printf("Part 2: %d\n", clearingSpace("input.txt", part2))
}

func clearingSpace(input string, part func(map[string]int) int) int {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// pwd will change over time to reflect the pwd in the VM
	pwd := ""
	// filesystem is map of filepaths absolute paths with values being their size
	filesystem := map[string]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "$ cd") {
			var destDir string
			if _, err := fmt.Sscanf(scanner.Text(), "$ cd %s", &destDir); err != nil {
				panic(err)
			}
			// path.Join intelligently handles ".." as well, which is convenient
			pwd = path.Join(pwd, destDir)
			continue
		}

		if scanner.Text() == "$ ls" || strings.HasPrefix(scanner.Text(), "dir ") {
			continue
		}

		var filesize int
		var filename string
		if _, err := fmt.Sscanf(scanner.Text(), "%d %s", &filesize, &filename); err != nil {
			panic(err)
		}
		filesystem[path.Join(pwd, filename)] = filesize
	}

	return part(sizeOfDirectories(filesystem))
}

// sizeOfDirectories converts a filesystem map into a map of directories with the key as the name of the directory
// and the value being the total size of its contents
func sizeOfDirectories(filesystem map[string]int) map[string]int {
	dirs := map[string]int{}

	for fpath, size := range filesystem {
		// For each file, add its size to all of its parents.
		// E.g. /a/b/c.txt would add the size of c.txt to all the directories /, /a and /a/b
		for dir := path.Dir(fpath); ; dir = path.Dir(dir) {
			dirs[dir] += size
			if dir == "/" {
				break
			}
		}
	}

	return dirs
}

// part1 returns the sum of the sizes of all directories with size < 100000
func part1(dirs map[string]int) int {
	ans := 0

	for _, size := range dirs {
		if size < 100000 {
			ans += size
		}
	}

	return ans
}

// part2 returns the smallest directory, when deleted, would bring the free space to at least
// 30000000 assuming a filesystem of total size 70000000
func part2(dirs map[string]int) int {
	targetFree := 30000000 - (70000000 - dirs["/"])
	currentMin := 70000000

	for _, size := range dirs {
		if size > targetFree && size < currentMin {
			currentMin = size
		}

	}

	return currentMin
}
