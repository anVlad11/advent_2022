package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part2(file []string) error {
	if len(file) == 0 {
		return nil
	}

	grid := Grid{}

	grid[500] = map[int]string{0: "+"}

	for _, line := range file {
		coordsRaw := strings.Split(line, " -> ")
		x := 0
		y := 0
		for _, coordRaw := range coordsRaw {
			coord := strings.Split(coordRaw, ",")
			targetX, _ := strconv.Atoi(coord[0])
			targetY, _ := strconv.Atoi(coord[1])
			if x == 0 && y == 0 {
				x, y = targetX, targetY
				continue
			}

			incX := 0
			if targetX < x {
				incX = -1
			} else if targetX > x {
				incX = 1
			}
			incY := 1
			if targetY < y {
				incY = -1
			} else if targetY > x {
				incX = 1
			}

			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}

			if _, exists := grid[x]; !exists {
				grid[x] = map[int]string{}
			}
			grid[x][y] = "#"

			for !(targetX == x && targetY == y) {
				if targetX != x {
					x += incX
				}
				if targetY != y {
					y += incY
				}
				if _, exists := grid[x]; !exists {
					grid[x] = map[int]string{}
				}
				grid[x][y] = "#"

				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
		x, y = 0, 0
	}

	maxY = maxY + 2

	fmt.Println(minX, maxX)

	sandCount := 0
	intoTheAbyss := false
	for !intoTheAbyss {
		x, y := sourceX, sourceY
		wall := grid[x][y]
		sandCount++
		for {
			if _, exists := grid[x]; !exists {
				intoTheAbyss = true
				break
			}
			if _, exists := grid[x][y]; !exists {
				intoTheAbyss = true
				break
			}
			grid[x][y] = "o"
			grid.refreshGridPartTwo()
			if grid[x][y+1] == "." {
				grid[x][y] = wall
				y++
				wall = grid[x][y]
				continue
			}

			if _, exists := grid[x-1]; !exists {
				grid[x-1] = map[int]string{}
				minX = x - 1
			}

			if _, exists := grid[x-1][y+1]; !exists {
				if y+1 == maxY {
					grid[x-1][y+1] = "#"
				} else {
					grid[x-1][y+1] = "."
				}
			}

			if grid[x-1][y+1] == "." {
				grid[x][y] = wall
				x--
				y++
				wall = grid[x][y]
				continue
			}

			if _, exists := grid[x+1]; !exists {
				grid[x+1] = map[int]string{}
				maxX = x - 1
			}

			if _, exists := grid[x+1][y+1]; !exists {
				if y+1 == maxY {
					grid[x+1][y+1] = "#"
				} else {
					grid[x+1][y+1] = "."
				}
			}

			if grid[x+1][y+1] == "." {
				grid[x][y] = wall
				x++
				y++
				wall = grid[x][y]
				continue
			}

			grid[x][y] = "o"
			break
		}
		//time.Sleep(25 * time.Millisecond)
		//fmt.Print("\033[H\033[2J")
		//grid.printGrid()
		println(sandCount)
		if x == sourceX && y == sourceY {
			break
		}
	}

	grid.printGrid()

	fmt.Println(sandCount)

	return nil
}
