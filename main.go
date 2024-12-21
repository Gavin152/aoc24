package main

import (
	"container/list"
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

// func printRegs() {
// 	fmt.Printf("A: %d, B: %d, C: %d\n", regA, regB, regC)
// }

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

func matchesProgram(output []int, expected []int) bool {
	if len(output) != len(expected) {
		return false
	}
	for i := range output {
		if output[i] != expected[i] {
			return false
		}
	}
	return true
}

func solve() bool {
	type QueueItem struct {
		a uint64
		n int
	}

	queue := list.New()
	queue.PushBack(QueueItem{a: 0, n: 1})

	for queue.Len() > 0 {
		item := queue.Remove(queue.Front()).(QueueItem)
		a, n := item.a, item.n

		if n > len(program) { // Base case
			return true
		}

		for i := uint64(0); i < 8; i++ {
			regA := (a << 3) | i
			run()
			fmt.Printf("Running: %d, %d\n", regA, n)
			fmt.Printf("Queued: %d, %d\n", regA, n)
			target := program[len(program)-n:]
			fmt.Printf("Target: %v\n", strings.Join(strings.Fields(fmt.Sprint(target)), ","))

			// Save correct partial solutions
			if matchesProgram(output, target) {
				queue.PushBack(QueueItem{a: regA, n: n + 1})
			}
		}
	}

	return false
}

func main() {
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

	// Try the value
	if solve() {
		fmt.Println("Current value of A produces matching output!")
	} else {
		fmt.Println("Current value of A does not produce matching output.")
	}

	fmt.Printf("Register A: %d\n", regA)
	fmt.Printf("Register B: %d\n", regB)
	fmt.Printf("Register C: %d\n", regC)
	fmt.Printf("Program: %v\n", program)
	if len(output) > 0 {
		fmt.Printf("Output:  %v\n", strings.Join(strings.Fields(fmt.Sprint(output)), ","))
	}
}
