package main

import (
	"fmt"

	"github.com/Gavin152/aoc24/internal/filereader"
)

var lines []string
var directions [][]int

func findWords() int {
	xmasCounter := 0

	for i, line := range lines {
		for j, rawChar := range line {
			char := string(rawChar)

			if char == "X" {
				for _, vec := range directions {
					if checkDirection([]int{i, j}, vec) {
						xmasCounter++
						fmt.Printf("COUNTER: %d\n", xmasCounter)
					}
				}
			}
		}
	}

	return xmasCounter
}

func checkDirection(anchor []int, vector []int) bool {
	fmt.Printf("checking X in [%d,%d] in dir %v\n", anchor[0], anchor[1], vector)
	var h int
	var v int
	for i := 1; i < 4; i++ {
		h = anchor[1] + vector[1]*i
		v = anchor[0] + vector[0]*i
		if !isInBounds([]int{v, h}) {
			return false
		} else {
			char := string(lines[v][h])
			// fmt.Printf("i: %d ", i)
			// fmt.Printf("| char: %s\n", char)
			if i == 1 && char == "M" {
				continue
			} else if i == 2 && char == "A" {
				continue
			} else if i == 3 && char == "S" {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func isInBounds(pos []int) bool {
	// fmt.Printf("Checking [%d,%d]\n", pos[1], pos[0])
	if pos[0] < 0 || pos[0] > len(lines[0])-1 {
		fmt.Printf("[%d,%d] is out of bounds\n", pos[1], pos[0])
		return false
	}
	if pos[1] < 0 || pos[1] > len(lines)-1 {
		fmt.Printf("[%d,%d] is out of bounds\n", pos[1], pos[0])
		return false
	}
	return true
}

func main() {

	// safeCount := 0
	lines = []string{}
	directions = [][]int{
		{-1, 0},  // up
		{-1, 1},  // up right
		{0, 1},   // right
		{1, 1},   // btm right
		{1, 0},   // btm
		{1, -1},  // btm left
		{0, -1},  // left
		{-1, -1}, // up
	}

	// filePath := "example.txt"
	filePath := "data_a"
	// filePath := "data_b"

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		lines = append(lines, line)
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	xmass := findWords()

	// fmt.Printf("%c\n", lines[0][3])
	fmt.Printf("Found %d words", xmass)
}
