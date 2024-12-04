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

			if char == "A" {
				masCnt := 0
				skip := false
				for _, vec := range directions {
					if checkDirection([]int{i, j}, vec) {
						masCnt++
					}
					if masCnt == 2 && !skip {
						fmt.Printf("Valid cross: [%d,%d]\n", i+1, j+1)
						xmasCounter++
						skip = true
					}
				}
			}
		}
	}

	return xmasCounter
}

func checkDirection(anchor []int, vector []int) bool {
	// fmt.Printf("checking X in [%d,%d] in dir %v\n", anchor[0], anchor[1], vector)
	// var h int
	// var v int
	h := anchor[1] + vector[1]
	v := anchor[0] + vector[0]

	hinv := anchor[1] + vector[1]*-1
	vinv := anchor[0] + vector[0]*-1

	if !isInBounds([]int{v, h}) || !isInBounds([]int{vinv, hinv}) {
		return false
	} else {
		char := string(lines[v][h])
		charinv := string(lines[vinv][hinv])
		if char == "M" && charinv == "S" {
			return true
		} else {
			return false
		}
	}
}

func isInBounds(pos []int) bool {
	// fmt.Printf("Checking [%d,%d]\n", pos[1], pos[0])
	if pos[0] < 0 || pos[0] > len(lines[0])-1 {
		// fmt.Printf("[%d,%d] is out of bounds\n", pos[1], pos[0])
		return false
	}
	if pos[1] < 0 || pos[1] > len(lines)-1 {
		// fmt.Printf("[%d,%d] is out of bounds\n", pos[1], pos[0])
		return false
	}
	return true
}

func main() {

	// safeCount := 0
	lines = []string{}
	directions = [][]int{
		{-1, 1},  // up right
		{1, 1},   // btm right
		{1, -1},  // btm left
		{-1, -1}, // up left
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
