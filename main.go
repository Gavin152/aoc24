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
	bLength := 0
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	for _, plot := range region.Plots {
		for _, direction := range directions {
			if plot.Coordinates[0] + direction[0] < 0 || plot.Coordinates[0] + direction[0] >= len(garden) {
				bLength++
				continue
			} else if plot.Coordinates[1] + direction[1] < 0 || plot.Coordinates[1] + direction[1] >= len(garden[0]) {
				bLength++
				continue
			} else if garden[plot.Coordinates[0] + direction[0]][plot.Coordinates[1] + direction[1]] != region.Identifier {
				bLength++
				continue
			}
		}
	}
	region.BoundaryLength = bLength
}

func printRegions() {
	for _, region := range regions {
		fmt.Printf("Region %s: Area: %d Boundary: %d Score: %d\n", string(region.Identifier), region.Area, region.BoundaryLength, region.Area * region.BoundaryLength)
	}
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
