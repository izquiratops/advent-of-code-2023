package main

import (
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Reading the input file into a list of strings
	input, err := os.ReadFile("input_a.txt")
	if err != nil {
		panic(err)
	}

	inputString := string(input)
	inputLines := strings.Split(inputString, "\n")

	// Setting up the regex pattern to find numbers
	pattern := "[1-9]"
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err)
		return
	}

	calibrationValues := list.New()
	for _, line := range inputLines {
		// Looking for all the numbers in each line
		matchedList := re.FindAllString(line, -1)
		if len(matchedList) == 0 {
			continue
		}

		// Looking for the first and last number in each line
		matchedValue := matchedList[0] + matchedList[len(matchedList)-1]

		// Converting the matched string to an integer
		calibrationValue, err := strconv.Atoi(matchedValue)
		if err != nil {
			panic(err)
		}

		calibrationValues.PushBack(calibrationValue)
	}

	// Summing up the calibration values
	total := 0
	for e := calibrationValues.Front(); e != nil; e = e.Next() {
		total += e.Value.(int)
	}

	fmt.Println("Day 1, Part A: ", total)
}
