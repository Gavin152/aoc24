package main

import (
	"fmt"
	"slices"

	"github.com/Gavin152/aoc24/internal/filereader"
	"github.com/Gavin152/aoc24/internal/util"
)

type Plot struct {
	Identifier 	rune
	Coordinates []int
	BoundaryLength int
}

type Region struct {
	Identifier rune
	Plots []Plot
	BoundaryLength int
	Area int
}

var garden = [][]rune{}
var regions = []Region{}
var plotsHandled = [][]int{}

func drawMap(startingPoint []int) {
	for i, col := range garden {
		for j := range col {
			setToRegion([]int{i, j})
		}
		fmt.Println()
	}	
}

func plotHandled(coordinates []int) bool {
	return slices.ContainsFunc(plotsHandled, func(p []int) bool {
		return p[0] == coordinates[0] && p[1] == coordinates[1]
	})
}

func setToRegion(coordinates []int) {
	if plotHandled(coordinates) {
		fmt.Printf("Plot %v already handled\n", coordinates)
		return
	}
	plot := Plot{
		Identifier: rune(garden[coordinates[0]][coordinates[1]]),
		Coordinates: coordinates,
		BoundaryLength: -1,
	}
	if !plotHandled(plot.Coordinates) {
		fmt.Printf("Creating region for %c\n", plot.Identifier)
		region := Region{
			Identifier: plot.Identifier,
			Plots: []Plot{},
			BoundaryLength: -1,
			Area: -1,
		}
		regions = append(regions, region)
		growRegion(&regions[len(regions)-1], plot)
	}
}

func growRegion(region *Region, plot Plot) {
	addPlotToRegion(region, plot)
	neighbors := findCommonNeighbors(plot)
	for _, neighbor := range neighbors {
		growRegion(region, neighbor)
	}	
}

func addPlotToRegion(region *Region, plot Plot) {
	if !plotHandled(plot.Coordinates) {
		region.Plots = append(region.Plots, plot)
		plotsHandled = append(plotsHandled, plot.Coordinates)
		fmt.Printf("Added plot %c:%v to region %s\n", plot.Identifier, plot.Coordinates, string(region.Identifier))
	}
}

