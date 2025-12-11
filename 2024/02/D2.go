package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// process input file, and obtain set of data arrays
	inputData := processInputFile(os.Getenv("D2INPUT"))

	// verify safety of data per input file
	// fmt.Println(inputData)

	safeCount := 0
	for _, row := range inputData {
		safe := isSafe((row))
		fmt.Printf("%v : %v\n", row, safe)
		if safe {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func isSafe(dataRow []int) bool {
	safe := true
	incr := false // series is increasing

	for i := 1; i < len(dataRow); i++ {
		fmt.Printf("isSafe(): comparing %d to %d..\n", dataRow[i-1], dataRow[i])
		fmt.Printf("isSafe(): %d - %d = Abs(%d)\n", dataRow[i-1], dataRow[i], Abs(dataRow[i-1]-dataRow[i]))

		if Abs(dataRow[i-1]-dataRow[i]) < 1 && Abs(dataRow[i-1]-dataRow[i]) > 3 {
			fmt.Printf("isSafe(): non-safe datapoint distance\n")
			safe = false
			return safe
		}

		fmt.Printf("isSafe(): safe datapoint distance - considering inc/dec\n")
		if i == 1 {
			fmt.Printf("isSafe(): safe datapoint distance - considering inc/dec\n")
			if dataRow[i-1]-dataRow[i] < 0 {
				incr = true
			} else {
				incr = false
			}
		} else {
			if (incr && dataRow[i-1]-dataRow[i] > 0) || (!incr && dataRow[i-1]-dataRow[i] < 0) {
				safe = false
				return safe
			}
		}
	}

	return safe
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

func processInputFile(inputFile string) [][]int {
	inputData := [][]int{}
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	inputScanner := bufio.NewScanner(file)
	for inputScanner.Scan() {
		line := inputScanner.Text()
		lineCompsStr := strings.Split(line, " ")
		lineCompsInt := intDataSlice(lineCompsStr)
		inputData = append(inputData, lineCompsInt)
	}

	return inputData
}

func intDataSlice(strSlice []string) []int {
	res := []int{}
	for _, v := range strSlice {
		intVal, err := strconv.Atoi(v)
		if err == nil {
			res = append(res, intVal)
		}
	}

	return res
}
