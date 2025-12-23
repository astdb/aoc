package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./d4 <inputfile>")
	}

	paperLocations := processInput(os.Args[1])
	fmt.Printf("Paper locations map: %v\n", paperLocations)

	accessibleRolls := countAccessibleRolls(paperLocations)

	fmt.Println(accessibleRolls)
}

func countAccessibleRolls(rollLocations [][]rune) int {
	accessibleRolls := 0
	for row := 0; row < len(rollLocations); row++ {
		for col := 0; col < len(rollLocations[row]); col++ {
			if rollAccessible(row, col, rollLocations) {
				accessibleRolls++
			}
		}
	}

	return accessibleRolls
}

func rollAccessible(row, col int, locs [][]rune) bool {
	totalAdjRolls := 0

	// check the top three
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

		// row[col-1], row[col], row col[]
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

		// row[col-1], row[col], row col[]
	}

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