package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := ReadInput("input_a.txt")
	histories := ParseInput(input)

	var nextValues []int
	for _, history := range histories {
		// Sequence index
		i := 0
		// Sequence of differences
		sequences := [][]int{history}
		// The next sequence for each prediction
		nextSequence := make([]int, len(history)-1)

		// Prediction loops
		for {
			for j := 0; j < len(sequences[i])-1; j++ {
				nextSequence[j] = sequences[i][j+1] - sequences[i][j]
			}

			sequences = append(sequences, nextSequence)

			// Check if we're done
			sum := 0
			for _, stepValue := range nextSequence {
				if stepValue < 0 {
					sum -= stepValue
				} else {
					sum += stepValue
				}
			}

			i++
			if sum == 0 {
				// Adding a 0 to the next sequence before extrapolating
				sequences[i] = append(sequences[i], 0)
				break
			} else {
				// Setting an empty new array with a known lenght
				nextSequence = make([]int, len(sequences[i])-1)
			}
		}

		// Extrapolation loops
		for {
			currentSequence := sequences[i][len(sequences[i])-1]
			prevSequence := sequences[i-1][len(sequences[i-1])-1]
			sequences[i-1] = append(sequences[i-1], currentSequence+prevSequence)

			if i == 1 {
				break
			} else {
				i--
			}
		}

		nextValue := sequences[0][len(sequences[0])-1]
		nextValues = append(nextValues, nextValue)
	}

	result := 0
	for _, curr := range nextValues {
		result += curr
	}

	fmt.Println("Result 🎊", result)
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
