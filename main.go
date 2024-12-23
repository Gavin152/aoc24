package main

import (
	"fmt"
	_ "slices"

	"github.com/Gavin152/aoc24/internal/filereader"
	"github.com/Gavin152/aoc24/internal/util"
)

type Position struct {
	X int
	Y int
}

type Cheat struct {
	Start Position
	End   Position
	Path  [][]int
}

func isInBounds(row, col int, grid [][]rune) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}

func getPath(grid [][]rune) ([][]int, []int) {
	path := [][]int{}

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
		return path, start
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

			// Check bounds
			if !isInBounds(newRow, newCol, grid) {
				continue
			}

			// Check if we found the end
			if grid[newRow][newCol] == 'E' {
				path = append(path, []int{newRow, newCol})
				return path, start
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

	return path, start
}

func cheat(grid [][]rune, path [][]int) []Cheat {
	cheats := []Cheat{}
	for _, pos := range path {
		cheat, ok := cheatAtPosition(grid, path, Position{X: pos[0], Y: pos[1]})
		if ok {
			cheats = append(cheats, cheat...)
		}
		// if i == 17 {
		// 	util.PrintGridWithPath(grid, cheat[0].Path)
		// }
	}
	return cheats
}

func cheatAtPosition(grid [][]rune, path [][]int, position Position) ([]Cheat, bool) {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	cheats := []Cheat{}

	// Find position index in the original path
	posIdx := -1
	for i, p := range path {
		if p[0] == position.X && p[1] == position.Y {
			posIdx = i
			break
		}
	}
	if posIdx == -1 {
		return nil, false
	}

	for _, dir := range dirs {
		cStart := Position{X: position.X + dir[0], Y: position.Y + dir[1]}
		cEnd := Position{X: cStart.X + dir[0], Y: cStart.Y + dir[1]}

		// Find the end position in the original path
		cEndIdx := -1
		for i, p := range path {
			if p[0] == cEnd.X && p[1] == cEnd.Y {
				cEndIdx = i
				break
			}
		}

		if grid[cStart.X][cStart.Y] == '#' &&
			isInBounds(cEnd.X, cEnd.Y, grid) &&
			(grid[cEnd.X][cEnd.Y] == '.' || grid[cEnd.X][cEnd.Y] == 'E') &&
			cEndIdx > posIdx {
			
			// Create new path slices to avoid modifying the original
			path1 := make([][]int, posIdx+1)
			copy(path1, path[:posIdx+1])
			
			path2 := make([][]int, len(path)-cEndIdx)
			copy(path2, path[cEndIdx:])

			// Create the cheated path
			cPath := make([][]int, 0, len(path1)+1+len(path2))
			cPath = append(cPath, path1...)
			cPath = append(cPath, []int{cStart.X, cStart.Y})
			cPath = append(cPath, path2...)

			cheats = append(cheats, Cheat{Start: cStart, End: cEnd, Path: cPath})
		}
	}

	return cheats, len(cheats) > 0
}

func main() {
	// filePath := "example"
	filePath := "data"

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

	path, start := getPath(grid)
	// fmt.Printf("Path: %v\n", path)
	fmt.Printf("Path length: %v\n", len(path))
	// util.PrintGridWithPath(grid, path)

	cheats := []Cheat{}

	c1, c1Found := cheatAtPosition(grid, path, Position{X: start[0], Y: start[1]})
	if c1Found {
		fmt.Println("Found cheat at start")
		cheats = append(cheats, c1...)
	}

	cheats = append(cheats,cheat(grid, path)...)
	saved100 := 0
	for _, cheat := range cheats {
		if len(path) - len(cheat.Path) > 99 {
			// fmt.Printf("\nFound cheat if length: %d ... a saving of: %d\n", len(cheat.Path), len(path)-len(cheat.Path))
			// fmt.Printf("Start: %v\n", cheat.Start)
			// fmt.Printf("End: %v\n", cheat.End)
			// fmt.Printf("Path: %v\n", cheat.Path)
			// util.PrintGridWithPath(grid, cheat.Path)
			saved100++
		}
		// fmt.Printf("Start: %v\n", cheat.Start)
		// fmt.Printf("End: %v\n", cheat.End
		// fmt.Printf("Path: %v\n", cheat.Path)
		// util.PrintGridWithPath(grid, cheat.Path)
	}

	fmt.Printf("\n==============================\n")
	fmt.Printf("Cheats that save at least 100: %d\n", saved100)
	fmt.Printf("Total cheats: %d\n", len(cheats))
}
