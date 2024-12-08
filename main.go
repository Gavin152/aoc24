package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Gavin152/aoc24/internal/filereader"
)

var grid [][]rune
var antinodes map[string]bool

func linesToGrid(lines []string) [][]rune {
	// First dimension is Y (rows), second is X (columns)
	height := len(lines)
	width := len(lines[0]) // Assumes all lines have same width
	
	grid := make([][]rune, height)
	for y := range grid {
		grid[y] = make([]rune, width)
		for x := range lines[y] {
			grid[y][x] = rune(lines[y][x])
		}
	}
	return grid
}

func fillMatrix(lines []string) {
	xlen := len(lines[0])
	ylen := len(lines)
	grid = make([][]rune, xlen)
	for i, _ := range grid {
		grid[i] = make([]rune, ylen)
	}

	// fmt.Printf("lab length: %d, lab element length: %d\n", len(lab), len(lab[0]))
	for i, line := range lines {
		col := []rune(line)
		for j, character := range col {
			// fmt.Printf("i: %d, j: %d\n", i, j)
			grid[j][i] = character
		}
	}
}

func printGrid() {
	anodes := [][]int{}
	// fmt.Printf("Antinodes: %v\n", antinodes)
	for k := range antinodes {
		coords := strings.Split(k, "|")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		anodes = append(anodes, []int{x, y})
	}
	fmt.Printf("\n")
	for i, col := range grid {
		for j, _ := range col {
			if containsCoord(anodes, []int{j, i}) {
				fmt.Printf("#")
			} else {
				fmt.Printf("%s", string(grid[j][i]))
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func isInBounds(coord []int) bool {
	return coord[0] >= 0 && coord[0] < len(grid[0]) && coord[1] >= 0 && coord[1] < len(grid)
}

func getAntinodes(coord []int) [][]int {
	anodes := [][]int{}
	character := grid[coord[0]][coord[1]]
	for x, row := range grid {
		for y, _ := range row {
			if grid[x][y] == character && !(x == coord[0] && y == coord[1]) {
				vecX := x - coord[0]
				vecY := y - coord[1]
				newNode := []int{coord[0]-vecX, coord[1]-vecY}
				for isInBounds(newNode) {
					anodes = append(anodes, newNode)
					newNode[0] -= vecX
					newNode[1] -= vecY
				}
			}
		}
	}

	// fmt.Printf("================================\n")
	return anodes
}

func containsCoord(anodes [][]int, coord []int) bool {
	for _, anode := range anodes {
		if anode[0] == coord[0] && anode[1] == coord[1] {
			return true
		}
	}
	return false
}

func main() {

	filePath := "example.txt"
	// filePath := "data"

	antinodes = make(map[string]bool)

	lines := []string{}
	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		lines = append(lines, line)
		grid = append(grid, []rune(line))
		
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fillMatrix(lines)
	for x, row := range grid {
		for y, _ := range row {
			if grid[x][y] != '.' {
				anodes := getAntinodes([]int{x, y})
				for _, anode := range anodes {
					antinodes[fmt.Sprintf("%d|%d", anode[0], anode[1])] = true
				}
			}
		}
	}


	fmt.Printf("Antinodes: %v\n", len(antinodes))
	// fmt.Printf("Antinodes:asdfa %v\n\n", antinodes)
	printGrid()
}