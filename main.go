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

func findInvalidIndex(report []int) int {
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
		if distance < 0 {
			distance = distance * -1
		}

		isSequential := (rising && level > previous) || (!rising && previous > level)
		isValid := 1 <= distance && distance <= 3
		if !isSequential || !isValid {
			return i
		}
		previous = level
	}
	// fmt.Printf("Safe report found: %v\n", report)
	return -1
}

func isSafe(report []int) bool {
	invalidIndex := findInvalidIndex(report)
	if invalidIndex < 0 {
		fmt.Printf("Valid report: %v\n", report)
		return true
	}
	index1, index2 := invalidIndex-1, invalidIndex
	swappedReport1 := deleteLevelAt(report, index1)
	swappedReport2 := deleteLevelAt(report, index2)
	swap1 := findInvalidIndex(swappedReport1) == -1
	swap2 := findInvalidIndex(swappedReport2) == -1
	if swap1 {
		fmt.Printf("valid on retry %v\n", swappedReport1)
		return true
	}
	if swap2 {
		fmt.Printf("valid on retry %v\n", swappedReport2)
		return true
	}
	return false
}

func deleteLevelAt(levels []int, idx int) []int {
	deleted := make([]int, len(levels)-1)
	copy(deleted[:idx], levels[:idx])
	copy(deleted[idx:], levels[idx+1:])
	return deleted
}

func main() {

	safeCount := 0

	// filePath := "example.txt"
	filePath := "data_a"
	// filePath := "data_b"

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		clean_line := s.TrimSpace(line)
		split_line := s.Split(clean_line, " ")

		reps := []int{}
		for _, val := range split_line {
			num, _ := strconv.Atoi(val)
			reps = append(reps, num)
		}

		if isSafe(reps) {
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
