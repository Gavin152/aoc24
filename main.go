package main

import (
	"fmt"

	"github.com/Gavin152/aoc24/internal/filereader"
	"github.com/Gavin152/aoc24/internal/util"
)

var (
	trailHeads [][]int
	grid [][]int
)	

func getTrailHeads(grid [][]int) [][]int {
	trailHeads := [][]int{}
	for i, row := range grid {
		for j, col := range row {
			if col == 0 {
				trailHeads = append(trailHeads, []int{i, j})
			}
		}
	}
	return trailHeads
}

func hike(pos []int) [][]int {
	// fmt.Printf("Hiking from %v\n", pos)
	summits := [][]int{}
	for _, nextStep := range getNextStep(pos) {
		if grid[nextStep[0]][nextStep[1]] == 9 {
			summits = append(summits, nextStep)
		} else {
			summits = append(summits, hike(nextStep)...)
		}
	}
	return summits
}

func getNextStep(pos []int) [][]int {
	nextSteps := [][]int{}
	fmt.Printf("Getting next steps for %v\n", pos)
	possibleSteps := [][]int{{pos[0]-1, pos[1]}, {pos[0]+1, pos[1]}, {pos[0], pos[1]-1}, {pos[0], pos[1]+1}}
	
	filteredSteps := [][]int{}
	for _, step := range possibleSteps {
		if isInGrid(step) {
			filteredSteps = append(filteredSteps, step)
		}
	}
	possibleSteps = filteredSteps
	fmt.Printf("Possible Steps are: %v\n", possibleSteps)

	for _, dir := range possibleSteps {
		if isInGrid(dir) && grid[dir[0]][dir[1]] == grid[pos[0]][pos[1]] + 1 {
			fmt.Printf("Hiking from %v to %v\n", pos, dir)
			nextSteps = append(nextSteps, dir)
		}
	}
	return nextSteps
}

func isInGrid(pos []int) bool {
	return pos[0] >= 0 && pos[0] < len(grid) && pos[1] >= 0 && pos[1] < len(grid[0])
}

func main() {

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

	rawGrid := util.SliceToGrid(lines)
	grid, _ = util.ParseGridToInt(rawGrid)
	trailHeads = getTrailHeads(grid)

	fullScore := 0
	trails := make(map[string	]int)
	for _, trailHead := range trailHeads {
		trailSummits := hike(trailHead)
		trailSummitsMap := make(map[string]bool)	
		for _, summit := range trailSummits {
			trailSummitsMap[fmt.Sprintf("%d|%d", summit[0], summit[1])] = true
		}
		trails[fmt.Sprintf("%d|%d", trailHead[0], trailHead[1])] = len(trailSummitsMap)
		fullScore += len(trailSummits)
		fmt.Printf("Trail Head: %v, Trail Summits: %v\n", trailHead, trailSummitsMap)
	}

	fmt.Printf("Trail Heads: %v\n", trailHeads)
	
	fmt.Printf("Trails: %v\n", trails)
	fmt.Printf("Full Score: %d\n", fullScore)
}
