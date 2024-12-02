package main

import (
	"fmt"
	"slices"
	"strconv"
	s "strings"

	"github.com/Gavin152/aoc24/internal/filereader"
)

func countOccurance(n int, pool []int) int {
	count := 0
	for _, v := range pool {
		if v == n {
			count++
		}
	}
	return count
}

func main() {

	left := []int{}
	right := []int{}
	tally := 0

	// filePath := "example.txt"
	filePath := "data_a"
	// filePath := "data_b"

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		clean_line := s.TrimSpace(line)
		split_line := s.Split(clean_line, " ")

		left_number, l_err := strconv.Atoi(split_line[0])
		right_number, r_err := strconv.Atoi(split_line[len(split_line)-1])

		if l_err != nil || r_err != nil {
			fmt.Printf("Error parsing line: \nLEFT: %v\n RIGHT:%v\n", l_err, r_err)
		}

		left = append(left, left_number)
		right = append(right, right_number)
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	slices.Sort(left)
	slices.Sort(right)

	for _, val := range left {
		tally += val * countOccurance(val, right)
	}

	fmt.Printf("Final tally is %d", tally)
}
