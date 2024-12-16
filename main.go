package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Gavin152/aoc24/internal/filereader"
)

type Robot struct {
	PX int
	PY int
	VX int
	VY int
}

type Quadrant struct {
	MinX int
	MinY int
	MaxX int
	MaxY int
	RobotCount int
}

var robots = []Robot{}
var maxX = 11
var maxY = 7
var q1 = Quadrant{}
var q2 = Quadrant{}
var q3 = Quadrant{}
var q4 = Quadrant{}

func parseLine(line string) Robot {
	params := strings.Split(line, " ")
	pos := strings.Split(strings.Split(params[0], "=")[1], ",")
	vel := strings.Split(strings.Split(params[1], "=")[1], ",")

	px, _ := strconv.Atoi(pos[0])
	py, _ := strconv.Atoi(pos[1])
	vx, _ := strconv.Atoi(vel[0])
	vy, _ := strconv.Atoi(vel[1])

	return Robot{
		PX: px,
		PY: py,
		VX: vx,
		VY: vy,
	}
}

func move(robot *Robot) {
	newX := robot.PX + robot.VX
	newY := robot.PY + robot.VY

	if newX < 0 {
		newX = maxX + newX
	}
	if newX >= maxX {
		newX = newX - maxX
	}
	if newY < 0 {
		newY = maxY + newY
	}
	if newY >= maxY {
		newY = newY - maxY
	}

	robot.PX = newX
	robot.PY = newY
}

func printGrid() {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			found := 0
			for _, robot := range robots {
				if robot.PX == x && robot.PY == y {
					found++
				}
			}
			if found > 0 {
				fmt.Printf("%d", found)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func setQuadrants() {
	q1 = Quadrant{
		MinX: 0,
		MinY: 0,
		MaxX: maxX / 2,
		MaxY: maxY / 2,
		RobotCount: 0,
	}
	q2 = Quadrant{
		MinX: maxX / 2 + 1,
		MinY: 0,
		MaxX: maxX,
		MaxY: maxY / 2,
		RobotCount: 0,
	}
	q3 = Quadrant{
		MinX: 0,
		MinY: maxY / 2 + 1,
		MaxX: maxX / 2,
		MaxY: maxY,
		RobotCount: 0,
	}
	q4 = Quadrant{
		MinX: maxX / 2 + 1,
		MinY: maxY / 2 + 1,
		MaxX: maxX,
		MaxY: maxY,
		RobotCount: 0,
	}
}

func hasVerticalNeighbour(robot *Robot, depth int) bool {
	if depth == 0 {
		return true
	}
	for _, robot2 := range robots {
		if robot2.PX == robot.PX && robot2.PY == robot.PY+1 && hasVerticalNeighbour(&robot2, depth-1) {
			return true
		}
	}
	return false
}

func findCluster() bool {
	for _, robot := range robots {
		if hasVerticalNeighbour(&robot, 8) {
			return true
		}
	}
	return false
}

func main() {

	filePath := "data"
	// filePath := "example"

	if filePath == "example" {
		maxX = 11
		maxY = 7
	} else {
		maxX = 101
		maxY = 103
	}

	setQuadrants()

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		robot := parseLine(line)
		robots = append(robots, robot)
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// printGrid()
	for i := 0; true; i++ {
		for r := range robots {
			move(&robots[r])
		}

		if findCluster() {
			fmt.Printf("Cluster found at step %d\n", i+1)
			printGrid()
			return
		}
	}

	for _, robot := range robots {
		if robot.PX >= q1.MinX && robot.PX < q1.MaxX && robot.PY >= q1.MinY && robot.PY < q1.MaxY {
			q1.RobotCount++
		}
		if robot.PX >= q2.MinX && robot.PX < q2.MaxX && robot.PY >= q2.MinY && robot.PY < q2.MaxY {
			q2.RobotCount++
		}
		if robot.PX >= q3.MinX && robot.PX < q3.MaxX && robot.PY >= q3.MinY && robot.PY < q3.MaxY {
			q3.RobotCount++
		}
		if robot.PX >= q4.MinX && robot.PX < q4.MaxX && robot.PY >= q4.MinY && robot.PY < q4.MaxY {
			q4.RobotCount++
		}
	}

	// printGrid()
	// fmt.Printf("Q1: %d\n", q1.RobotCount)
	// fmt.Printf("Q2: %d\n", q2.RobotCount)
	// fmt.Printf("Q3: %d\n", q3.RobotCount)
	// fmt.Printf("Q4: %d\n", q4.RobotCount)
	// fmt.Printf("Total: %d\n", q1.RobotCount*q2.RobotCount*q3.RobotCount*q4.RobotCount)
}
