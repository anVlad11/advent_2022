package main

import (
	"fmt"
	"strings"
)

type Grid map[int]map[int]string

func (g Grid) refreshGridPartOne() {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, exists := g[x]; !exists {
				g[x] = map[int]string{}
			}
			if _, exists := g[x][y]; !exists {
				g[x][y] = "."
			}
			if x == sourceX && y == sourceY && g[x][y] != "o" {
				g[x][y] = "+"
			}
		}
	}
	for x := minX; x <= maxX; x++ {
		g[x][maxY] = "#"
	}
}

func (g Grid) refreshGridPartTwo() {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, exists := g[x]; !exists {
				g[x] = map[int]string{}
			}
			if _, exists := g[x][y]; !exists {
				g[x][y] = "."
			}
			if x == sourceX && y == sourceY && g[x][y] != "o" {
				g[x][y] = "+"
			}
		}
	}
	for x := minX; x <= maxX; x++ {
		g[x][maxY] = "#"
	}
}

func (g Grid) printGrid() {
	sb := strings.Builder{}

	for y := minY; y <= maxY; y++ {
		sb.WriteString(fmt.Sprintf("%d\t", y))
		for x := minX; x <= maxX; x++ {
			if _, exists := g[x]; !exists {
				g[x] = map[int]string{}
			}
			if _, exists := g[x][y]; !exists {
				g[x][y] = "."
			}
			if x == sourceX && y == sourceY && g[x][y] != "o" {
				g[x][y] = "+"
			}

			val := g[x][y]

			sb.WriteString(val)
		}
		sb.WriteString("\n")
	}

	sb.WriteString("\n")

	fmt.Println(sb.String())
}
