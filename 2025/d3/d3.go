/**
- set x to 9
- find the first 9 (before the last character) and the next biggest number after that - that's your answer. 
- decrease x and retry. 
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: ./d3 <input filename>\n")
	}

	inputFileName := os.Args[1]

	banks, err := processInput(inputFileName)
	if err != nil {
		log.Fatalf("Error processing input file: %v", err)
	}

	fmt.Printf("Banks: %v\n", banks)
}

func processInput(fileName string) ([][]int, error) {
	res := [][]int{}
	
	file, err := os.Open(fileName)
	if err != nil {
		return res, err
	}

	inputScanner := bufio.NewScanner(file)
	for inputScanner.Scan() {
		inputLine := inputScanner.Text()

		thisBattBank, err := getNumsArr(inputLine)
		if err != nil {
			log.Printf("Error coverting battery bank to numeric: %v\n", err)
			continue
		}

		res = append(res, thisBattBank)
	}

	return res, nil
}

func getNumsArr(str string) ([]int, error) {
	bank := []int{}
	
	for _, ch := range str {
		i, err := strconv.Atoi(string(ch))
		if err != nil {
			return bank, err
		}

		bank = append(bank, i)
	}

	return bank, nil
}