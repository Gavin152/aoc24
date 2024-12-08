package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Gavin152/aoc24/internal/filereader"
)

type Line struct {
	result      int
	factors     []int
	solvable    bool
	calculation string
}

func dissectLine(line string) (Line, error) {
	// Split into result and factors parts
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return Line{}, fmt.Errorf("invalid line format: missing colon separator")
	}

	// Parse result
	result, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return Line{}, fmt.Errorf("invalid result: %v", err)
	}

	// Parse factors
	factorStrs := strings.Fields(parts[1])
	factors := make([]int, 0, len(factorStrs))

	for _, f := range factorStrs {
		factor, err := strconv.Atoi(f)
		if err != nil {
			return Line{}, fmt.Errorf("invalid factor: %v", err)
		}
		factors = append(factors, factor)
	}

	return Line{
		result:      result,
		factors:     factors,
		solvable:    false,
		calculation: "",
	}, nil
}

func solveLine(line *Line) {
	fmt.Printf("===========================\nSolving line: %v\n", line)
	// Try all possible combinations of operators between factors
	numFactors := len(line.factors)
	if numFactors < 2 && line.result == line.factors[0] {
		line.solvable = true
		return
	}

	// Generate all possible operator combinations
	numOperators := numFactors - 1
	maxCombinations := 1 << (numOperators * 2) // 3 operators need 2 bits each

	for i := 0; i < maxCombinations; i++ {
		// Build operator sequence from binary representation
		operators := make([]string, numOperators)
		for j := 0; j < numOperators; j++ {
			// Use 2 bits to represent 3 operators
			bits := (i >> (j * 2)) & 3

			switch bits {
			case 0:
				operators[j] = "+"
			case 1:
				operators[j] = "*"
			case 2, 3:
				operators[j] = "||"
			default:
				fmt.Printf("Invalid operator: %d\n", bits)
			}
		}
		// fmt.Printf("Num Operators: %d\n", numOperators)
		// fmt.Printf("Operators: %v\n", operators)

		// Build calculation string as we go
		calcStr := strconv.Itoa(line.factors[0])
		result := line.factors[0]
		for j := 0; j < numOperators; j++ {
			switch operators[j] {
			case "+":
				result += line.factors[j+1]
				calcStr += " + " + strconv.Itoa(line.factors[j+1])
			case "*":
				result *= line.factors[j+1]
				calcStr += " * " + strconv.Itoa(line.factors[j+1])
			case "||":
				// Convert current result and next factor to strings and concatenate
				resultStr := strconv.Itoa(result)
				nextStr := strconv.Itoa(line.factors[j+1])
				concatenated, _ := strconv.Atoi(resultStr + nextStr)
				result = concatenated
				calcStr += " || " + strconv.Itoa(line.factors[j+1])
			}
		}

		if result == line.result {
			line.solvable = true
			line.calculation = strconv.Itoa(result) + ":  " + calcStr
			fmt.Printf("%v\n", line.calculation)
			return
		}
	}
}

func main() {

	// filePath := "example.txt"
	filePath := "data"

	solvableLines := []Line{}

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		dissected, err := dissectLine(line)
		if err != nil {
			return err
		}
		solveLine(&dissected)
		if dissected.solvable {
			solvableLines = append(solvableLines, dissected)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	sum := 0
	for _, line := range solvableLines {
		sum += line.result
	}

	fmt.Printf("Sum of solvable lines: %d\n", sum)
	fmt.Printf("Solvable Lines %d\n", len(solvableLines))
}
