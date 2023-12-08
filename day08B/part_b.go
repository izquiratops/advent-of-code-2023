package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	Id    string
	Left  *Node
	Right *Node
}

func main() {
	input := ReadInput("input_b.txt")
	instructions, starterNodes := ParseInput(input)

	stepsPerGhost := make([]int, len(starterNodes))
	for i := 0; i < len(starterNodes); i++ {
		steps := 0
		for j := 0; strings.LastIndex(starterNodes[i].Id, "Z") != 2; j++ {
			// make sure we don't go out of bounds
			j %= len(instructions)
			// move the ghost
			steps++
			// take the next instruction
			if instructions[j] == "R" {
				starterNodes[i] = starterNodes[i].Right
			} else {
				starterNodes[i] = starterNodes[i].Left
			}
		}
		stepsPerGhost[i] = steps
	}

	fmt.Println(HandleSteps(stepsPerGhost))
}

func ReadInput(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(input)
}

func ParseInput(input string) ([]string, []*Node) {
	inputLines := strings.Split(input, "\n\r\n")

	// Intructions are a string of R's and L's
	instructions := strings.Split(strings.TrimSpace(inputLines[0]), "")
	// Nodes are represented as a string of 3 characters with 2 more for each direction
	nodeLines := strings.Split(inputLines[1], "\n")

	nodes := make(map[string]*Node, len(nodeLines))
	starterNodes := make([]*Node, 0)
	re := regexp.MustCompile(`[A-Z|1-9]{3}`)

	// Populate the nodes with their ids
	for _, n := range nodeLines {
		nodeId := re.FindString(n)
		nodes[nodeId] = &Node{
			Id: nodeId,
		}

		if strings.LastIndex(nodeId, "A") == 2 {
			starterNodes = append(starterNodes, nodes[nodeId])
		}
	}

	for _, n := range nodeLines {
		matches := re.FindAllString(n, 3)
		nodes[matches[0]].Left = nodes[matches[1]]
		nodes[matches[0]].Right = nodes[matches[2]]
	}

	return instructions, starterNodes
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func HandleSteps(numbers []int) int {
	result := numbers[0]

	for i := 1; i < len(numbers); i++ {
		result = LCM(result, numbers[i])
	}

	return result
}
