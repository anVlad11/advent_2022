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
	file, err := helpers.ReadFile("./data/input/day_8_input.txt")
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

	forest := [][]int{}
	forestRotated := [][]int{}

	for i, line := range file {
		if strings.TrimSpace(line) == "" {
			continue
		}
		forest = append(forest, []int{})
		for j, v := range line {
			val, _ := strconv.Atoi(string(v))
			if len(forestRotated) <= j {
				forestRotated = append(forestRotated, []int{})
			}
			forestRotated[j] = append(forestRotated[j], val)
			forest[i] = append(forest[i], val)
		}
	}

	visible := func(side []int, target int) bool {
		if len(side) == 0 {
			return true
		}

		for _, val := range side {
			if val >= target {
				return false
			}
		}

		return true
	}

	visibleTrees := 2*len(forest) + 2*(len(forest[0])-2)

	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest[i])-1; j++ {
			cur := forest[i][j]
			left := forest[i][0:j]
			right := forest[i][j+1:]
			top := forestRotated[j][0:i]
			down := forestRotated[j][i+1:]

			if visible(left, cur) || visible(right, cur) || visible(top, cur) || visible(down, cur) {
				visibleTrees++
			}
		}
	}

	fmt.Println(visibleTrees)

	return nil
}

func part2(file []string) error {
	if len(file) == 0 {
		return nil
	}

	forest := [][]int{}
	forestRotated := [][]int{}

	for i, line := range file {
		if strings.TrimSpace(line) == "" {
			continue
		}
		forest = append(forest, []int{})
		for j, v := range line {
			val, _ := strconv.Atoi(string(v))
			if len(forestRotated) <= j {
				forestRotated = append(forestRotated, []int{})
			}
			forestRotated[j] = append(forestRotated[j], val)
			forest[i] = append(forest[i], val)
		}
	}

	maxScenicValue := 0

	iterate := func(side []int, max int, reversed bool) int {
		if len(side) == 0 {
			return 0
		}

		start := 0
		iterator := 1
		if reversed {
			start = len(side) - 1
			iterator = -1
		}

		sum := 0
		for k := start; k >= 0 && k < len(side); k += iterator {
			sum++
			if side[k] >= max {
				return sum
			}
		}

		return sum
	}

	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest[i])-1; j++ {
			cur := forest[i][j]
			left := forest[i][0:j]
			right := forest[i][j+1:]
			top := forestRotated[j][0:i]
			down := forestRotated[j][i+1:]

			leftScenicValue := iterate(left, cur, true)
			rightScenicValue := iterate(right, cur, false)
			topScenicValue := iterate(top, cur, true)
			downScenicValue := iterate(down, cur, false)

			scenicValue := leftScenicValue * rightScenicValue * topScenicValue * downScenicValue
			if scenicValue > maxScenicValue {
				maxScenicValue = scenicValue
			}

		}
	}

	fmt.Println(maxScenicValue)

	return nil
}
