package main

import (
	"fmt"
	"strconv"
	s "strings"

	"github.com/Gavin152/aoc24/internal/filereader"
)

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func isSafe(report []int, retry bool) bool {
	previous := report[0]
	level := report[1]
	rising := true

	if previous > level {
		rising = false
	}

	for i, val := range report {
		if i == 0 {
			continue
		}
		level = val

		distance := level - previous
		if !rising {
			distance = distance * -1
		}

		if rising && level < previous {

		}

		if distance < 1 || distance > 3 {

		}

		previous = level
	}
	fmt.Printf("Safe report found: %v\n", report)
	return true
}

func main() {

	safeCount := 0

	filePath := "example.txt"
	// filePath := "data_a"
	// filePath := "data_b"

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		clean_line := s.TrimSpace(line)
		split_line := s.Split(clean_line, " ")

		reps := []int{}
		for _, val := range split_line {
			num, _ := strconv.Atoi(val)
			reps = append(reps, num)
		}

		if isSafe(reps, true) {
			// fmt.Printf("Safe report found: %v\n", split_line)
			safeCount++
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Found %d safe reports", safeCount)
}
