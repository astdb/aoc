package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Part 2
// - Start with the initial 

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./d4 <inputfile>")
	}

	paperLocations := processInput(os.Args[1])
	
	// DEBUG
	// fmt.Printf("Paper locations map: %v\n", paperLocations)

	// D4 - part 1
	// accessibleRolls := countAccessibleRolls(paperLocations)
	// fmt.Println(accessibleRolls)

	totalRollsRemoved := 0
	rollsRemoved, paperLocs := countAccessibleRolls(paperLocations)
	totalRollsRemoved = rollsRemoved

	for rollsRemoved > 0 {
		rollsRemoved, paperLocs = countAccessibleRolls(paperLocs)
		totalRollsRemoved = totalRollsRemoved + rollsRemoved
	}
	
	// fmt.Println(accessibleRolls)		// D4 - part 1
	fmt.Println(totalRollsRemoved)
}

func countAccessibleRolls(rollLocations [][]rune) (int, [][]rune) {
	accessibleRolls := 0
	rollLocsCopy := deepCopyRollLocs(rollLocations)
	
	for row := 0; row < len(rollLocations); row++ {
		for col := 0; col < len(rollLocations[row]); col++ {
			if rollLocations[row][col] == '@' && rollAccessible(row, col, rollLocations) {
				accessibleRolls++

				// 'remove' this accessible roll
				rollLocsCopy[row][col] = '.'
			}
		}
	}

	return accessibleRolls, rollLocsCopy
}

func deepCopyRollLocs(src [][]rune) [][]rune {
	dst := make([][]rune, len(src))

	for i, innerSrc := range src {
		innerDst := make([]rune, len(innerSrc))
		copy(innerDst, innerSrc)
		dst[i] = innerDst
	}

	return dst
}

func rollAccessible(row, col int, locs [][]rune) bool {
	totalAdjRolls := 0

	// check above three
	if row-1 >= 0 && row-1 < len(locs) {
		prevRow := locs[row-1]
		if col-1 >= 0 && col-1 < len(prevRow) {
			if prevRow[col-1] == '@' {
				totalAdjRolls++
			}
		}

		if col >= 0 && col < len(prevRow) {
			if prevRow[col] == '@' {
				totalAdjRolls++
			}
		}

		if col+1 >= 0 && col+1 < len(prevRow) {
			if prevRow[col+1] == '@' {
				totalAdjRolls++
			}
		}
	}

	// check left/right
	if col-1 >= 0 && col-1 < len(locs[row]) {
		if locs[row][col-1] == '@' {
			totalAdjRolls++
		}
	}

	if col+1 >= 0 && col+1 < len(locs[row]) {
		if locs[row][col+1] == '@' {
			totalAdjRolls++
		}
	}

	// check below three
	if row+1 >= 0 && row+1 < len(locs) {
		nxtRow := locs[row+1]
		if col-1 >= 0 && col-1 < len(nxtRow) {
			if nxtRow[col-1] == '@' {
				totalAdjRolls++
			}
		}

		if col >= 0 && col < len(nxtRow) {
			if nxtRow[col] == '@' {
				totalAdjRolls++
			}
		}

		if col+1 >= 0 && col+1 < len(nxtRow) {
			if nxtRow[col+1] == '@' {
				totalAdjRolls++
			}
		}
	}

	// DEBUG
	// fmt.Printf("Row: %d / Col: %d / Adj Roll Count: %d\n", row, col, totalAdjRolls)

	return totalAdjRolls < 4
}

func processInput(filename string) [][]rune {
	rows := [][]rune{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}

	inputScanner := bufio.NewScanner(file)
	for inputScanner.Scan() {
		row := []rune{}
		line := inputScanner.Text()
		for _, ch := range line {
			row = append(row, ch)
		}

		rows = append(rows, row)
	}

	return rows
}
