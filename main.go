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

func checkpattern(pattern string, cache map[string]int) int {

	val, ok := cache[pattern]
	if ok {
		return val
	}
	matches := 0
	if pattern == ""{
		// fmt.Println("Pattern matched")
		return 1
	}
	for _, towel := range towels {
		if towel == pattern {
			matches++
		} else if strings.HasPrefix(pattern, towel) {
			matches += checkpattern(pattern[len(towel):], cache)
		}
	}
	cache[pattern] = matches
	return matches
}

func checkpatterns() int {
	cache := map[string]int{}
	count := 0
	for _, pattern := range patterns {
		fmt.Printf("Checking pattern %s\n", pattern)
		patternMatches := checkpattern(pattern, cache)
		fmt.Printf("Pattern %s can be matched %d different ways\n", pattern, patternMatches)
		fmt.Println("================")
		count += patternMatches
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
