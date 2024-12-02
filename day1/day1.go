package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readNumbersFromFile(filename string) ([]int, []int, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var leftNumbers []int
	var rightNumbers []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		numberStrings := strings.Fields(line)

		if len(numberStrings) < 2 {
			continue
		}

		leftNum, err := strconv.Atoi(numberStrings[0])
		if err != nil {
			return nil, nil, fmt.Errorf("error converting left numbers: %v", err)
		}
		leftNumbers = append(leftNumbers, leftNum)

		rightNum, err := strconv.Atoi(numberStrings[1])
		if err != nil {
			return nil, nil, fmt.Errorf("error converting right numbers: %v", err)
		}
		rightNumbers = append(rightNumbers, rightNum)
	}

	return leftNumbers, rightNumbers, nil
}

func main() {
	leftNums, rightNums, err := readNumbersFromFile("input.txt")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	// fmt.Printf("variable leftNums is of type %T \n", leftNums)
	// fmt.Printf("variable rightNums is of type %T \n", rightNums)

	sort.Sort(sort.IntSlice(leftNums))
	sort.Sort(sort.IntSlice(rightNums))

	// fmt.Println("Sorted Left numbers: ", leftNums)
	// fmt.Println("Sorted right numbers: ", rightNums)

	var distance []int
	var dis int

	for i := 0; i < len(leftNums); i++ {
		difference := leftNums[i] - rightNums[i]
		positiveNumbers := int(math.Abs(float64(difference)))
		distance = append(distance, positiveNumbers)
		dis += positiveNumbers
	}

	fmt.Println("Difference between distances: ", distance)
	fmt.Println("Addition all distances: ", dis)
}
