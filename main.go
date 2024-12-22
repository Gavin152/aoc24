package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Gavin152/aoc24/internal/filereader"
)

var towels = []string{}
var patterns = []string{}

func parseInput(lines []string) {
	part2 := false
	for _, line := range lines {
		if line == "" {
			part2 = true
			continue
		}
		if part2 {
			patterns = append(patterns, line)
		} else {
			split := strings.Split(line, ", ")
			for _, towel := range split {
				towels = append(towels, strings.TrimSpace(towel))
			}
		}
	}
}

func checkpattern(pattern string) bool {
	isMatch := false
	if len(pattern) == 0 {
		return true
	}
	for _, towel := range towels {
		if isMatch {
			break
		}
		// fmt.Printf("Checking if pattern %s starts with %s\n", pattern, towel)
		if strings.HasPrefix(pattern, towel) {
			newPat := strings.TrimPrefix(pattern, towel)
			isMatch = checkpattern(newPat)
		} else {
			continue
		}
	}
	return isMatch
}

func checkpatterns() int {
	count := 0
	for _, pattern := range patterns {
		// fmt.Printf("Inspecting pattern %s\n", pattern)
		if checkpattern(pattern) {
			// fmt.Printf("Matching pattern found!\n\n")
			count++
		} else {
			// fmt.Printf("No matcches found for %s\n\n", pattern)
		}
	}
	return count
}

func main() {
	// filePath := "example"
	filePath := "data"

	var lines []string
	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		lines = append(lines, line)
		return nil
	})
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	parseInput(lines)
	slices.SortFunc(towels, func(a, b string) int {
		// Compare lengths in reverse order for descending sort
		return len(b) - len(a)
	})

	// fmt.Printf("Towles: %v\n", towels)
	// fmt.Printf("Patterns: %v\n", patterns)

	count := checkpatterns()

	fmt.Printf("%d matching patterns found", count)
}
