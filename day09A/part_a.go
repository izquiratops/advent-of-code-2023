package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	input := ReadInput("test_a.txt")
	histories := ParseInput(input)

	for _, history := range histories {
		prevHistory, nextHistory := SettingNextPrediction(nil, history)
		predictionIters := 0

	predictionLoop:
		for {
			for i := 0; i < len(prevHistory)-1; i++ {
				nextHistory[i] = prevHistory[i+1] - prevHistory[i]
			}

			// Check if we're done
			sum := 0
			for _, stepValue := range nextHistory {
				sum += stepValue
			}

			// Moving on to the next prediction or breaking the loop if every step is 0
			if sum == 0 {
				break predictionLoop
			} else {
				prevHistory, nextHistory = SettingNextPrediction(prevHistory, nextHistory)
				predictionIters++
			}
		}

		prevHistory = append(prevHistory, prevHistory[len(prevHistory)-1])
		prevHistory, nextHistory = SettingNextExtrapolation(nil, prevHistory)
		extrapolationIters := 0

	extrapolationLoop:
		for {
			value := prevHistory[0]
			for j := 0; j < len(prevHistory)-1; j++ {
				value += prevHistory[j+1]
				nextHistory[j] = value
			}

			// Checking if we have reached the desired iteration
			if extrapolationIters == predictionIters {
				break extrapolationLoop
			} else {
				prevHistory, nextHistory = SettingNextExtrapolation(prevHistory, nextHistory)
				extrapolationIters++
			}
		}
	}
}

func ReadInput(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(input)
}

func ParseInput(input string) [][]int {
	inputLines := strings.Split(input, "\n")
	histories := make([][]int, len(inputLines))

	for i := 0; i < len(inputLines); i++ {
		// Getting values from the line as strings
		trimmedLine := strings.TrimSpace(inputLines[i])
		splittedLine := strings.Split(trimmedLine, " ")

		// Converting the strings to ints
		for _, item := range splittedLine {
			num, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			} else {
				histories[i] = append(histories[i], num)
			}
		}
	}

	return histories
}

func SettingNextPrediction(prev []int, next []int) ([]int, []int) {
	prev = make([]int, len(next))
	copy(prev, next)
	next = make([]int, len(prev)-1)
	return prev, next
}

func SettingNextExtrapolation(prev []int, next []int) ([]int, []int) {
	prev = make([]int, len(next))
	copy(prev, next)
	next = make([]int, len(prev)+1)
	return prev, next
}
