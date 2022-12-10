package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/anVlad11/advent_2022/internal/helpers"
)

func main() {
	err := do()
	if err != nil {
		fmt.Println(err)
	}
}

func do() error {
	file, err := helpers.ReadFile("./data/input/day_10_input.txt")
	if err != nil {
		return err
	}

	err = part1(file)
	if err != nil {
		return err
	}

	err = part2(file)
	if err != nil {
		return err
	}

	return nil
}

func part1(file []string) error {
	if len(file) == 0 {
		return nil
	}

	costs := map[string]int{
		"noop": 1,
		"addx": 2,
	}

	x := 1
	cycle := 0
	signalStrength := 0

	for _, line := range file {
		if strings.TrimSpace(line) == "" {
			continue
		}
		instruction := strings.Split(line, " ")
		operator := instruction[0]
		for i := 0; i < costs[operator]; i++ {
			cycle++
			if cycle > 220 {
				continue
			}
			if cycle == 20 || (cycle > 20 && ((cycle-20)%40 == 0)) {
				signalStrength += x * cycle
			}
		}
		switch operator {
		case "addx":
			{
				operand, _ := strconv.Atoi(instruction[1])
				x += operand
			}
		}
	}

	fmt.Println(signalStrength)

	return nil
}

func part2(file []string) error {
	if len(file) == 0 {
		return nil
	}

	costs := map[string]int{
		"noop": 1,
		"addx": 2,
	}

	x := 1
	cycle := 0

	screen := [][]string{}
	screenLine := make([]string, 40)
	screenI := 0

	for _, line := range file {
		if strings.TrimSpace(line) == "" {
			continue
		}
		instruction := strings.Split(line, " ")
		operator := instruction[0]
		for i := 0; i < costs[operator]; i++ {
			cycle++

			screenLine[screenI] = "."
			if x == screenI || x-1 == screenI || x+1 == screenI {
				screenLine[screenI] = "#"
			}
			screenI++

			if cycle%40 == 0 {
				screen = append(screen, screenLine)
				fmt.Println(screenLine)
				screenLine = make([]string, 40)
				screenI = 0
			}
		}
		switch operator {
		case "addx":
			{
				operand, _ := strconv.Atoi(instruction[1])
				x += operand
			}
		}
	}

	return nil
}
