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
	file, err := helpers.ReadFile("./data/input/day_1_input.txt")
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
	var err error

	if len(file) == 0 {
		return nil
	}

	max := int64(0)
	current := int64(0)
	for _, line := range file {
		if strings.TrimSpace(line) == "" {
			if current > max {
				max = current
			}
			current = 0
			continue
		}
		var intval int64
		intval, err = strconv.ParseInt(line, 10, 64)
		if err != nil {
			return err
		}
		current += intval
	}
	fmt.Println(max)

	return nil
}

func part2(file []string) error {
	var err error

	if len(file) == 0 {
		return nil
	}
	top := []int64{}
	cur := int64(0)
	for lineNum, line := range file {
		if strings.TrimSpace(line) != "" {
			var intval int64
			intval, err = strconv.ParseInt(line, 10, 64)
			if err != nil {
				return err
			}
			cur += intval
		}

		if strings.TrimSpace(line) == "" || lineNum == len(file)-1 {
			if len(top) == 0 {
				top = append(top, cur)
			} else {
				added := false
				for i, val := range top {
					if cur > val {
						left := []int64{}
						if i > 0 {
							left = top[0:i]
						}
						right := []int64{}
						if i+1 < len(top) {
							right = top[i:]
						}
						newTop := append(left, cur)
						newTop = append(newTop, right...)
						top = newTop
						added = true
						break
					}
				}
				if !added {
					top = append(top, cur)
				}
			}
			cur = 0

			continue
		}

	}
	topSum := int64(0)
	for i, item := range top {
		topSum += item
		if i == 2 {
			break
		}
	}

	fmt.Println(topSum)

	return nil
}
