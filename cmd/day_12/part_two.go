package main

import (
	"fmt"
	"math"
	"strings"
)

func part2(file []string) error {
	if len(file) == 0 {
		return nil
	}

	topography := &Graph{}

	sourceNodes := []*Node{}
	var targetNode *Node

	for x, line := range file {
		if strings.TrimSpace(line) == "" {
			continue
		}

		for y, v := range line {
			char := v

			node := &Node{
				X:     x,
				Y:     y,
				Edges: []*Node{},
			}

			switch {
			case char == 'S':
				{
					char = 'a'
				}
			case char == 'E':
				{
					char = 'z'
					targetNode = node
				}
			}

			if char == 'a' {
				sourceNodes = append(sourceNodes, node)
			}

			node.Value = int(char-'a') + 1

			topography.AddNode(node)
		}
	}

	for _, node := range topography.GetNodes() {
		for _, neighbour := range topography.GetNeighbours(node) {
			if neighbour.Value <= node.Value+1 {
				node.Edges = append(node.Edges, neighbour)
			}
		}
	}

	// task does not specify that you need to start from 'E'
	// this graph is not weighted, too
	// let's just find all distances from all 'a' to 'E' lol
	shortestPathLengthTotal := math.MaxInt
	for _, node := range sourceNodes {
		if node.Value == 1 {
			shortestPathLength := topography.FindShortestPath(node, targetNode)
			if shortestPathLength != 0 && shortestPathLength < shortestPathLengthTotal {
				shortestPathLengthTotal = shortestPathLength
			}
		}
	}

	fmt.Println(shortestPathLengthTotal)

	return nil
}
