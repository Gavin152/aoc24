package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Gavin152/aoc24/internal/filereader"
)

func unpack(line string) string {
	blocks := strings.Split(line, "")
	blockString := ""
	idx := 0
	for i, block := range blocks {
		bump := true
		blockSize, _ := strconv.Atoi(block)
		if blockSize == 0 {
			idx++
			continue	
		}
		for j := 0; j < blockSize; j++ {
			if i%2 == 0 {
				blockString += strconv.Itoa(idx)
			} else {
				blockString += "."
				if bump {
					idx++
					bump = false
				}
			}
		}
	}

	return blockString
}

func findLastDigit(runes []rune, pos int) (digit rune, index int) {
	for i := len(runes) - 1; i >= 0; i-- {
		if i < pos {
			break
		}
		if runes[i] != '.' {
			// fmt.Printf("Found digit %c at index %d\n", runes[i], i)
			return runes[i], i
		}
	}
	return -1, pos
}

func defrag(line string) string {
	raw_blocks := []rune(line)
	new_blocks := []rune(line)
	for i := 0; i < len(raw_blocks); i++ {
		if raw_blocks[i] != '.' {
			continue
		} else {
			lastDigit, lastIndex := findLastDigit(raw_blocks, i)
			if lastDigit != -1 {
				new_blocks[i] = lastDigit
				raw_blocks[lastIndex] = '.'
			}
			new_blocks[lastIndex] = '.'
		}
	}
	return string(new_blocks)
}

func calculateChecksum(line string) int {
	blocks := strings.Split(line, "")
	checksum := 0
	for i, block := range blocks {
		blockNum, _ := strconv.Atoi(block)
		checksum += blockNum * i
	}
	return checksum
}

func main() {

	filePath := "example.txt"
	// filePath := "data"

	checksum := 0
	unpacked := ""
	defragged := ""

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		unpacked = unpack(line)
		defragged = defrag(unpacked)
		checksum += calculateChecksum(defragged)
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Unpacked: %s\n", unpacked)
	fmt.Printf("Defragged: %s\n", defragged)
	fmt.Printf("Checksum: %d\n", checksum)
}
