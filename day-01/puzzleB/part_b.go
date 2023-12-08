package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	words     = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	numberMap = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
)

func main() {
	// Reading the input file into a list of strings
	input, err := os.ReadFile("input_b.txt")
	if err != nil {
		panic(err)
	}

	inputString := string(input)
	inputLines := strings.Split(inputString, "\n")

	calibrationValues := list.New()
	for index, inputLine := range inputLines {
		matches := findValidMatch(inputLine)
		fmt.Println(index, matches)

		calibrationValue := getCalibrationValue(matches)
		fmt.Println(index, calibrationValue)

		calibrationValues.PushBack(calibrationValue)
	}

	// Summing up the calibration values
	total := 0
	for e := calibrationValues.Front(); e != nil; e = e.Next() {
		total += e.Value.(int)
	}

	fmt.Println("Day 1, Part B: ", total)
}

func findValidMatch(inputLine string) []string {
	var matches []string
	for i := 0; i < len(inputLine); i++ {
		for _, word := range words {
			if strings.HasPrefix(inputLine[i:], word) {
				fmt.Println("Found match: ", word)
				matches = append(matches, word)
			}
		}
	}

	return matches
}

func getCalibrationValue(matchedList []string) int {
	firstValue := getNumericValue(matchedList[0])
	lastValue := getNumericValue(matchedList[len(matchedList)-1])
	calibrationValue, err := strconv.Atoi(firstValue + lastValue)
	if err != nil {
		panic(err)
	}

	return calibrationValue
}

func getNumericValue(value string) string {
	number, ok := numberMap[value]
	if ok {
		return number
	} else {
		return value
	}
}
