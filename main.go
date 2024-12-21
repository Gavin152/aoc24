package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Gavin152/aoc24/internal/filereader"
)

var memSpace [][]rune

func initGrid(size int) {
	grid := make([][]rune, size)
	for i := range grid {
		grid[i] = make([]rune, size)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	memSpace = grid
}

func fillMemSpace(bytes []string, start int, end int) []int {
	if start >= len(bytes) {
		return []int{-1, -1}
	}
	if end > len(bytes) {
		end = len(bytes)
	}
	for i := start; i < end; i++ {
		split := strings.Split(bytes[i], ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		if memSpace[x][y] == 'O' {
			return []int{x, y}
		}
		memSpace[x][y] = '#'
	}
	return []int{-1, -1}
}

type Point struct {
	x, y int
}

func (p Point) isValid(size int) bool {
	return p.x >= 0 && p.x < size && p.y >= 0 && p.y < size
}

func printGrid() {
	for i := range memSpace {
		for j := range memSpace[i] {
			fmt.Print(string(memSpace[j][i]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func findShortestPath() int {
	size := len(memSpace)
	if size == 0 {
		return -1 // Invalid grid
	}

	dist := make([][]int, size)
	visited := make([][]bool, size)
	
	// Initialize matrices
	for i := range dist {
		dist[i] = make([]int, size)
		visited[i] = make([]bool, size)
		for j := range dist[i] {
			dist[i][j] = int(^uint(0) >> 1) // Max int
		}
	}

	type queueItem struct {
		p    Point
		dist int
	}
	pq := []queueItem{{Point{0, 0}, 0}}
	dist[0][0] = 0

	moves := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for len(pq) > 0 {
		current := pq[0]
		pq = pq[1:]

		if !current.p.isValid(size) || visited[current.p.x][current.p.y] {
			continue
		}

		visited[current.p.x][current.p.y] = true

		if current.p.x == size-1 && current.p.y == size-1 {
			return current.dist
		}

		for _, move := range moves {
			newX, newY := current.p.x+move[0], current.p.y+move[1]
			next := Point{newX, newY}

			if !next.isValid(size) || visited[newX][newY] || memSpace[newX][newY] == '#' {
				continue
			}

			newDist := current.dist + 1
			if newDist < dist[newX][newY] {
				dist[newX][newY] = newDist

				inserted := false
				for i, p := range pq {
					if newDist < p.dist {
						pq = append(pq[:i], append([]queueItem{{next, newDist}}, pq[i:]...)...)
						inserted = true
						break
					}
				}
				if !inserted {
					pq = append(pq, queueItem{next, newDist})
				}
			}
		}
	}

	return -1
}

func main() {
	// filePath := "example"
	filePath := "data"

	start := 0

	if filePath == "example" {
		initGrid(7)
		start = 12
	} else {
		initGrid(71)
		start = 1024
	}

	var lines []string
	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		lines = append(lines, line)
		return nil
	})
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fillMemSpace(lines, 0, start)

	lastPos := 0
	for i := start; i < len(lines); i++ {
		fillMemSpace(lines, i, i+1)
		dist := findShortestPath()
		if dist == -1 {
			lastPos = i
			break
		}
	}

	fmt.Printf("Shortest Path of %d would be blocked at %v\n", lastPos, lines[lastPos])
	printGrid()
}
