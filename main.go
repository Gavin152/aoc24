package main

import (
	"fmt"
	"strconv"
	s "strings"

	"github.com/Gavin152/aoc24/internal/filereader"
)

func isSafe(report []string) bool {
	previous, _ := strconv.Atoi(report[0])
	level, _ := strconv.Atoi(report[1])

	rising := true
	if previous > level {
		rising = false
	}

	for i, string_val := range report {
		if i == 0 {
			continue
		}
		level, _ = strconv.Atoi(string_val)

		// if level > previous && (level-previous < 1 || level-previous > 3) {
		// 	fmt.Printf("Round: %d | Distance: %d\n", i+1, level-previous)
		// 	return false
		// }
		// if previous > level && (previous-level < 1 || previous-level > 3) {
		// 	fmt.Printf("Round: %d | Distance: %d\n", i+1, previous-level)
		// 	return false
		// }
		// if level == previous {
		// 	fmt.Printf("Round: %d | Distance: %d\n", i+1, level-previous)
		// 	return false
		// }

		distance := level - previous
		if !rising {
			distance = distance * -1
		}

		if rising && level < previous {
			return false
		}

		if distance < 1 || distance > 3 {
			return false
		}

		previous = level
	}
	fmt.Printf("Safe report found: %v\n", report)
	return true
}

func main() {

	safeCount := 0

	// filePath := "example.txt"
	filePath := "data_a"
	// filePath := "data_b"

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		clean_line := s.TrimSpace(line)
		split_line := s.Split(clean_line, " ")

		if isSafe(split_line) {
			safeCount++
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Found %d safe reports", safeCount)
}
