package filereader

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFileLineByLine reads a file line by line and executes a callback function on each line.
func ReadFileLineByLine(filePath string, lineProcessor func(string) error) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Process each line using the provided callback
		if err := lineProcessor(scanner.Text()); err != nil {
			return fmt.Errorf("error processing line: %w", err)
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}
