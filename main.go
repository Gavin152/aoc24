package main

import (
	"fmt"
	"strings"

	"github.com/Gavin152/aoc24/internal/filereader"
)

type Position struct {
	x, y int
}

func (p Position) add(other Position) Position {
	return Position{p.x + other.x, p.y + other.y}
}

type Warehouse struct {
	grid     [][]rune
	robot    Position
	boxes    []Position
	width    int
	height   int
	moves    []rune
}

func parseWarehouse(input []string) (*Warehouse, error) {
	var warehouse Warehouse
	var mapLines []string
	var moveStr string
	parsingMap := true

	// Separate map and moves
	for _, line := range input {
		if len(strings.TrimSpace(line)) == 0 {
			parsingMap = false
			continue
		}
		if parsingMap {
			if strings.Contains(line, "#") {
				mapLines = append(mapLines, line)
			}
		} else {
			// Only include actual movement characters
			for _, ch := range line {
				if ch == '^' || ch == 'v' || ch == '<' || ch == '>' {
					moveStr += string(ch)
				}
			}
		}
	}

	// Parse map
	warehouse.height = len(mapLines)
	if warehouse.height == 0 {
		return nil, fmt.Errorf("no map data found")
	}
	warehouse.width = len(mapLines[0])
	warehouse.grid = make([][]rune, warehouse.height)

	for y, line := range mapLines {
		warehouse.grid[y] = make([]rune, warehouse.width)
		for x, char := range line {
			warehouse.grid[y][x] = char
			switch char {
			case '@':
				warehouse.robot = Position{x, y}
				warehouse.grid[y][x] = '.' // Replace robot with empty space in grid
			case 'O':
				warehouse.boxes = append(warehouse.boxes, Position{x, y})
			}
		}
	}

	// Parse moves
	warehouse.moves = []rune(moveStr)

	return &warehouse, nil
}

func (w *Warehouse) hasBox(pos Position) bool {
	for _, box := range w.boxes {
		if box == pos {
			return true
		}
	}
	return false
}

func (w *Warehouse) moveBox(from, to Position) bool {
	moved := true
	if w.hasBox(to) {
		dir := Position{to.x - from.x, to.y - from.y}
		newPos := Position{to.x + dir.x, to.y + dir.y}
		if w.isValidPosition(newPos) && !w.isWall(newPos) {
			moved = w.moveBox(to, newPos)
		}else {
			moved = false
		}
	}
	if moved && !w.isWall(to) {
		for i, box := range w.boxes {
			if box == from {
				w.grid[from.y][from.x] = '.'
				w.boxes[i] = to
				return true
			}
		}
	}
	return false
}

func (w *Warehouse) isWall(pos Position) bool {

	if w.grid[pos.y][pos.x] == '#' {
		return true
	}
	return false
}

func (w *Warehouse) isValidPosition(pos Position) bool {
	return pos.y >= 0 && pos.y < w.height && pos.x >= 0 && pos.x < w.width
}

func (w *Warehouse) move(direction rune) bool {
	var delta Position
	switch direction {
	case '^':
		delta = Position{0, -1}
	case 'v':
		delta = Position{0, 1}
	case '<':
		delta = Position{-1, 0}
	case '>':
		delta = Position{1, 0}
	default:
		return false
	}

	newPos := w.robot.add(delta)
	if !w.isValidPosition(newPos) || w.isWall(newPos) {
		return false
	}

	if w.hasBox(newPos) {
		// Try to push the box
		newBoxPos := newPos.add(delta)
		if !w.moveBox(newPos, newBoxPos) {
			return false
		}
	}

	w.robot = newPos
	return true
}

func (w *Warehouse) calculateGPSSum() int {
	sum := 0
	for _, box := range w.boxes {
		// For each box:
		// - Multiply its distance from top (y) by 100
		// - Add its distance from left (x)
		gpsCoordinate := (box.y * 100) + box.x
		sum += gpsCoordinate
	}
	return sum
}

func (w *Warehouse) executeAllMoves() {
	fmt.Printf("\nExecuting %d moves:\n", len(w.moves))
	fmt.Printf("Move sequence: %s\n\n", string(w.moves))
	
	for _, move := range w.moves {
		w.move(move)
	}
	
	fmt.Printf("\nFinal GPS sum: %d\n", w.calculateGPSSum())
}

func (w *Warehouse) String() string {
	result := make([][]rune, w.height)
	for y := 0; y < w.height; y++ {
		result[y] = make([]rune, w.width)
		copy(result[y], w.grid[y])
	}

	// Place boxes
	for _, box := range w.boxes {
		result[box.y][box.x] = 'O'
	}

	// Place robot
	result[w.robot.y][w.robot.x] = '@'

	var sb strings.Builder
	for _, row := range result {
		sb.WriteString(string(row))
		sb.WriteRune('\n')
	}
	return sb.String()
}

func main() {
	filePath := "data"
	// filePath := "example"

	var lines []string
	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		lines = append(lines, line)
		return nil
	})
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	warehouse, err := parseWarehouse(lines)
	if err != nil {
		fmt.Printf("Error parsing warehouse: %v\n", err)
		return
	}

	fmt.Println("Initial state:")
	fmt.Print(warehouse)

	warehouse.executeAllMoves()

	fmt.Println("\nFinal state:")
	fmt.Print(warehouse)
}