func findCommonNeighbors(plot Plot) []Plot {
	fmt.Printf("Finding neighbors for %c:%v\n", plot.Identifier, plot.Coordinates)
	neighbors := []Plot{}
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	for _, direction := range directions {
		if plot.Coordinates[0] + direction[0] < 0 || plot.Coordinates[0] + direction[0] >= len(garden) || plot.Coordinates[1] + direction[1] < 0 || plot.Coordinates[1] + direction[1] >= len(garden[0]) {
			continue
		}
		if plotHandled([]int{plot.Coordinates[0] + direction[0], plot.Coordinates[1] + direction[1]}) {
			continue
		}
		if plot.Identifier == rune(garden[plot.Coordinates[0] + direction[0]][plot.Coordinates[1] + direction[1]]) {
			neighbor := Plot{
				Identifier: rune(garden[plot.Coordinates[0] + direction[0]][plot.Coordinates[1] + direction[1]]),
				Coordinates: []int{plot.Coordinates[0] + direction[0], plot.Coordinates[1] + direction[1]},
				BoundaryLength: -1,
			}
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func setArea(region *Region) {
	region.Area = len(region.Plots)
}

func setBoundaryLength(region *Region) {
	sides := 0

	for _, plot := range region.Plots {
		rightNeigh := '.'
		leftNeigh := '.'
		upNeigh := '.'
		downNeigh := '.'
		lc := []int{plot.Coordinates[0], plot.Coordinates[1] - 1}
		rc := []int{plot.Coordinates[0], plot.Coordinates[1] + 1}
		uc := []int{plot.Coordinates[0] - 1, plot.Coordinates[1]}
		dc := []int{plot.Coordinates[0] + 1, plot.Coordinates[1]}
		if lc[0] < 0 || lc[1] < 0 || lc[0] >= len(garden) || lc[1] >= len(garden[0]) {
			// fmt.Printf("Left neighbor is out of bounds\n")
			leftNeigh = '.'
		} else {
			leftNeigh = garden[lc[0]][lc[1]]
		}

		if rc[0] < 0 || rc[1] < 0 || rc[0] >= len(garden) || rc[1] >= len(garden[0]) {
			//fmt.Printf("Right neighbor is out of bounds\n")
			rightNeigh = '.'
		} else {
			rightNeigh = garden[rc[0]][rc[1]]
		}

		if uc[0] < 0 || uc[1] < 0 || uc[0] >= len(garden) || uc[1] >= len(garden[0]) {
			//fmt.Printf("Up neighbor is out of bounds\n")
			upNeigh = '.'
		} else {
			upNeigh = garden[uc[0]][uc[1]]
		}

		if dc[0] < 0 || dc[1] < 0 || dc[0] >= len(garden) || dc[1] >= len(garden[0]) {
			// fmt.Printf("Down neighbor is out of bounds\n")
			downNeigh = '.'
		} else {
			downNeigh = garden[dc[0]][dc[1]]
		}

		if rightNeigh != region.Identifier && upNeigh != region.Identifier {
			sides++
		}
		if rightNeigh != region.Identifier && downNeigh != region.Identifier {
			sides++
		}
		if leftNeigh != region.Identifier && upNeigh != region.Identifier {
			sides++
		}
		if leftNeigh != region.Identifier && downNeigh != region.Identifier {
			sides++
		}
		sides += checkInnerCorner(plot)
	}
	region.BoundaryLength = sides
}

func getNeighbors(plot Plot) [][]int {
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	neighbors := [][]int{}
	for _, direction := range directions {
		x := plot.Coordinates[0] + direction[0]
		y := plot.Coordinates[1] + direction[1]
		if isInBounds([]int{x, y}) {
			neighbors = append(neighbors, []int{x, y})
			fmt.Printf("Neighbor of %c:%v: %c:%v\n", plot.Identifier, plot.Coordinates, garden[x][y], []int{x, y})
		}
	}
	return neighbors
}

func checkNeighbor(plot Plot, direction string) bool {
	var rightNeigh, leftNeigh, upNeigh, downNeigh rune

	if plot.Coordinates[1]+1 >= len(garden[0]) {
		rightNeigh = '.'
	} else {
		rightNeigh = garden[plot.Coordinates[0]][plot.Coordinates[1]+1]
	}
	
	if plot.Coordinates[1]-1 < 0 {
		leftNeigh = '.'
	} else {
		leftNeigh = garden[plot.Coordinates[0]][plot.Coordinates[1]-1]
	}
	
	if plot.Coordinates[0]-1 < 0 {
		upNeigh = '.'
	} else {
		upNeigh = garden[plot.Coordinates[0]-1][plot.Coordinates[1]]
	}
	
	if plot.Coordinates[0]+1 >= len(garden) {
		downNeigh = '.'
	} else {
		downNeigh = garden[plot.Coordinates[0]+1][plot.Coordinates[1]]
	}

	if rightNeigh == plot.Identifier && upNeigh == plot.Identifier && direction == "topright" {
		return true
	}
	if rightNeigh == plot.Identifier && downNeigh == plot.Identifier && direction == "bottomright" {
		return true
	}
	if leftNeigh == plot.Identifier && upNeigh == plot.Identifier && direction == "topleft" {
		return true
	}
	if leftNeigh == plot.Identifier && downNeigh == plot.Identifier && direction == "bottomleft" {
		return true
	}
	return false
}

func checkInnerCorner(plot Plot) int {
	corners := 0
	directions := map[string][]int{
		"bottomright": {1, 1},
		"bottomleft": {1, -1},
		"topright": {-1, 1},
		"topleft": {-1, -1},
	}
	for name, direction := range directions {
		x := plot.Coordinates[0] + direction[0]
		y := plot.Coordinates[1] + direction[1]
		// fmt.Printf("Checking inner corner at %d|%d\n", x, y)
		if isInBounds([]int{x, y}) && garden[x][y] != plot.Identifier && checkNeighbor(plot, name) {
			fmt.Printf("Found inner corner at of %c (%d|%d) %d|%d\n", plot.Identifier, plot.Coordinates[0]+1, plot.Coordinates[1]+1, x+1, y+1)
			corners++
		}
	}
	return corners
}

func isInBounds(coordinates []int) bool {
	if coordinates[0] < 0 || coordinates[0] >= len(garden) || coordinates[1] < 0 || coordinates[1] >= len(garden[0]) {
		return false
	}
	return true
}

func printRegions() {
	for _, region := range regions {
		fmt.Printf("Region %s: Area: %d Boundary: %d Score: %d\n", string(region.Identifier), region.Area, region.BoundaryLength, region.Area * region.BoundaryLength)
	}
}

func main() {

	filePath := "data"
	// filePath := "example"

	lines := []string{}

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		lines = append(lines, line)

		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	garden = util.SliceToGrid(lines)
	drawMap([]int{0, 0})

	for i:=0; i<len(regions); i++ {
		setArea(&regions[i])
		setBoundaryLength(&regions[i])
	}

	fmt.Printf("The garden has %v regions\n", len(regions))
	printRegions()

	fullScore := 0
	for _, region := range regions {
		fullScore += region.Area * region.BoundaryLength
	}
	fmt.Printf("Full score: %d\n", fullScore)
}
