package util

import (
	"fmt"
	"strconv"
)

func SliceToGrid(lines []string) [][]rune {
	xlen := len(lines[0])
	ylen := len(lines)
	grid := make([][]rune, xlen)
	for i := range grid {
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

func ParseGridToInt(grid [][]rune) ([][]int, error) {
	gridInt := [][]int{}

	for i, col := range grid {
		intcol := []int{}
		for j := range col {
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

func PrintGrid[T any](grid [][]T) {
	for i, col := range grid {
		for j := range col {
			fmt.Printf("%c", grid[j][i])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func containsPoint(path [][]int, point []int) bool {
    for _, p := range path {
        if len(p) == len(point) && p[0] == point[0] && p[1] == point[1] {
            return true
        }
    }
    return false
}

func PrintGridWithPath[T any](grid [][]T, path [][]int) {
	for i, col := range grid {
		for j := range col {
			if containsPoint(path, []int{j, i}) {
				fmt.Printf("%c", 'O')
			} else {
				fmt.Printf("%c", grid[j][i])
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
