package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/Gavin152/aoc24/internal/filereader"
	// "github.com/Gavin152/aoc24/internal/util"
)

var stones []int

type StoneBlinkRes struct {
	Value 	int
	Blinks 	int
}

var bCache = make(map[StoneBlinkRes]int)

func cblink(stone int, blinks int) int {
	if val, ok := bCache[StoneBlinkRes{Value: stone, Blinks: blinks}]; ok {
		return val
	}
	if blinks == 0 {
		return 1
	}
	if stone == 0 {
		return cblink(1, blinks-1)
	}
	if splitable := strconv.Itoa(stone); len(splitable) % 2 == 0 {
		left, right := splitStone(stone)
		created := cblink(left, blinks-1) + cblink(right, blinks-1)
		bCache[StoneBlinkRes{Value: stone, Blinks: blinks}] = created
		return created
	} else {
		created := cblink(stone * 2024, blinks-1)
		bCache[StoneBlinkRes{Value: stone, Blinks: blinks}] = created
		return created
	}
}

func blink(stones []int) []int {
	newStones := []int{}
	if len(stones) > 100 {
		for _, stone := range stones {
			newStones = append(newStones, blink([]int{stone})...)
		}
	} else {
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
	}
	return newStones
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

	totalStones := 0
	for _, stone := range stones {
		totalStones += cblink(stone, 75)
	}

	fmt.Println(bCache)
	fmt.Printf("The total number of stones is %d\n", totalStones)
}
