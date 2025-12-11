package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	leftList, rightList, err := processInputFile(os.Getenv("D2INPUT"))
	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	simScore := 0
	rightScanStart := 0
	rightAppearances := 0
	for i := 0; i < len(leftList); i++ {
		leftNum := leftList[i]

		if !(i > 0 && leftNum == leftList[i-1]) {
			rightAppearances = 0
			for j := rightScanStart; j < len(rightList); j++ {
				if rightList[j] == leftNum {
					rightAppearances++
				}
				if rightList[j] > leftNum {
					rightScanStart = j
					break
				}
			}
		}

		// fmt.Printf("leftNum: %d, rightAppearances: %d\n", leftNum, rightAppearances)
		simScore += leftNum * rightAppearances
	}

	fmt.Println(simScore)
}

func processInputFile(inputFile string) ([]int, []int, error) {
	// fmt.Printf("inputFile: %s\n", inputFile)
	left := []int{}
	right := []int{}

	file, err := os.Open(inputFile)
	if err != nil {
		return left, right, err
	}

	inputScanner := bufio.NewScanner(file)
	for inputScanner.Scan() {
		line := inputScanner.Text()
		// fmt.Printf("line: %s\n", line)
		lineComps := strings.Split(line, "   ")
		// fmt.Printf("lineComps: %v\n", lineComps)
		if len(lineComps) == 2 {
			// fmt.Println("len(lineComps) == 2")
			leftStr := strings.TrimSpace(lineComps[0])
			rightStr := strings.TrimSpace(lineComps[1])

			// fmt.Printf("leftStr: %s, rightStr: %s\n", leftStr, rightStr)

			leftInt, leftIntErr := strconv.Atoi(leftStr)
			rightInt, rightIntErr := strconv.Atoi(rightStr)
			if leftIntErr == nil && rightIntErr == nil {
				// fmt.Printf("leftInt: %d, rightInt: %d\n", leftInt, rightInt)
				left = append(left, leftInt)
				right = append(right, rightInt)
			}
			// else {
			// 	fmt.Printf("leftIntErr: %v\n", leftIntErr)
			// 	fmt.Printf("rightIntErr: %v\n", rightIntErr)
			// }
		}
	}

	// fmt.Printf("left: %v\n", left)
	// fmt.Printf("right: %v\n", right)
	return left, right, nil
}
