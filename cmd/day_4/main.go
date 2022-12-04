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
	file, err := helpers.ReadFile("./data/input/day_4_input.txt")
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

	result := 0

	for _, line := range file {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		pairRaw := strings.Split(line, ",")
		pair := make([][]int64, len(pairRaw))
		for i, s := range pairRaw {
			pair[i] = []int64{}
			values := strings.Split(s, "-")
			for _, valueRaw := range values {
				value, err := strconv.ParseInt(valueRaw, 10, 64)
				if err != nil {
					return err
				}
				pair[i] = append(pair[i], value)
			}
		}

		if (pair[0][0] >= pair[1][0] && pair[0][1] <= pair[1][1]) ||
			(pair[1][0] >= pair[0][0] && pair[1][1] <= pair[0][1]) {
			result++
		}
	}

	fmt.Println(result)

	return nil
}

func part2(file []string) error {
	if len(file) == 0 {
		return nil
	}

	result := 0

	for _, line := range file {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		pairRaw := strings.Split(line, ",")
		pair := make([][]int64, len(pairRaw))
		for i, s := range pairRaw {
			pair[i] = []int64{}
			values := strings.Split(s, "-")
			for _, valueRaw := range values {
				value, err := strconv.ParseInt(valueRaw, 10, 64)
				if err != nil {
					return err
				}
				pair[i] = append(pair[i], value)
			}
		}

		if (pair[0][0] <= pair[1][1]) && (pair[0][1] >= pair[1][0]) {
			result++
		}
	}

	fmt.Println(result)

	return nil
}
