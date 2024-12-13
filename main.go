package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/Gavin152/aoc24/internal/filereader"
	// "github.com/Gavin152/aoc24/internal/util"
)

var stones []int

func blink() {
	newStones := []int{}
	for _, stone := range stones {
		splitable := splitable(stone)
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if splitable {
			left, right := splitStone(stone)
			newStones = append(newStones, left)
			newStones = append(newStones, right)
		} else {
			newStones = append(newStones, stone * 2024)
		}
	}
	stones = newStones
}

func splitable(stone int) bool {
	str := strconv.Itoa(stone)
	splitable := len(str) % 2 == 0
	return splitable
}

func splitStone(stone int) (int, int) {
	str := strconv.Itoa(stone)
	split, err := strconv.Atoi(str[:len(str)/2])
	split2, err2 := strconv.Atoi(str[len(str)/2:])
	if err != nil || err2 != nil {
		panic(err)
	}
	return split, split2
}

func main() {

	// filePath := "example.txt"
	filePath := "data"

	// lines := []string{}

	err := filereader.ReadFileLineByLine(filePath, func(line string) error {
		// lines = append(lines, line)
		strLine := strings.Split(line, " ")
		for _, stone := range strLine {
			stoneInt, err := strconv.Atoi(stone)
			if err != nil {
				panic(err)
			}
			stones = append(stones, stoneInt)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println(stones)

	for i := 0; i < 25; i++ {
		blink()
		// fmt.Println(stones)
		// fmt.Println("======================================")
	}

	fmt.Printf("The total number of stones is %d\n", len(stones))
}
