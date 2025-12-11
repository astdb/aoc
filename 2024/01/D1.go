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
	leftList, rightList, err := processInputFile(os.Getenv("D1INPUT"))
	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)
	// fmt.Printf("leftList: %v\n", leftList)
	// fmt.Printf("rightList: %v\n", rightList)

	res := 0
	for i := 0; i < len(leftList); i++ {
		if i < len(rightList) {
			res += Abs(leftList[i] - rightList[i])
		}
	}

	fmt.Println(res)
}

func Abs(x int) int {
	if x < 0 {
		return -1 * x
	}

	return x
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
