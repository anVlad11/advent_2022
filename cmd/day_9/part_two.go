package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part2(file []string) error {
	if len(file) == 0 {
		return nil
	}

	ropeLength := 9

	rope := map[int]map[int]map[int]bool{}
	previousSegmentXs := map[int]int{}
	previousSegmentYs := map[int]int{}
	segmentXs := map[int]int{}
	segmentYs := map[int]int{}

	for i := 0; i <= ropeLength; i++ {
		rope[i] = map[int]map[int]bool{0: {0: true}}
		previousSegmentXs[i] = 0
		previousSegmentYs[i] = 0
		segmentXs[i] = 0
		segmentYs[i] = 0
	}

	visitedByTail := 1
	for _, line := range file {
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			continue
		}

		direction := parts[0]
		steps, _ := strconv.Atoi(parts[1])
		for step := 0; step < steps; step++ {

			switch direction {
			case "L":
				{
					segmentXs[0]--
				}
			case "R":
				{
					segmentXs[0]++
				}
			case "U":
				{
					segmentYs[0]++
				}
			case "D":
				{
					segmentYs[0]--
				}
			}
			rope[0][previousSegmentXs[0]][previousSegmentYs[0]] = false
			if _, exists := rope[0][segmentXs[0]]; !exists {
				rope[0][segmentXs[0]] = map[int]bool{}
			}
			rope[0][segmentXs[0]][segmentYs[0]] = true
			previousSegmentXs[0] = segmentXs[0]
			previousSegmentYs[0] = segmentYs[0]

			for segmentI := 1; segmentI <= ropeLength; segmentI++ {
				headX := segmentXs[segmentI-1]
				headY := segmentYs[segmentI-1]

				tailX := segmentXs[segmentI]
				tailY := segmentYs[segmentI]

				distX := headX - tailX
				distY := headY - tailY

				distXPow := math.Pow(float64(distX), 2)
				distYPow := math.Pow(float64(distY), 2)

				distance := math.Sqrt(distXPow + distYPow)
				if distance == 2 {
					if distX < 0 {
						segmentXs[segmentI]--
					}
					if distX > 0 {
						segmentXs[segmentI]++
					}
					if distY < 0 {
						segmentYs[segmentI]--
					}
					if distY > 0 {
						segmentYs[segmentI]++
					}
				} else if distance > 2 {
					if distXPow > distYPow {
						if distX < 0 {
							segmentXs[segmentI]--
						}
						if distX > 0 {
							segmentXs[segmentI]++
						}
						segmentYs[segmentI] = segmentYs[segmentI-1]
					} else if distXPow < distYPow {
						if distY < 0 {
							segmentYs[segmentI]--
						}
						if distY > 0 {
							segmentYs[segmentI]++
						}
						segmentXs[segmentI] = segmentXs[segmentI-1]
					} else {
						if distX < 0 {
							segmentXs[segmentI]--
						}
						if distX > 0 {
							segmentXs[segmentI]++
						}
						if distY < 0 {
							segmentYs[segmentI]--
						}
						if distY > 0 {
							segmentYs[segmentI]++
						}
					}
				}

				rope[segmentI][previousSegmentXs[segmentI]][previousSegmentYs[segmentI]] = false
				if _, exists := rope[segmentI][segmentXs[segmentI]]; !exists {
					rope[segmentI][segmentXs[segmentI]] = map[int]bool{}
				}
				if segmentI == ropeLength {
					if _, exists := rope[segmentI][segmentXs[segmentI]][segmentYs[segmentI]]; !exists {
						visitedByTail++
					}
				}
				rope[segmentI][segmentXs[segmentI]][segmentYs[segmentI]] = true

				previousSegmentXs[segmentI] = segmentXs[segmentI]
				previousSegmentYs[segmentI] = segmentYs[segmentI]

			}
		}

	}

	fmt.Println(visitedByTail)

	return nil
}
