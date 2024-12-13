package util

import (
	"fmt"
	"strconv"
)

func SliceToGrid (lines []string) [][]rune {
	xlen := len(lines[0])
	ylen := len(lines)
	grid := make([][]rune, xlen)
	for i, _ := range grid {
		grid[i] = make([]rune, ylen)
	}

	for i, line := range lines {
		col := []rune(line)
		for j, character := range col {
			grid[j][i] = character
		}
	}
	return grid
}

func ParseGridToInt (grid [][]rune) ([][]int, error) {
	gridInt := [][]int{}
	
	for i, col := range grid {
		intcol := []int{}
		for j, _ := range col {
			newInt, err := strconv.Atoi(string(grid[j][i]))
			if err != nil {
				newInt = -1
			}
			intcol = append(intcol, newInt)
		}
		gridInt = append(gridInt, intcol)
	}
	return gridInt, nil
}

func PrintGrid [T any](grid [][]T) {
	for i, col := range grid {
		for j, _ := range col {
			fmt.Printf("%v", grid[j][i])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}