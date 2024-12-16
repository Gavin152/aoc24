package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Gavin152/aoc24/internal/filereader"
)

type Button struct {
	X int
	Y int
	Cost int
	Count int
}

type clawMachine struct {
	A Button
	B Button
	X int
	Y int
	Cost int
	Winable bool
}

func parseBlock(lines []string) clawMachine {
	machine := clawMachine{}
	for _, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), " ")
		if strings.HasPrefix(line, "Button") {
			xParts := strings.Split(parts[2], "+")
			yParts := strings.Split(parts[3], "+")
			x, _ := strconv.Atoi(strings.TrimSuffix(xParts[1], ","))
			y, _ := strconv.Atoi(strings.TrimSpace(yParts[1]))
			if strings.HasPrefix(line, "Button A") {
				machine.A = Button{X: x, Y: y, Cost: 3, Count: 0}
			} else if strings.HasPrefix(line, "Button B") {
				machine.B = Button{X: x, Y: y, Cost: 1, Count: 0}
			}
		} else if strings.HasPrefix(line, "Prize") {
			xParts := strings.Split(parts[1], "=")
			yParts := strings.Split(parts[2], "=")
			x, _ := strconv.Atoi(strings.TrimSuffix(xParts[1], ","))
			y, _ := strconv.Atoi(strings.TrimSpace(yParts[1]))
			machine.X = x
			machine.Y = y
		}
	}
	machine.Winable = false
	machine.Cost = -1
	return machine
}

func solve(machine *clawMachine) {
	bTimes := (machine.Y * machine.A.X - machine.X * machine.A.Y) / (machine.A.X * machine.B.Y - machine.B.X * machine.A.Y)
	aTimes := (machine.X - bTimes * machine.B.X) / machine.A.X

	machine.A.Count = aTimes
	machine.B.Count = bTimes
	if verify(*machine) {
		machine.Winable = true
		machine.Cost = machine.A.Cost * machine.A.Count + machine.B.Cost * machine.B.Count
	} else {
		machine.Cost = -1
		machine.Winable = false
		machine.A.Count = -1
		machine.B.Count = -1
	}
}

func verify(machine clawMachine) bool {
	if machine.A.Count > 100 || machine.B.Count > 100 {
		return false
	}
	vx := machine.A.X * machine.A.Count + machine.B.X * machine.B.Count
	vy := machine.A.Y * machine.A.Count + machine.B.Y * machine.B.Count
	return vx == machine.X && vy == machine.Y
}

func main() {

	filePath := "data"
	//filePath := "example"

	arcade := []clawMachine{}
	lines := []string{}

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		if line == "" {
			clawMachine := parseBlock(lines)
			arcade = append(arcade, clawMachine)
			lines = []string{}
		} else {
			lines = append(lines, line)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	totalCost := 0

	clawMachine := parseBlock(lines)
	arcade = append(arcade, clawMachine)

	fmt.Println(len(arcade))
	for _, machine := range arcade {
		solve(&machine)
		if machine.Winable {
			totalCost += machine.Cost
		}
	}

	fmt.Println(totalCost)
}
