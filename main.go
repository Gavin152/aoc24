package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/Gavin152/aoc24/internal/filereader"
)

func splitPages(raw *[]string, split *[][]int) {
	for _, line := range *raw {
		s := strings.Split(line, "|")
		first, _ := strconv.Atoi(s[0])
		second, _ := strconv.Atoi(s[1])
		*split = append(*split, []int{first, second})
	}
}

func splitUpdates(raw *[]string, split *[][]int) {
	for _, line := range *raw {
		s := strings.Split(line, ",")
		p := []int{}
		for _, num := range s {
			parsed, _ := strconv.Atoi(num)
			p = append(p, parsed)
		}
		*split = append(*split, p)
	}
}

func checkUpdates(updates *[][]int, rules *[][]int) int {
	sum := 0
	for _, update := range *updates {
		isValid := true
		for i, rule := range *rules {
			fmt.Printf("Checking Rule %d\n", i)
			u_index1 := slices.Index(update, rule[0])
			u_index2 := slices.Index(update, rule[1])
			if u_index1 == -1 || u_index2 == -1 {
				continue
			}
			if !(u_index1 < u_index2) {
				fmt.Printf("Break on %d\n", i)
				isValid = false
				break
			}
		}
		if isValid {
			sum += findCenter(&update)
		}
	}
	return sum
}

func findCenter(update *[]int) int {
	idx := (len(*update) - 1) / 2
	return (*update)[idx]
}

func main() {

	// safeCount := 0
	raw_pages := []string{}
	raw_updates := []string{}

	pages := [][]int{}
	updates := [][]int{}

	readUpdates := false

	// filePath := "example.txt"
	filePath := "data_a"
	// filePath := "data_b"

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		if line == "" {
			readUpdates = true
		}
		if !readUpdates {
			raw_pages = append(raw_pages, line)
		} else {
			raw_updates = append(raw_updates, line)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	splitPages(&raw_pages, &pages)
	splitUpdates(&raw_updates, &updates)

	tally := checkUpdates(&updates, &pages)

	// fmt.Printf("%v\n", pages)
	// fmt.Printf("%v\n", updates)
	fmt.Printf("Total tally of valid updates' center pages is %d", tally)
}
