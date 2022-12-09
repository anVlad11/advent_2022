package main

import "fmt"

func printField(head, tail map[int]map[int]bool) {
	result := map[int]map[int]string{0: {0: "s"}}

	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	visitedByTail := 0

	for x, m := range tail {
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
		if _, exists := result[x]; !exists {
			result[x] = map[int]string{}
		}
		for y, v := range m {
			if y > maxY {
				maxY = y
			}
			if y < minY {
				minY = y
			}
			result[x][y] = "#"
			if v {
				result[x][y] = "T"
			}
			visitedByTail++
		}
	}

	for x, m := range head {
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
		if _, exists := result[x]; !exists {
			result[x] = map[int]string{}
		}
		for y, v := range m {
			if y > maxY {
				maxY = y
			}
			if y < minY {
				minY = y
			}
			if v {
				result[x][y] = "H"
			}
		}
	}

	result[0][0] = "s"

	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			v, exists := result[x][y]
			if !exists {
				fmt.Print(".")
			} else {
				fmt.Print(v)
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func printField2(field map[int]map[int]map[int]bool) {
	result := map[int]map[int]string{0: {0: "s"}}

	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	for segmentI := len(field) - 1; segmentI >= 0; segmentI-- {
		segment := field[segmentI]
		value := fmt.Sprintf("%d", segmentI)
		if segmentI == 0 {
			value = "H"
		}
		for x, m := range segment {
			if x > maxX {
				maxX = x
			}
			if x < minX {
				minX = x
			}
			if _, exists := result[x]; !exists {
				result[x] = map[int]string{}
			}
			for y, v := range m {
				if y > maxY {
					maxY = y
				}
				if y < minY {
					minY = y
				}
				if v && result[x][y] != "#" {
					result[x][y] = value
				}
				if segmentI == len(field)-1 {
					result[x][y] = "#"
				}
			}
		}
	}

	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			v, exists := result[x][y]
			if !exists {
				fmt.Print(".")
			} else {
				fmt.Print(v)
			}
		}
		fmt.Println()
	}
}
