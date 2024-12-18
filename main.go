package main

import (
	"container/heap"
	"fmt"

	"github.com/Gavin152/aoc24/internal/filereader"
)

type Position struct {
	x, y int
}

type State struct {
	pos      Position
	dir      rune
	score    int
	priority int // lower is better
}

// Priority queue implementation
type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*State)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

var (
	grid  [][]rune
	start Position
	end   Position
)

const (
	North = 'N'
	South = 'S'
	East  = 'E'
	West  = 'W'
)

// Returns the new position after moving one step in the given direction
func move(pos Position, dir rune) Position {
	switch dir {
	case North:
		return Position{pos.x, pos.y - 1}
	case South:
		return Position{pos.x, pos.y + 1}
	case East:
		return Position{pos.x + 1, pos.y}
	case West:
		return Position{pos.x - 1, pos.y}
	}
	return pos
}

// Returns the directions to try (left turn and right turn from current direction)
func getTurnDirections(dir rune) []rune {
	switch dir {
	case North:
		return []rune{West, East}
	case South:
		return []rune{East, West}
	case East:
		return []rune{North, South}
	case West:
		return []rune{South, North}
	}
	return nil
}

func parseLabyrinth(lines []string) {
	grid = make([][]rune, len(lines))
	for y, line := range lines {
		grid[y] = []rune(line)
		for x, ch := range grid[y] {
			switch ch {
			case 'S':
				start = Position{x, y}
			case 'E':
				end = Position{x, y}
			}
		}
	}
}

func printLabyrinth(grid [][]rune) {
	for y := range grid {
		fmt.Println(string(grid[y]))
	}
}

func isValid(pos Position) bool {
	return pos.x >= 0 && pos.y >= 0 && pos.x < len(grid[0]) && pos.y < len(grid) && grid[pos.y][pos.x] != '#'
}

func findPath() int {
	// Keep track of visited states (position + direction)
	type visitKey struct {
		pos Position
		dir rune
	}
	visited := make(map[visitKey]int)

	// Initialize priority queue with starting state
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &State{
		pos:      start,
		dir:      East,
		score:    0,
		priority: 0,
	})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*State)
		key := visitKey{current.pos, current.dir}

		// If we've seen this state before with a better score, skip it
		if bestScore, exists := visited[key]; exists && bestScore <= current.score {
			continue
		}
		visited[key] = current.score

		// If we reached the end, return the score
		if current.pos == end {
			return current.score
		}

		// Try moving forward
		nextPos := move(current.pos, current.dir)
		if isValid(nextPos) {
			heap.Push(&pq, &State{
				pos:      nextPos,
				dir:      current.dir,
				score:    current.score + 1,
				priority: current.score + 1,
			})
		}

		// Try turning and moving
		for _, newDir := range getTurnDirections(current.dir) {
			nextPos := move(current.pos, newDir)
			if isValid(nextPos) {
				newScore := current.score + 1000 + 1 // 1000 for turn, 1 for move
				heap.Push(&pq, &State{
					pos:      nextPos,
					dir:      newDir,
					score:    newScore,
					priority: newScore,
				})
			}
		}
	}

	return -1 // No path found
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

	parseLabyrinth(lines)
	fmt.Printf("Start position: x=%d, y=%d\n", start.x, start.y)
	fmt.Printf("End position: x=%d, y=%d\n", end.x, end.y)
	fmt.Println("Labyrinth:")
	printLabyrinth(grid)

	if result := findPath(); result != -1 {
		fmt.Printf("\nBest path found with score: %d\n", result)
	} else {
		fmt.Println("\nNo path found!")
	}
}
