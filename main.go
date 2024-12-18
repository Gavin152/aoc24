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

type BoxPosition struct {
	left, right Position
}

type Warehouse struct {
	grid   [][]rune
	robot  Position
	boxes  []BoxPosition
	width  int
	height int
	moves  []rune
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
				// Transform the line to double-width format
				var wideLine strings.Builder
				for _, ch := range line {
					switch ch {
					case '#':
						wideLine.WriteString("##")
					case 'O':
						wideLine.WriteString("[]")
					case '.':
						wideLine.WriteString("..")
					case '@':
						wideLine.WriteString("@.")
					}
				}
				mapLines = append(mapLines, wideLine.String())
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
		for x, ch := range line {
			warehouse.grid[y][x] = ch
			if x%2 == 0 { // Check for box start
				if x+1 < len(line) && line[x] == '[' && line[x+1] == ']' {
					warehouse.boxes = append(warehouse.boxes, BoxPosition{
						left:  Position{x, y},
						right: Position{x + 1, y},
					})
				} else if line[x] == '@' {
					warehouse.robot = Position{x, y}
				}
			}
		}
	}

	// Parse moves
	warehouse.moves = []rune(moveStr)

	return &warehouse, nil
}

func (w *Warehouse) isBox(pos Position) bool {
	for _, box := range w.boxes {
		if box.left == pos || box.right == pos {
			return true
		}
	}
	return false
}

func (w *Warehouse) moveBox(from, to Position) bool {
	// Find which box we're trying to move
	var boxIndex int
	var targetBox BoxPosition
	found := false

	for i, box := range w.boxes {
		if box.left == from || box.right == from {
			boxIndex = i
			targetBox = box
			found = true
			break
		}
	}

	if !found {
		// If we're not moving into a box, just check if the position is valid
		return w.isValidPosition(to) && !w.isWall(to)
	}

	// Calculate the direction of movement
	dir := Position{to.x - from.x, to.y - from.y}

	// Calculate new box position
	var newLeft, newRight Position
	if dir.x > 0 {
		newLeft = Position{targetBox.left.x + 1, targetBox.left.y}
		newRight = Position{targetBox.right.x + 1, targetBox.right.y}
	} else if dir.x < 0 {
		newLeft = Position{targetBox.left.x - 1, targetBox.left.y}
		newRight = Position{targetBox.right.x - 1, targetBox.right.y}
	} else if dir.y != 0 {
		newLeft = Position{targetBox.left.x, targetBox.left.y + dir.y}
		newRight = Position{targetBox.right.x, targetBox.right.y + dir.y}
	}

	// Check if new positions are valid
	if !w.isValidPosition(newLeft) || !w.isValidPosition(newRight) ||
		w.isWall(newLeft) || w.isWall(newRight) {
		return false
	}

	// First phase: Collect all boxes that need to move
	type boxMove struct {
		index    int
		box      BoxPosition
		newLeft  Position
		newRight Position
	}
	moves := []boxMove{{index: boxIndex, box: targetBox, newLeft: newLeft, newRight: newRight}}

	// Keep checking for boxes until we find no more that need to move
	checked := make(map[int]bool)
	checked[boxIndex] = true

	for i := 0; i < len(moves); i++ { // Use index loop to allow appending during iteration
		move := moves[i]

		// Check all other boxes for collisions with this move
		for j, otherBox := range w.boxes {
			if checked[j] {
				continue
			}

			// Check if new positions would overlap with any other box
			if otherBox.left == move.newLeft || otherBox.right == move.newLeft ||
				otherBox.left == move.newRight || otherBox.right == move.newRight {

				// Calculate new positions for this box
				otherNewLeft := Position{otherBox.left.x + dir.x, otherBox.left.y + dir.y}
				otherNewRight := Position{otherBox.right.x + dir.x, otherBox.right.y + dir.y}

				// Check if the move would be valid
				if !w.isValidPosition(otherNewLeft) || !w.isValidPosition(otherNewRight) ||
					w.isWall(otherNewLeft) || w.isWall(otherNewRight) {
					return false
				}

				// Add this box to the moves list
				moves = append(moves, boxMove{
					index:    j,
					box:      otherBox,
					newLeft:  otherNewLeft,
					newRight: otherNewRight,
				})
				checked[j] = true
			}
		}
	}

	// Second phase: All moves are valid, execute them in reverse order
	for i := len(moves) - 1; i >= 0; i-- {
		move := moves[i]
		w.grid[w.boxes[move.index].left.y][w.boxes[move.index].left.x] = '.'
		w.grid[w.boxes[move.index].right.y][w.boxes[move.index].right.x] = '.'
		w.boxes[move.index] = BoxPosition{left: move.newLeft, right: move.newRight}
		w.grid[move.newLeft.y][move.newLeft.x] = '['
		w.grid[move.newRight.y][move.newRight.x] = ']'
	}

	return true
}

func (w *Warehouse) isWall(pos Position) bool {
	if !w.isValidPosition(pos) {
		return true
	}
	return w.grid[pos.y][pos.x] == '#'
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
	}

	newPos := Position{w.robot.x + delta.x, w.robot.y + delta.y}
	if !w.isValidPosition(newPos) || w.isWall(newPos) {
		return false
	}

	// If moving into a box, try to push it
	if w.isBox(newPos) {
		if !w.moveBox(newPos, Position{newPos.x + delta.x, newPos.y + delta.y}) {
			return false
		}
	}

	// Move robot
	w.grid[w.robot.y][w.robot.x] = '.' // Clear the old position
	w.robot = newPos
	w.grid[newPos.y][newPos.x] = 'R' // Set the new position
	return true
}

func (w *Warehouse) calculateGPSSum() int {
	sum := 0
	for _, box := range w.boxes {
		// Find the closest edge from top (minimum y)
		topDistance := min(box.left.y, box.right.y)

		// Find the closest edge from left (minimum x)
		leftDistance := min(box.left.x, box.right.x)

		// Calculate GPS coordinate using closest edges
		gpsCoordinate := (topDistance * 100) + leftDistance
		sum += gpsCoordinate
	}
	return sum
}

func (w *Warehouse) String() string {
	var sb strings.Builder
	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			isBoxPart := false
			for _, box := range w.boxes {
				if box.left.x == x && box.left.y == y {
					sb.WriteRune('[')
					isBoxPart = true
					break
				} else if box.right.x == x && box.right.y == y {
					sb.WriteRune(']')
					isBoxPart = true
					break
				}
			}
			if !isBoxPart {
				if w.robot.x == x && w.robot.y == y {
					sb.WriteRune('@')
				} else {
					sb.WriteRune(w.grid[y][x])
				}
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (w *Warehouse) executeAllMoves() {
	// reader := bufio.NewReader(os.Stdin)
	fmt.Println("Initial state:")
	fmt.Println(w.String())

	for _, move := range w.moves {
		// fmt.Printf("\nMove %d: %c (Press Enter to continue...)\n", i, move)
		// reader.ReadString('\n')

		w.move(move)
		// fmt.Println(w.String())
	}
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

	warehouse, err := parseWarehouse(lines)
	if err != nil {
		fmt.Printf("Error parsing warehouse: %v\n", err)
		return
	}

	warehouse.executeAllMoves()
	fmt.Printf("The final GPS sum is: %d\n", warehouse.calculateGPSSum())
	fmt.Println(warehouse.String())
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
