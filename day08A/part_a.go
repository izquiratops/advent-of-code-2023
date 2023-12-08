package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	re = regexp.MustCompile(`[A-Z]{3}`)
)

type Node struct {
	Id    string
	Left  string
	Right string
}

func main() {
	input, err := os.ReadFile("input_a.txt")
	if err != nil {
		panic(err)
	}

	inputString := string(input)
	inputLines := strings.Split(inputString, "\n\r\n")

	// Actual data
	instructions := strings.Split(strings.TrimSpace(inputLines[0]), "")
	nodes := strings.Split(inputLines[1], "\n")

	// Creating a map of nodes
	mapNodes := make(map[string]Node)

	// Populating the map of nodes
	for _, node := range nodes {
		matches := re.FindAllString(node, 3)
		node := Node{
			Id:    matches[0],
			Left:  matches[1],
			Right: matches[2],
		}

		mapNodes[node.Id] = node
	}

	currentNode := "AAA"
	step := 1
outerLoop:
	for {
		for _, direction := range instructions {
			fmt.Println(step, direction, currentNode)

			if direction == "R" {
				currentNode = mapNodes[currentNode].Right
			} else {
				currentNode = mapNodes[currentNode].Left
			}

			if currentNode == "ZZZ" {
				break outerLoop
			} else {
				step++
			}
		}
	}

	fmt.Println("Day 8, Part A: ", step)
}
