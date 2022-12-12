package main

import (
	"fmt"
	"strings"
)

func part1(file []string) error {
	if len(file) == 0 {
		return nil
	}

	topography := &Graph{}

	var sourceNode *Node
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
					sourceNode = node
				}
			case char == 'E':
				{
					char = 'z'
					targetNode = node
				}
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

	shortestPathLength := topography.FindShortestPath(sourceNode, targetNode)
	fmt.Println(shortestPathLength)

	return nil
}
