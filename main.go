package main

import (
	"fmt"
	"slices"

	"github.com/Gavin152/aoc24/internal/filereader"
	"github.com/Gavin152/aoc24/internal/util"
)

type Cheat struct {
	Start []int
	End   []int
}

func isInBounds(row, col int, grid [][]rune) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}

func cheatedAtPosition(row, col int, cheats *[]Cheat) bool {
	for _, cheat := range *cheats {
		if cheat.Start[0] == row && cheat.Start[1] == col {
			return true
		}
	}
	return false
}

func getPath(grid [][]rune, withCheat bool, cheats *[]Cheat) [][]int {
	path := [][]int{}
	cheated := false

	// Find starting point 'S'
	var start []int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'S' {
				start = []int{i, j}
				break
			}
		}
		if start != nil {
			break
		}
	}
	if start == nil {
		return path
	}

	// Add starting point
	current := start

	// Directions: up, right, down, left
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for {
		row, col := current[0], current[1]
		found := false

		// Try all directions
		for _, dir := range dirs {
			newRow, newCol := row+dir[0], col+dir[1]

			if withCheat && !cheated && !cheatedAtPosition(newRow, newCol, cheats) {
				// check we're not going out of bounds
				if !isInBounds(newRow, newCol, grid) {
					continue
				}

				cStartX := newRow
				cStartY := newCol
				cEndX := newRow + dir[0]
				cEndY := newCol + dir[1]

				if grid[cStartX][cStartY] == '#' &&
					isInBounds(cEndX, cEndY, grid) &&
					grid[cEndX][cEndY] == '.' {
					*cheats = append(*cheats, Cheat{Start: []int{cStartX, cStartY}, End: []int{cEndX, cEndY}})
					newRow = cEndX
					newCol = cEndY
					cheated = true
					path = append(path, []int{cStartX, cStartY})
					path = append(path, []int{cEndX, cEndY})
					continue
				}
			}

			// Check bounds
			if !isInBounds(newRow, newCol, grid) {
				continue
			}

			// Check if we found the end
			if grid[newRow][newCol] == 'E' {
				path = append(path, []int{newRow, newCol})
				return path
			}

			// Check if it's a path tile we haven't visited yet
			if grid[newRow][newCol] == '.' {
				// Check if this position is not already in our path
				isNew := true
				for _, pos := range path {
					if pos[0] == newRow && pos[1] == newCol {
						isNew = false
						break
					}
				}

				if isNew {
					current = []int{newRow, newCol}
					path = append(path, current)
					found = true
					break
				}
			}
		}

		// If we couldn't find a new position to move to, we're done
		if !found {
			break
		}
	}

	return path
}

func main() {
	filePath := "example"
	// filePath := "data"

	var lines []string
	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		lines = append(lines, line)
		return nil
	})
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	grid := util.SliceToGrid(lines)
	util.PrintGrid(grid)

	path := getPath(grid, false, nil)
	fmt.Printf("Path: %v\n", path)
	fmt.Printf("Path length: %v\n", len(path))

	cheats := []Cheat{}
	cheatedPaths := [][][]int{}
	for range path {
		cp := getPath(grid, true, &cheats)
		cheatedPaths = append(cheatedPaths, cp)
		// fmt.Printf("Found a cheated path of length %v at position %v\n", len(cp), p)
	}

	// fmt.Printf("Number of cheated paths: %v\n", len(cheatedPaths))

	slices.SortFunc(cheatedPaths, func(a, b [][]int) int {
		return len(a) - len(b)
	})

	for idx := 0; idx < 20; idx++ {
		// fmt.Printf("Cheat %d: %v\n", idx, cheats[idx])
		// fmt.Printf("Cheat %d path: %v\n", idx, cheatedPaths[idx])
		fmt.Printf("Cheat %d path length: %v\n", idx, len(cheatedPaths[idx]))
		fmt.Printf("A saving of %d\n\n", len(path) - len(cheatedPaths[idx]))
		
	}

	// grid[cheats[chIdx].Start[0]][cheats[chIdx].Start[1]] = '1'
	// grid[cheats[chIdx].End[0]][cheats[chIdx].End[1]] = '2'
	// util.PrintGrid(grid)
}
