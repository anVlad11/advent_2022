package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part1(file []string) error {
	if len(file) == 0 {
		return nil
	}

	head := map[int]map[int]bool{
		0: {0: true},
	}
	prevHeadX := 0
	prevHeadY := 0
	headX := 0
	headY := 0

	tail := map[int]map[int]bool{
		0: {0: true},
	}
	prevTailX := 0
	prevTailY := 0
	tailX := 0
	tailY := 0

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
		for i := 0; i < steps; i++ {
			distance := 0.0

			switch direction {
			case "L":
				{
					headX--
					distance = math.Sqrt(math.Pow(float64(headX-tailX), 2) + math.Pow(float64(headY-tailY), 2))
					if distance >= 2 {
						tailX--
						tailY = headY
					}
				}
			case "R":
				{
					headX++
					distance = math.Sqrt(math.Pow(float64(headX-tailX), 2) + math.Pow(float64(headY-tailY), 2))
					if distance >= 2 {
						tailX++
						tailY = headY
					}
				}
			case "U":
				{
					headY++
					distance = math.Sqrt(math.Pow(float64(headX-tailX), 2) + math.Pow(float64(headY-tailY), 2))
					if distance >= 2 {
						tailY++
						tailX = headX
					}
				}
			case "D":
				{
					headY--
					distance = math.Sqrt(math.Pow(float64(headX-tailX), 2) + math.Pow(float64(headY-tailY), 2))
					if distance >= 2 {
						tailY--
						tailX = headX
					}
				}
			}

			head[prevHeadX][prevHeadY] = false
			if _, exists := head[headX]; !exists {
				head[headX] = map[int]bool{}
			}
			head[headX][headY] = true

			tail[prevTailX][prevTailY] = false
			if _, exists := tail[tailX]; !exists {
				tail[tailX] = map[int]bool{}
			}
			if _, exists := tail[tailX][tailY]; !exists {
				visitedByTail++
			}
			tail[tailX][tailY] = true

			prevHeadX = headX
			prevHeadY = headY
			prevTailX = tailX
			prevTailY = tailY

		}
	}

	fmt.Println(visitedByTail)

	return nil
}
