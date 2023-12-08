package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	idRegex          = regexp.MustCompile(`[A-Z|1-9]{3}`)
	startNodeIdRegex = regexp.MustCompile(`..A`)
	endNodeIdRegex   = regexp.MustCompile(`..Z`)
)

type Node struct {
	Id    string
	Left  string
	Right string
}

type Ghost struct {
	NodeId string
}

func main() {
	input, err := os.ReadFile("input_b.txt")
	if err != nil {
		panic(err)
	}

	inputString := string(input)
	inputLines := strings.Split(inputString, "\n\r\n")

	// Actual data
	instructions := strings.Split(strings.TrimSpace(inputLines[0]), "")
	nodes := strings.Split(inputLines[1], "\n")

	// Creating the map of nodes and a list of ghosts
	var ghosts []*Ghost
	mapNodes := make(map[string]Node)

	// Populating the map of nodes and ghosts to start with
	for _, node := range nodes {
		matches := idRegex.FindAllString(node, 3)
		node := Node{
			Id:    matches[0],
			Left:  matches[1],
			Right: matches[2],
		}

		if startNodeIdRegex.MatchString(node.Id) {
			ghosts = append(ghosts, &Ghost{
				NodeId: node.Id,
			})
		}

		mapNodes[node.Id] = node
	}

	step := 1
outerLoop:
	for {
		for _, direction := range instructions {
			for _, ghost := range ghosts {
				fmt.Println(step, direction, ghost)

				if direction == "R" {
					ghost.NodeId = mapNodes[ghost.NodeId].Right
				} else {
					ghost.NodeId = mapNodes[ghost.NodeId].Left
				}
			}

			ghostEnd := 0
			for _, ghostPosition := range ghosts {
				if endNodeIdRegex.MatchString(ghostPosition.NodeId) {
					ghostEnd++
				}
			}

			if ghostEnd == len(ghosts) {
				break outerLoop
			} else {
				step++
			}
		}
	}

	fmt.Println("Day 8, Part B: ", step)
}
