package util

func sliceToMatrix(lines []string) [][]rune {
	xlen := len(lines[0])
	ylen := len(lines)
	lab = make([][]rune, xlen)
	for i, _ := range lab {
		lab[i] = make([]rune, ylen)
	}

	// fmt.Printf("lab length: %d, lab element length: %d\n", len(lab), len(lab[0]))
	for i, line := range lines {
		col := []rune(line)
		for j, character := range col {
			// fmt.Printf("i: %d, j: %d\n", i, j)
			lab[j][i] = character
			if slices.Index([]rune{'^', '>', 'v', '<'}, character) > -1 {
				startPoint = []int{j, i}
				startDirection = character
			}
		}
	}
}

func printLab() {
	for i, col := range lab {
		for j, _ := range col {
			fmt.Printf("%s", string(lab[j][i]))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
