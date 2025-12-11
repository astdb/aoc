package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// read input file - get file name from stdin and return an array of Turn objects
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./d1 <input file name>")
	}

	inputFile := os.Args[1]
	turns, err := processInput(inputFile)
	if err != nil {
		log.Fatalf("Error processing input: %s\n", err)
	}

	// create lock dial
	dial := &Dial{Point: 50}
	zCount := 0
	for _, turn := range turns {
		if turn.Direction == "L" {
			for i := 0; i < turn.Count; i++ {
				dial.Left()
				zCount = checkZCount(dial.Point, zCount)
			}
		}

		if turn.Direction == "R" {
			for i := 0; i < turn.Count; i++ {
				dial.Right()
				zCount = checkZCount(dial.Point, zCount)
			}
		}

		// zCount = checkZCount(dial.Point, zCount)
	}

	fmt.Println(zCount)
	return
}

func checkZCount(dpoint , zcount int) int {
	if dpoint == 0 {
		zcount++
	}

	return zcount
}

func processInput(fileName string) ([]*Turn, error) {
	turns := []*Turn{}

	file, err := os.Open(fileName)
	if err != nil {
		return turns, err
	}

	inputScanner := bufio.NewScanner(file)
	for inputScanner.Scan() {
		inputLine := inputScanner.Text()

		// fmt.Println(inputLine)
		turnDir := inputLine[:1]
		if !(turnDir == "L" || turnDir == "R") {
			log.Printf("Incorrect direction on input line: %s\n", inputLine)
			continue
		}

		count, err := strconv.Atoi(inputLine[1:])
		if err != nil {
			log.Printf("Error getting valid turn count from input line %s: %v\n", inputLine, err)
			continue
		}

		turn := &Turn{Direction: turnDir, Count: count}
		turns = append(turns, turn)
	}

	return turns, err
}

type Turn struct {
	Direction string
	Count     int
}

type Dial struct {
	Point int
}

func (d *Dial) Left() {
	if d.Point == 0 {
		d.Point = 99
		return
	}

	d.Point--
	return
}

func (d *Dial) Right() {
	if d.Point == 99 {
		d.Point = 0
		return
	}

	d.Point++
	return
}
