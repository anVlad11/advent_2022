package main

import (
	"fmt"

	"github.com/anVlad11/advent_2022/internal/helpers"
)

func main() {
	err := do()
	if err != nil {
		fmt.Println(err)
	}
}

func do() error {
	file, err := helpers.ReadFile("./data/input/day_6_input.txt")
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

	markerPosition := 4

	for _, line := range file {
		i := markerPosition - 1
		for ; i < len(line); i++ {
			dict := map[uint8]bool{}
			for j := markerPosition - 1; j >= 0; j-- {
				dict[line[i-j]] = true
			}
			if len(dict) == markerPosition {
				fmt.Println(i + 1)
				break
			}
		}
	}

	return nil
}

func part2(file []string) error {
	if len(file) == 0 {
		return nil
	}

	markerPosition := 14

	for _, line := range file {
		i := markerPosition - 1
		for ; i < len(line); i++ {
			dict := map[uint8]bool{}
			for j := markerPosition - 1; j >= 0; j-- {
				dict[line[i-j]] = true
			}
			if len(dict) == markerPosition {
				fmt.Println(i + 1)
				break
			}
		}
	}

	return nil
}
