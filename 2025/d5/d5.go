package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./d5 <inputfile>")
	}

	freshRanges, ingIDs := processInput(os.Args[1])
	fmt.Printf("FreshRanges: %v\n", freshRanges)
	fmt.Printf("Ingredient IDs: %v\n", ingIDs)
}

func processInput(filename string) ([][]int,[]int) {
	freshRanges := [][]int{}
	ingIDs := []int{}
	
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Cannot open input file: %v\n", err)
	}

	inputScanner := bufio.NewScanner(file)

	freshRangeRE := regexp.MustCompile("^[0-9]+\\-[0-9]+$")
	ingIDRE := regexp.MustCompile("^[0-9]+$")

	for inputScanner.Scan() {
		line := inputScanner.Text()

		if freshRangeRE.MatchString(strings.TrimSpace(line)) {
			// fresh ingredient ID range
			idRange, err := getIDRange(line)
			if err != nil {
				log.Printf("Creating ID range: %v\n", err)
				continue
			}

			freshRanges = append(freshRanges, idRange)
		} else if ingIDRE.MatchString(line) {
			// ingredient ID
			ingID, err := strconv.Atoi(line)
			if err != nil {
				log.Printf("Converting ingredient ID: %v\n", err)
				continue
			}

			ingIDs = append(ingIDs, ingID)
		}
	}

	return freshRanges, ingIDs
}

func getIDRange(line string) ([]int, error) {
	freshRange := []int{}
	lineComps := strings.Split(line, "-")

	if len(lineComps) != 2 {
		return freshRange, errors.New(fmt.Sprintf("Invalid fresh range string: %s", line))
	}

	lower, err := strconv.Atoi(lineComps[0])
	if err != nil {
		return freshRange, errors.New(fmt.Sprintf("Invalid lower range: %s\n", lineComps[0]))
	}

	upper, err := strconv.Atoi(lineComps[1])
	if err != nil {
		return freshRange, errors.New(fmt.Sprintf("Invalid upper range: %s\n", lineComps[1]))
	}

	freshRange = append(freshRange, lower, upper)
	return freshRange, nil
}