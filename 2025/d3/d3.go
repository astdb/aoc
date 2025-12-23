/**
- set x to 9
- find the first x (before the last character) and the next biggest number after that - that's your answer. 
- decrease x and retry. 
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	var totMaxJ int64 = 0

	for _, bank := range banks {
		// maxJ := getJoltage(bank)
		// totMaxJ = totMaxJ + maxJ

		maxJ := getJoltage2(bank)
		totMaxJ = totMaxJ + int64(maxJ)
	}

	fmt.Println(totMaxJ)

}

// given a 15-digit string, determine the largest integer that can be made with 12 digits, without changing their order.
// Note: previously we built a tw-digit number with the same constraints. There we started looking for a large starting digit
// and the next-largest digit from the rest of the string and joining them together. Can that be extended to larger-digit count
// ints? How could it work for say a three-digit one? 

// start by picking the largest in the first four - given the int we're creating is 12-digits long, we can't start later than that.
// we HAVE to pick the first digit from thefirst four. 
// Then it gets interesting - if we pick the fourth one, then we have no choice, we have to grab everything that follows. But if we pick
// something that's either 1st/2nd or 3rd, then we have a choice. If we pick the first we can drop further three digits, 2nd we can drop 2,
// 3rd we can drop 1 and 4th - none. 

// So lets say we found the largest starting digit (and its not the 4th) - how do we determine the rest? 
// Yes - that's what we do. So let's say we the largest of the first four digits is indeed the very first. That means, we can drop a further 
// three digits. So after the first pick we try to determine the largest of positions 2, 3, 4 and 5. Lets say we picked 4 

// Set toDrop = 3 (this means the number of digits we can 'drop' to construct the 12-digit int)
// Find the largest digit in str[:toDrop+1] range. Let's say the index is x. Make note of it. 
// Decrement toDrop by x (toDrop = toDrop-x). 
// Find the largest digit in [x:toDrop+1]
// if toDrop is zero, we'll need to add all remaining digits.

func getJoltage2(bank []int) int {
	joltage := []int{}
	toDrop := len(bank) - 12

	start := 0
	end := start + (toDrop + 1) 
	fmt.Printf("joltage: %v / toDrop: %d / start: %d / End: %d\n", joltage, toDrop, start, end)

	for len(joltage) < 12 {
		largest, index := largest2(start, end, bank)
		fmt.Printf(" ---------------------------- \nLargest: %d / index: %d / slice: %v\n", largest, index, bank[start:end])

		joltage = append(joltage, largest)

		toDrop = toDrop - (index-start)
		start = index + 1
		end = index + toDrop + 1 + 1

		fmt.Printf("joltage: %v (length: %d) / toDrop: %d / start: %d / End: %d\n", joltage, len(joltage), toDrop, start, end)
	}

	if len(joltage) < 12 {
		joltage = append(joltage, bank[start:]...)
	}
	
	fmt.Println(joltage)
	joltageInt := toInt(joltage)
	fmt.Println(joltageInt)
	return joltageInt
}

func toInt(jolt []int) int {
	i := 0
	pow := 0
	base := 10
	for k := len(jolt)-1; k >= 0; k-- {
		i = i + int(math.Pow(float64(base), float64(pow))) * jolt[k]
		pow = pow + 1
	}

	return i
}

// largest2 accepts an int slice (bank), and a start and end index. It then returns the largest int within bank[start:end], and its index
func largest2(start, end int, bank []int) (int, int) {
	largest := 0
	index := 0

	for i := start; i < end; i++ {
		if bank[i] > largest {
			largest = bank[i]
			index = i
		}
	}

	return largest, index
}

func getJoltage(bank []int) (int) {
	// fmt.Println(bank)
	first := 0
	sec := 0

	found := false
	for x := 9; x >=0; x-- {
		// find first 9
		for i := 0; i < len(bank); i++ {
			if bank[i] == x && i < len(bank)-1 {
				first = bank[i]
				
				// find the next largest item in the rest of the bank
				sec = largest1(bank[i+1:])
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	return (10*first) + sec
}

func largest1(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
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