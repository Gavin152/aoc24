package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/Gavin152/aoc24/internal/filereader"
)

func unpack(line string) []int {
	blocks := strings.Split(line, "")
	blockString := []int{}
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
				blockString = append(blockString, idx)
			} else {
				blockString = append(blockString, -1)
				if bump {
					idx++
					bump = false
				}
			}
		}
	}

	return blockString
}

func findLastDigit(runes []int, pos int) (digit int, index int) {
	for i := len(runes) - 1; i >= 0; i-- {
		if i < pos {
			break
		}
		if runes[i] != -1 {
			// fmt.Printf("Found digit %c at index %d\n", runes[i], i)
			return runes[i], i
		}
	}
	return -1, pos
}

func defrag(line []int) []int {
	raw_blocks := line
	new_blocks := line
	for i := 0; i < len(raw_blocks); i++ {
		if raw_blocks[i] != -1 {
			continue
		} else {
			lastDigit, lastIndex := findLastDigit(raw_blocks, i)
			if lastDigit != -1 {
				new_blocks[i] = lastDigit
				raw_blocks[lastIndex] = -1
			}
			new_blocks[lastIndex] = -1
		}
	}
	cleaned := slices.DeleteFunc(new_blocks, func(r int) bool {
		return r == -1
	})
	return cleaned
}

func getFreeBlocks(line []int) [][]int {
	free_blocks := [][]int{}
	count_free := 0
	index_free := -1
	for i, block := range line {

		// fmt.Printf("Block: %d at index %d\n", block, i)
		if block == -1 {
			count_free++
			if index_free == -1 {
				index_free = i
			}
		} else if count_free > 0 {
			free_blocks = append(free_blocks, []int{index_free, count_free})
			count_free = 0
			index_free = -1
		}
	}
	// fmt.Printf("Free blocks: %v\n", free_blocks)
	return free_blocks
}

func reshuffle(line []int) []int {
	free_blocks := getFreeBlocks(line)

	file_blocks := [][]int{}
	tmp_block := []int{}

	block_id := -1
	// block_length := 0
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= 0 {
			if block_id == -1 {
				block_id = line[i]
			}
			if line[i] == block_id {
				tmp_block = append(tmp_block, line[i])
			} else {
				file_blocks = append(file_blocks, tmp_block)
				if line[i] != -1 {
					block_id = line[i]
					tmp_block = []int{block_id}
				} else {
					tmp_block = []int{}
				}
			}
		}
	}
	// fmt.Printf("File blocks: %v\n", file_blocks)

	for _, fb := range file_blocks {
		// fmt.Printf("File block: %v\n", fb)
		dest_index := slices.IndexFunc(free_blocks, func(free_block []int) bool {
			return free_block[1] >= len(fb)
		})
		if dest_index != -1 {
			dest_block := free_blocks[dest_index]
			if dest_block[0] >= slices.Index(line, fb[0]) {
				continue
			}
			for slices.Index(line, fb[0]) != -1 {
				idx := slices.Index(line, fb[0])
				line[idx] = -1
			}

			
			for j := 0; j < len(fb); j++ {
				// fmt.Printf("Moving %d to index %d\n", fb[0], dest_block[0]+j)
				line[dest_block[0]+j] = fb[0]
			}
			free_blocks = getFreeBlocks(line)
		} else {
			// fmt.Printf("No free space found for block %v\n", fb)
		}
		
	}

	return line
}

func calculateChecksum(line []int) int {
	checksum := 0
	for i, block := range line {
		if block == -1 {
			continue
		}
		checksum += block * i
	}
	return checksum
}

func main() {

	// filePath := "example.txt"
	filePath := "data"

	checksum := 0
	unpacked := []int{}
	// defragged := []int{}

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		unpacked = unpack(line)
		// defragged = defrag(unpacked)
		// checksum += calculateChecksum(defragged)

		// fmt.Printf("Unpacked: %v\n", unpacked)
		reshuffle(unpacked)
		// fmt.Printf("Reshuffled: %v\n", unpacked)
		checksum += calculateChecksum(unpacked)

		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// fmt.Printf("Defragged: %v\n", defragged)
	fmt.Printf("Checksum: %d\n", checksum)
}
