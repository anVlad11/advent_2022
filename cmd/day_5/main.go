package main

import (
	"fmt"
	"regexp"
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
	file, err := helpers.ReadFile("./data/input/day_5_input.txt")
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
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
*/

func part1(file []string) error {
	if len(file) == 0 {
		return nil
	}

	stacks := map[int][]string{}

	i := 0
	for ; i < len(file); i++ {
		line := file[i]
		if strings.HasPrefix(line, " 1") {
			i = i + 2
			break
		}

		column := 0
		for j := 1; j < len(line); j = j + 4 {
			column++
			val := string(line[j])
			if val == " " {
				continue
			}

			if _, exists := stacks[column]; !exists {
				stacks[column] = []string{val}
			} else {
				stacks[column] = append(stacks[column], val)
			}
		}
	}

	moveRegexp, err := regexp.Compile(`move (\d*) from (\d*) to (\d*)`)
	if err != nil {
		return err
	}

	for ; i < len(file); i++ {
		line := file[i]

		values := moveRegexp.FindAllStringSubmatch(line, -1)
		if len(values) == 0 || len(values[0]) != 4 {
			continue
		}

		var amount int
		amount, err = strconv.Atoi(values[0][1])
		if err != nil {
			return err
		}

		var from int
		from, err = strconv.Atoi(values[0][2])
		if err != nil {
			return err
		}

		var to int
		to, err = strconv.Atoi(values[0][3])
		if err != nil {
			return err
		}

		if _, exists := stacks[to]; !exists {
			stacks[to] = []string{}
		}

		for amount > 0 {
			amount--

			if len(stacks[from]) == 0 {
				continue
			}

			if len(stacks[to]) == 0 {
				stacks[to] = []string{stacks[from][0]}
			} else {
				stacks[to] = append([]string{stacks[from][0]}, stacks[to]...)
			}

			if len(stacks[from]) > 1 {
				stacks[from] = stacks[from][1:len(stacks[from])]
			} else {
				stacks[from] = []string{}
			}
		}
	}
	for column := 1; ; column++ {
		stack, exists := stacks[column]
		if !exists {
			break
		}
		if len(stack) == 0 {
			continue
		}
		fmt.Print(stack[0])
	}
	fmt.Println()

	return nil
}

func part2(file []string) error {
	if len(file) == 0 {
		return nil
	}

	stacks := map[int][]string{}

	i := 0
	for ; i < len(file); i++ {
		line := file[i]
		if strings.HasPrefix(line, " 1") {
			i = i + 2
			break
		}

		column := 0
		for j := 1; j < len(line); j = j + 4 {
			column++
			val := string(line[j])
			if val == " " {
				continue
			}

			if _, exists := stacks[column]; !exists {
				stacks[column] = []string{val}
			} else {
				stacks[column] = append(stacks[column], val)
			}
		}
	}

	moveRegexp, err := regexp.Compile(`move (\d*) from (\d*) to (\d*)`)
	if err != nil {
		return err
	}

	for ; i < len(file); i++ {
		line := file[i]

		values := moveRegexp.FindAllStringSubmatch(line, -1)
		if len(values) == 0 || len(values[0]) != 4 {
			continue
		}

		amount, _ := strconv.Atoi(values[0][1])
		from, _ := strconv.Atoi(values[0][2])
		to, _ := strconv.Atoi(values[0][3])

		if _, exists := stacks[to]; !exists {
			stacks[to] = []string{}
		}

		leftPart := make([]string, len(stacks[from][0:amount]))
		for j, s := range stacks[from][0:amount] {
			leftPart[j] = s
		}

		rightPart := make([]string, len(stacks[to]))
		for j, s := range stacks[to] {
			rightPart[j] = s
		}

		// stacks[to] = append(stacks[from][0:amount], stacks[to]...) won't work
		stacks[to] = append(leftPart, rightPart...)
		if len(stacks[from]) == amount {
			stacks[from] = []string{}
		} else {
			stacks[from] = stacks[from][amount:]
		}
	}
	for column := 1; ; column++ {
		stack, exists := stacks[column]
		if !exists {
			break
		}
		if len(stack) == 0 {
			continue
		}
		fmt.Print(stack[0])
	}

	return nil
}
