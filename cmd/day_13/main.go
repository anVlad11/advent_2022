package main

import (
	"encoding/json"
	"fmt"
	"sort"
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
	file, err := helpers.ReadFile("./data/input/day_13_input.txt")
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

var logger = Logger{Level: -1, Enabled: false}

func part1(file []string) error {
	if len(file) == 0 {
		return nil
	}

	if strings.TrimSpace(file[len(file)-1]) != "" {
		file = append(file, "")
	}

	rightPairsSum := 0
	pair := 0
	for i := 0; i < len(file); i += 3 {
		pair++
		var left []interface{}
		var right []interface{}
		json.Unmarshal([]byte(file[i]), &left)
		json.Unmarshal([]byte(file[i+1]), &right)

		logger.Level = -1
		logger.Println(fmt.Sprintf("== Pair %d ==", pair))
		res := parse(left, right)
		if res == Lt {
			rightPairsSum += pair
		}
		logger.Println()
	}

	fmt.Println(rightPairsSum)

	return nil
}

func part2(file []string) error {
	if len(file) == 0 {
		return nil
	}

	packets := [][]interface{}{{[]interface{}{2.0}}, {[]interface{}{6.0}}}

	for i := 0; i < len(file); i += 1 {
		if strings.TrimSpace(file[i]) == "" {
			continue
		}
		var packet []interface{}
		json.Unmarshal([]byte(file[i]), &packet)
		packets = append(packets, packet)
	}

	sort.Slice(packets, func(i, j int) bool {
		return parse(packets[i], packets[j]) == Lt
	})

	first := 0
	last := 0

	for i, packet := range packets {
		if len(packet) == 1 {
			switch typed := packet[0].(type) {
			case []interface{}:
				{
					if len(typed) == 1 {
						switch typedItem := typed[0].(type) {
						case float64:
							{
								switch typedItem {
								case 2:
									first = i + 1
								case 6:
									last = i + 1
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(first * last)

	return nil
}
