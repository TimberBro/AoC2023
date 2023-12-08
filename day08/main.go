package main

import (
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	log.Println("Part 1: ", firstPart("day08\\input.txt"))

	log.Println("Part 2: ", secondPart("day08\\input.txt"))
}

func firstPart(file string) int {
	fileContent, err := os.ReadFile(file)
	if err != nil {
		log.Printf("unable to read file: %v\n", err)
	}

	fileString := string(fileContent)
	navigation, nodes := parseMap(fileString)

	result := findFinalNode(navigation, nodes)
	return result
}

func secondPart(file string) int {
	fileContent, err := os.ReadFile(file)
	if err != nil {
		log.Printf("unable to read file: %v\n", err)
	}

	fileString := string(fileContent)
	navigation, nodes := parseMap(fileString)

	result := countGhostSteps(navigation, nodes)
	return result
}

func parseMap(line string) (string, map[string]Node) {
	pattern := `(\w+)`
	patternReg, _ := regexp.Compile(pattern)

	navigation := patternReg.FindAllString(line, -1)[:1]
	//log.Println(navigation)

	network := patternReg.FindAllString(line, -1)[1:]
	//log.Println(network)

	nodes := make(map[string]Node, 0)
	for i := 0; i < len(network); i += 3 {
		node := Node{}
		node.LeftValue = network[i+1]
		node.RightValue = network[i+2]
		nodes[network[i]] = node
	}
	//log.Printf("nodes: %+v", nodes)

	return navigation[0], nodes
}

// First part
func findFinalNode(navigation string, nodes map[string]Node) int {
	result := 0
	foundNode := false
	var currentNode string
	for {
		if len(currentNode) == 0 {
			currentNode = "AAA"
		}
		for _, instruction := range navigation {
			//log.Printf("Begin of the navigation\n")
			if currentNode != "ZZZ" {
				switch string(instruction) {
				case "R":
					currentNode = nodes[currentNode].RightValue
					//log.Printf("picking right side, next node: %v\n", currentNode)
					result++
				case "L":
					currentNode = nodes[currentNode].LeftValue
					//log.Printf("picking left side, next node: %v\n", currentNode)
					result++
				}
			} else {
				//log.Printf("node ZZZ found, exiting from loop\n")
				foundNode = true
				break
			}
		}

		if foundNode {
			//log.Printf("node ZZZ found, exiting from infinite loop\n")
			break
		}
	}

	return result
}

// Second part
func countGhostSteps(navigation string, nodes map[string]Node) int {
	result := 0

	keys := make([]string, len(nodes))
	for s := range nodes {
		keys = append(keys, s)
	}

	finalStep := false
	startingNodes := make([]string, 0)
	for {
		// init starting list
		if len(startingNodes) == 0 {
			startingNodes = findStartingNodes(keys)
		}
		for _, instruction := range navigation {
			//log.Printf("Begin of the navigation\n")

			//log.Printf("Before step: %+v", startingNodes)
			startingNodes = makeStep(string(instruction), nodes, startingNodes)
			result++
			//log.Printf("After step: %+v", startingNodes)
			if isFinalStep(startingNodes) {
				//log.Printf("final step: %v\n Exiting loop\n", startingNodes)
				finalStep = true
				break
			}
		}

		if finalStep {
			//log.Printf("exiting from infinite loop\n")
			break
		}
	}

	return result
}

func makeStep(navigation string, nodes map[string]Node, travelingNodes []string) []string {
	resultNodes := make([]string, 0)
	for _, node := range travelingNodes {
		switch navigation {
		case "R":
			resultNodes = append(resultNodes, nodes[node].RightValue)
			//log.Printf("picking right side, next node: %v\n", node)
		case "L":
			resultNodes = append(resultNodes, nodes[node].LeftValue)
			//log.Printf("picking left side, next node: %v\n", node)
		}
	}

	return resultNodes
}

func findStartingNodes(nodes []string) []string {
	result := make([]string, 0)

	for _, node := range nodes {
		if strings.HasSuffix(node, "A") {
			result = append(result, node)
		}
	}

	return result
}

func isFinalStep(nodes []string) bool {
	counter := 0

	for _, node := range nodes {
		if strings.HasSuffix(node, "Z") {
			counter++
		}
	}

	return counter == len(nodes)
}

type Node struct {
	LeftValue  string
	RightValue string
}
