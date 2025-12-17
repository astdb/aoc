package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./d2 <inputfilename>")
	}

	inputFileName := os.Args[1]
	input, err := processInput(inputFileName)
	if err != nil {
		log.Fatalf("Error processing input file: %v\n", err)
	}
	// fmt.Println(input)

	tot := 0
	for _, inputRange := range input {
		// inputRange := []int{11,22}
		for i := inputRange[0]; i <= inputRange[1]; i++ {
			iStr := strconv.Itoa(i)

			for subStrLen := 1; subStrLen <= len(iStr)/2; subStrLen++ {
				x := subStrLen
				subStrFound := true
				for y := subStrLen; y <= len(iStr)-x; y+=subStrLen {
					z := y + subStrLen
					if z > len(iStr) {
						z = len(iStr)
					}

					// fmt.Printf(" - Comparing: str[:%d] != str[%d:%d] (%s != %s) [subStrFound => %v]\n", x, y, z, iStr[:x], iStr[y:z], subStrFound)
					if iStr[:x] != iStr[y:z] {
						// fmt.Printf(" - Setting subStrFound to false..\n")
						subStrFound = false		// indicates that the search for subStrLen matching substring was false
						break
					}
				}

				if subStrFound && len(iStr)%subStrLen == 0 {
					// fmt.Printf(" - Matching substring found (iStr: %s, subStrLen: %d)\n", iStr, subStrLen)
					tot = tot + i
					subStrFound = true
					break	// go to next iStr
				}
			}
		}
	}

	fmt.Println(tot)
}

func processInput(filename string) ([][]int, error) {
	result := [][]int{}

	input, err := os.ReadFile(filename)
	if err != nil {
		return result, err
	}

	rangesStr := strings.Split(strings.TrimSpace(string(input)), ",")

	for _, rangeStr := range rangesStr {
		ranges := strings.Split(rangeStr, "-")

		if len(ranges) < 2 {
			log.Printf("Invalid range: %s\n", rangeStr)
			continue
		}

		lower, err := strconv.Atoi(strings.TrimSpace(ranges[0]))
		if err != nil {
			log.Printf("Invalid lower bound (%s): %v\n", err)
		}

		upper, err := strconv.Atoi(strings.TrimSpace(ranges[1]))
		if err != nil {
			log.Printf("Invalid upper bound (%s): %v\n", err)
		}

		res := []int{lower, upper}
		result = append(result, res)
	}

	return result, nil
}
