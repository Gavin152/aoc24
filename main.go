package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Gavin152/aoc24/internal/filereader"
)

var regA, regB, regC int
var program []int
var output []int

func parse(lines []string) error {
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "Register A:") {
			val, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "Register A:")))
			if err != nil {
				return fmt.Errorf("failed to parse Register A: %v", err)
			}
			regA = val
		} else if strings.HasPrefix(line, "Register B:") {
			val, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "Register B:")))
			if err != nil {
				return fmt.Errorf("failed to parse Register B: %v", err)
			}
			regB = val
		} else if strings.HasPrefix(line, "Register C:") {
			val, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "Register C:")))
			if err != nil {
				return fmt.Errorf("failed to parse Register C: %v", err)
			}
			regC = val
		} else if strings.HasPrefix(line, "Program:") {
			program = []int{} // Reset program
			numStr := strings.TrimSpace(strings.TrimPrefix(line, "Program:"))
			nums := strings.Split(numStr, ",")
			for _, num := range nums {
				val, err := strconv.Atoi(strings.TrimSpace(num))
				if err != nil {
					return fmt.Errorf("failed to parse program number: %v", err)
				}
				program = append(program, val)
			}
		}
	}
	return nil
}

// resolveOperand returns the actual value of an operand based on whether it's a literal or combo operand
func resolveOperand(operand int, isCombo bool) int {
	if !isCombo {
		return operand
	}
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return regA
	case 5:
		return regB
	case 6:
		return regC
	default: // case 7 is reserved
		return 0
	}
}

func run() {
	for i := 0; i < len(program); i += 2 {
		opcode := program[i]
		rawOperand := program[i+1]

		// determine if operand is combo based on instruction
		isCombo := opcode == 0 || opcode == 2 || opcode == 5 || opcode == 6 || opcode == 7
		operand := resolveOperand(rawOperand, isCombo)

		switch opcode {
		case 0: // adv - divide A by 2^operand, store in A
			divisor := 1 << operand
			regA = regA / divisor

		case 1: // bxl - XOR B with literal operand
			regB = regB ^ operand

		case 2: // bst - store operand mod 8 in B
			regB = operand % 8

		case 3: // jnz - jump if A is not zero
			if regA != 0 {
				i = operand - 2 // -2 because the loop will add 2
				continue
			}

		case 4: // bxc - XOR B with C (ignores operand)
			regB = regB ^ regC

		case 5: // out - output operand mod 8
			output = append(output, operand%8)

		case 6: // bdv - divide A by 2^operand, store in B
			divisor := 1 << operand
			regB = regA / divisor

		case 7: // cdv - divide A by 2^operand, store in C
			divisor := 1 << operand
			regC = regA / divisor
		}
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

	if err := parse(lines); err != nil {
		fmt.Printf("Error parsing input: %v\n", err)
		return
	}

	run()

	fmt.Printf("Register A: %d\n", regA)
	fmt.Printf("Register B: %d\n", regB)
	fmt.Printf("Register C: %d\n", regC)
	fmt.Printf("Program: %v\n", program)
	if len(output) > 0 {
		fmt.Printf("Output: %v\n", strings.Join(strings.Fields(fmt.Sprint(output)), ","))
	}
}
