package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Gavin152/aoc24/internal/filereader"
)

func main() {

	safeCount := 0

	// filePath := "example.txt"
	filePath := "data_a"
	// filePath := "data_b"

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
		multies := re.FindAllString(line, -1)

		for _, mul := range multies {
			s, _ := strings.CutSuffix(mul, ")")
			s, _ = strings.CutPrefix(s, "mul(")
			factors := strings.Split(s, ",")
			factor1, _ := strconv.Atoi(factors[0])
			factor2, _ := strconv.Atoi(factors[1])
			product := factor1 * factor2
			fmt.Printf("%d x %d = %d\n", factor1, factor2, product)
			safeCount += product
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Found %d safe reports", safeCount)
}
