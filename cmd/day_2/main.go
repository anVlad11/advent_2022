package main

import (
	"fmt"
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
	file, err := helpers.ReadFile("./data/input/day_2_input.txt")
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

/**
A - X - Rock
B - Y - Paper
C - Z - Scissors
*/

func part1(file []string) error {
	if len(file) == 0 {
		return nil
	}

	equals := map[byte]byte{'A': 'X', 'B': 'Y', 'C': 'Z'}
	// A is beaten by Y, B is beaten by Z, C is beaten by X
	isBeatenBy := map[byte]byte{'A': 'Y', 'B': 'Z', 'C': 'X'}
	cost := map[byte]int{'X': 1, 'Y': 2, 'Z': 3}

	score := 0
	for _, line := range file {
		if strings.TrimSpace(line) == "" {
			continue
		}
		roundScore := cost[line[2]]
		if equals[line[0]] == line[2] {
			roundScore += 3
		} else if isBeatenBy[line[0]] == line[2] {
			roundScore += 6
		}
		score += roundScore
	}

	fmt.Println(score)

	return nil
}

func part2(file []string) error {

	equals := map[byte]byte{'A': 'X', 'B': 'Y', 'C': 'Z'}
	// A is beaten by Y, B is beaten by Z, C is beaten by X
	isBeatenBy := map[byte]byte{'A': 'Y', 'B': 'Z', 'C': 'X'}
	// A beats Z, B beats X, C beats Y
	beats := map[byte]byte{'A': 'Z', 'B': 'X', 'C': 'Y'}

	letterCost := map[byte]int{'X': 1, 'Y': 2, 'Z': 3}
	resultCost := map[byte]int{'X': 0, 'Y': 3, 'Z': 6}

	solution := map[byte]map[byte]byte{
		'X': beats,
		'Y': equals,
		'Z': isBeatenBy,
	}

	score := 0
	for _, line := range file {
		if strings.TrimSpace(line) == "" {
			continue
		}

		roundScore := letterCost[solution[line[2]][line[0]]]
		roundScore += resultCost[line[2]]

		score += roundScore
	}

	fmt.Println(score)

	return nil
}
