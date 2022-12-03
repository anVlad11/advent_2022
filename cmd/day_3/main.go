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
	file, err := helpers.ReadFile("./data/input/day_3_input.txt")
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
Letter	ASCII	Weird
A		65		27
B		66		28
....
Z		90		52
...
a		97		1
b		98		2
...
z		122		26
*/

func convert(letter int32) int32 {
	if letter >= 'A' && letter <= 'Z' {
		return letter - 38
	}

	return letter - 96
}

func part1(file []string) error {
	if len(file) == 0 {
		return nil
	}

	result := int32(0)

	for _, line := range file {
		if strings.TrimSpace(line) == "" {
			continue
		}

		if len(line)%2 > 0 {
			return fmt.Errorf("this line is odd: %s", line)
		}

		letters := map[rune]bool{}

		leftPart := line[0 : len(line)/2]
		for _, letter := range leftPart {
			letters[letter] = false
		}

		rightPart := line[len(line)/2:]
		for _, letter := range rightPart {
			if checked, exists := letters[letter]; exists {
				if checked {
					continue
				}
				letters[letter] = true
				result += convert(letter)
			}
		}
	}

	fmt.Println(result)

	return nil
}

func part2(file []string) error {
	if len(file) == 0 {
		return nil
	}

	result := int32(0)

	groupSize := 3

	groupLetters := map[int32]map[int]bool{}

	for i, line := range file {
		if i%groupSize == 0 {
			for letter, value := range groupLetters {
				if len(value) == groupSize {
					result += convert(letter)
				}
			}
			groupLetters = map[int32]map[int]bool{}
		}

		if strings.TrimSpace(line) == "" {
			continue
		}

		if len(line)%2 > 0 {
			return fmt.Errorf("this line is odd: %s", line)
		}

		for _, letter := range line {
			if _, exists := groupLetters[letter]; exists {
				groupLetters[letter][i%3] = true
			} else {
				groupLetters[letter] = map[int]bool{i % 3: true}
			}

		}

	}

	for letter, value := range groupLetters {
		if len(value) == groupSize {
			result += convert(letter)
		}
	}

	fmt.Println(result)

	return nil
}
