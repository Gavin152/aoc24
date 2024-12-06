package main

import (
	"fmt"
	"slices"

	"github.com/Gavin152/aoc24/internal/filereader"
)

var directions map[rune][]int
var startPoint []int
var startDirection rune

func fillMatrix(lines []string) [][]rune {
	xlen := len(lines[0])
	ylen := len(lines)
	lab := make([][]rune, xlen)
	for i, _ := range lab {
		lab[i] = make([]rune, ylen)
	}

	fmt.Printf("lab length: %d, lab element length: %d\n", len(lab), len(lab[0]))
	for i, line := range lines {
		col := []rune(line)
		for j, character := range col {
			// fmt.Printf("i: %d, j: %d\n", i, j)
			lab[j][i] = character
			if slices.Index([]rune{'^', '>', 'v', '<'}, character) > -1 {
				startPoint = []int{j, i}
				startDirection = character
			}
		}
	}
	return lab
}

func walk(lab [][]rune) {
	current := startPoint
	direction := startDirection
	fmt.Printf("Startpoint is %d|%d, direction %c\n", current[0], current[1], direction)
	for {
		lab[current[0]][current[1]] = 'X'
		next := []int{
			current[0] + directions[direction][0],
			current[1] + directions[direction][1],
		}
		if isOutOfBounds(lab, next) {
			break
		}
		if lab[next[0]][next[1]] == '#' {
			direction = turn(direction)
			continue
		}
		current = next
	}
}

func turn(current rune) rune {
	if current == '^' {
		return '>'
	}
	if current == '>' {
		return 'v'
	}
	if current == 'v' {
		return '<'
	}
	return '^'

}

func isOutOfBounds(lab [][]rune, pos []int) bool {
	x := pos[0]
	y := pos[1]

	if x < 0 || x >= len(lab) {
		return true
	}
	if y < 0 || y >= len(lab[0]) {
		return true
	}
	return false
}

func tallyPositions(lab [][]rune) int {
	tally := 0
	for i, col := range lab {
		for j, _ := range col {
			if lab[i][j] == 'X' {
				tally++
			}
		}
	}
	return tally
}

func printLab(lab [][]rune) {
	for i, col := range lab {
		for j, _ := range col {
			fmt.Printf("%s", string(lab[j][i]))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func main() {

	directions = map[rune][]int{
		'^': []int{0, -1},
		'>': []int{1, 0},
		'v': []int{0, 1},
		'<': []int{-1, 0},
	}
	// filePath := "example.txt"
	filePath := "data_a"
	// filePath := "data_b"

	lines := []string{}

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		lines = append(lines, line)
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	lab := fillMatrix(lines)
	fmt.Printf("%s", string(lab[4]))
	walk(lab)

	total := tallyPositions(lab)
	printLab(lab)

	fmt.Printf("The guard visited %d positions\n", total)
}
