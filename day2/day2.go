package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readNumbersFromFile(filename string) ([][]int, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var allNumbers [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		numberStrings := strings.Fields(line)

		numbers := make([]int, len(numberStrings))

		for i, numStr := range numberStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, fmt.Errorf("error converting number %s: %v", numStr, err)
			}
			numbers[i] = num
		}

		allNumbers = append(allNumbers, numbers)

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return allNumbers, nil

}

func main() {
	// Read all numbers from the file
	numbers, err := readNumbersFromFile("sample.txt")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	fmt.Printf("%v \n", numbers)

	// Print out all lines of numbers
	fmt.Println("All numbers in the file:")
	for i, lineNumbers := range numbers {
		fmt.Printf("Line %d: %v\n", i+1, lineNumbers)
	}
}
