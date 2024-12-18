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
	priority int
	path     []Position
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

func printPath(path []Position) {
	// Create a copy of the grid
	display := make([][]rune, len(grid))
	for i := range grid {
		display[i] = make([]rune, len(grid[i]))
		copy(display[i], grid[i])
	}

	// Mark the path with 'O'
	for _, pos := range path {
		display[pos.y][pos.x] = 'O'
	}
	// Always mark start and end with S and E
	display[start.y][start.x] = 'S'
	display[end.y][end.x] = 'E'

	for _, row := range display {
		fmt.Println(string(row))
	}
}

// Count unique positions across all paths
func countUniqueLocations(paths [][]Position) int {
	unique := make(map[Position]bool)
	
	// Explicitly add start and end positions
	unique[start] = true
	unique[end] = true
	
	// Add all positions from all paths
	for _, path := range paths {
		for _, pos := range path {
			unique[pos] = true
		}
	}
	return len(unique)
}

func findPaths() (int, [][]Position) {
	type visitKey struct {
		pos Position
		dir rune
	}
	visited := make(map[visitKey]int)
	shortestPaths := make([][]Position, 0)
	bestScore := -1

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	
	// Include start position in initial path
	initialPath := []Position{start}
	
	heap.Push(&pq, &State{
		pos:      start,
		dir:      East,
		score:    0,
		priority: 0,
		path:     initialPath,
	})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*State)
		key := visitKey{current.pos, current.dir}

		// If we've seen this state before with a better score, skip it
		if bestScore, exists := visited[key]; exists && bestScore < current.score {
			continue
		}
		visited[key] = current.score

		// If we reached the end
		if current.pos == end {
			if bestScore == -1 || current.score < bestScore {
				// Found a better path
				// Include end position in path
				finalPath := append(current.path, end)
				bestScore = current.score
				shortestPaths = [][]Position{finalPath}
			} else if current.score == bestScore {
				// Found another path of same length
				finalPath := append(current.path, end)
				shortestPaths = append(shortestPaths, finalPath)
			}
			continue
		}

		// Try moving forward
		nextPos := move(current.pos, current.dir)
		if isValid(nextPos) {
			newPath := make([]Position, len(current.path))
			copy(newPath, current.path)
			heap.Push(&pq, &State{
				pos:      nextPos,
				dir:      current.dir,
				score:    current.score + 1,
				priority: current.score + 1,
				path:     append(newPath, nextPos),
			})
		}

		// Try turning and moving
		for _, newDir := range getTurnDirections(current.dir) {
			nextPos := move(current.pos, newDir)
			if isValid(nextPos) {
				newScore := current.score + 1000 + 1
				newPath := make([]Position, len(current.path))
				copy(newPath, current.path)
				heap.Push(&pq, &State{
					pos:      nextPos,
					dir:      newDir,
					score:    newScore,
					priority: newScore,
					path:     append(newPath, nextPos),
				})
			}
		}
	}

	return bestScore, shortestPaths
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

	parseLabyrinth(lines)
	fmt.Printf("Start position: x=%d, y=%d\n", start.x, start.y)
	fmt.Printf("End position: x=%d, y=%d\n", end.x, end.y)
	fmt.Println("Labyrinth:")
	printLabyrinth(grid)

	if bestScore, paths := findPaths(); bestScore != -1 {
		fmt.Printf("\nBest score: %d\n", bestScore)
		fmt.Printf("Number of paths with best score: %d\n", len(paths))
		fmt.Printf("Number of unique locations visited: %d\n", countUniqueLocations(paths))
		// fmt.Println("\nAll shortest paths:")
		// for i, path := range paths {
		// 	fmt.Printf("\nPath %d:\n", i+1)
		// 	printPath(path)
		// 	fmt.Println()
		// }
	} else {
		fmt.Println("\nNo path found!")
	}
}
