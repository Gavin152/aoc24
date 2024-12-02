package main

import (
	"fmt"

	"github.com/Gavin152/aoc24/internal/filereader"
)

func main() {
	// Example usage
	filePath := "example.txt"
	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		// Custom processing logic for each line
		fmt.Println(line)
		return nil // Return an error if processing fails
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
