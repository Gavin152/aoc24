package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/Gavin152/aoc24/internal/filereader"
)

var directions map[rune][]int
var startPoint []int
var startDirection rune
var lab [][]rune
var visited map[string][]rune

func fillMatrix(lines []string) {
	xlen := len(lines[0])
	ylen := len(lines)
	lab = make([][]rune, xlen)
	for i, _ := range lab {
		lab[i] = make([]rune, ylen)
	}

	// fmt.Printf("lab length: %d, lab element length: %d\n", len(lab), len(lab[0]))
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
}

func walk() bool {
	current := startPoint
	direction := startDirection
	// fmt.Printf("Startpoint is %d|%d, direction %c\n", current[0], current[1], direction)
	for {
		lab[current[0]][current[1]] = 'X'
		isLoop := setMarker(current, direction)
		if isLoop {
			return true
		}
		next := []int{
			current[0] + directions[direction][0],
			current[1] + directions[direction][1],
		}
		if isOutOfBounds(lab, next) {
			break
		}
		if lab[next[0]][next[1]] == '#' || lab[next[0]][next[1]] == 'O' {
			direction = turn(direction)
			continue
		}
		current = next
	}
	return false
}

func setMarker(pos []int, direction rune) bool {
	ps := strconv.Itoa(pos[0]) + "|" + strconv.Itoa(pos[1])
	dirs, exists := visited[ps]
	if exists && slices.Index(dirs, direction) > -1 {
		fmt.Printf("Already travelled this field in this direction\n")
		return true
	}
	visited[ps] = append(visited[ps], direction)
	return false
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

func tallyPositions() [][]int {
	tally := [][]int{}
	for i, col := range lab {
		for j, _ := range col {
			if lab[i][j] == 'X' {
				tally = append(tally, []int{i, j})
			}
		}
	}
	return tally
}

func printLab() {
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
	visited = map[string][]rune{}

	possibleObs := [][]int{}
	confirmedObs := [][]int{}

	// filePath := "example.txt"
	filePath := "data"

	lines := []string{}

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		lines = append(lines, line)
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fillMatrix(lines)
	walk()
	possibleObs = tallyPositions()

	for _, pos := range possibleObs {
		fillMatrix(lines)
		for k := range visited {
			delete(visited, k)
		}
		lab[pos[0]][pos[1]] = 'O'
		isLooped := walk()
		if isLooped {
			confirmedObs = append(confirmedObs, pos)
		}
		lab[pos[0]][pos[1]] = 'X'
	}

	total := tallyPositions()
	printLab()

	fmt.Printf("The guard visited %d positions\n", total)
	fmt.Printf("The guard visited %d positions\n", len(visited))

	fmt.Printf("Possible Obstruction Positions %d\n", len(possibleObs))
	fmt.Printf("Confirmed Obstruction Positions %d\n", len(confirmedObs))

}
