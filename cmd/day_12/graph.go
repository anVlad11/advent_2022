package main

import (
	"fmt"
)

type Node struct {
	X     int
	Y     int
	Value int
	Edges []*Node
}

func (n *Node) GetID() string {
	return fmt.Sprintf("%d-%d", n.X, n.Y)
}

type FIFO struct {
	queue []*Node
}

func (f *FIFO) Push(node *Node) {
	if f.queue == nil {
		f.queue = []*Node{}
	}
	f.queue = append(f.queue, node)
}

func (f *FIFO) Pop() *Node {
	if len(f.queue) == 0 {
		return nil
	}

	node := f.queue[0]
	f.queue = f.queue[1:]

	return node
}

type Graph struct {
	graphMap map[int]map[int]*Node
	nodes    []*Node
	nodeMap  map[string]*Node
}

func (g *Graph) AddNode(node *Node) {
	if g.graphMap == nil {
		g.graphMap = map[int]map[int]*Node{}
	}
	if _, exists := g.graphMap[node.X]; !exists {
		g.graphMap[node.X] = map[int]*Node{}
	}
	g.graphMap[node.X][node.Y] = node

	if g.nodes == nil {
		g.nodes = []*Node{}
	}
	g.nodes = append(g.nodes, node)

	if g.nodeMap == nil {
		g.nodeMap = map[string]*Node{}
	}
	g.nodeMap[node.GetID()] = node
}

func (g *Graph) GetNode(x, y int) *Node {
	if g.graphMap == nil {
		return nil
	}

	if _, exists := g.graphMap[x]; !exists {
		return nil
	}

	val, exists := g.graphMap[x][y]
	if !exists {
		return nil
	}

	return val
}

func (g *Graph) GetNodes() []*Node {
	if g.nodes == nil {
		return []*Node{}
	}

	return g.nodes
}

func (g *Graph) GetNeighbours(node *Node) []*Node {
	neighbours := make([]*Node, 0, 4)
	left := g.GetNode(node.X-1, node.Y)
	if left != nil {
		neighbours = append(neighbours, left)
	}
	right := g.GetNode(node.X+1, node.Y)
	if right != nil {
		neighbours = append(neighbours, right)
	}
	up := g.GetNode(node.X, node.Y-1)
	if up != nil {
		neighbours = append(neighbours, up)
	}
	down := g.GetNode(node.X, node.Y+1)
	if down != nil {
		neighbours = append(neighbours, down)
	}

	return neighbours
}

func (g *Graph) FindShortestPath(from *Node, to *Node) int {
	if existing, exist := g.nodeMap[from.GetID()]; !exist || existing != from {
		return 0
	}

	if existing, exist := g.nodeMap[to.GetID()]; !exist || existing != to {
		return 0
	}

	visited := make(map[string]*Node)
	queue := &FIFO{}
	queue.Push(from)

	visited[from.GetID()] = from
	previousMap := map[*Node]*Node{}

	for node := queue.Pop(); node != nil; node = queue.Pop() {
		for _, edge := range node.Edges {
			if _, ok := visited[edge.GetID()]; !ok {
				previousMap[edge] = node
				if edge == to {
					break
				}
				visited[edge.GetID()] = edge
				queue.Push(edge)
			}
		}
	}

	depth := 0
	for previous := previousMap[to]; previous != nil; previous = previousMap[previous] {
		depth++
	}

	return depth
}
