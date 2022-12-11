package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	ID              int
	Items           []int64
	InspectionCount int64
	Operation       func(int64) int64
	Test            func(int64) int
}

func part1(file []string) error {
	if len(file) == 0 {
		return nil
	}

	if file[len(file)-1] != "" {
		file = append(file, "")
	}

	monkeys := map[int]*Monkey{}

	for i := 0; i < len(file); i += 7 {
		monkeyRaw := file[i : i+7]

		id, _ := strconv.Atoi(strings.Split(strings.Split(monkeyRaw[0], ":")[0], " ")[1])

		itemsRaw := strings.Split(strings.Join(strings.Split(strings.TrimSpace(monkeyRaw[1]), " ")[2:], ""), ",")
		items := make([]int64, 0, len(itemsRaw))
		for _, s := range itemsRaw {
			item, _ := strconv.ParseInt(s, 10, 64)
			items = append(items, item)
		}

		operationFunc := func(line string) func(old int64) int64 {
			operationRaw := strings.Split(strings.TrimSpace(line), " ")[1:]

			leftOperand := operationRaw[2]
			leftOperandInt, _ := strconv.ParseInt(leftOperand, 10, 64)
			rightOperand := operationRaw[4]
			rightOperandInt, _ := strconv.ParseInt(rightOperand, 10, 64)
			operator := operationRaw[3]

			return func(old int64) int64 {
				left := leftOperandInt
				if leftOperand == "old" {
					left = old
				}
				right := rightOperandInt
				if rightOperand == "old" {
					right = old
				}

				result := int64(0)

				switch operator {
				case "*":
					result = left * right
				case "+":
					result = left + right
				case "-":
					result = left - right
				case "/":
					result = int64(math.Round(float64(left) / float64(right)))
				}

				return result
			}
		}(monkeyRaw[2])

		testFunc := func(testRaw, ifTrueRaw, ifFalseRaw string) func(int64) int {
			testDivisibleBy, _ := strconv.ParseInt(strings.Split(strings.TrimSpace(testRaw), " ")[3], 10, 64)
			ifTrueMonkeyID, _ := strconv.Atoi(strings.Split(strings.TrimSpace(ifTrueRaw), " ")[5])
			ifFalseMonkeyID, _ := strconv.Atoi(strings.Split(strings.TrimSpace(ifFalseRaw), " ")[5])

			return func(value int64) int {
				if value%testDivisibleBy == 0 {
					return ifTrueMonkeyID
				}

				return ifFalseMonkeyID
			}
		}(monkeyRaw[3], monkeyRaw[4], monkeyRaw[5])

		monkey := &Monkey{
			ID:        id,
			Items:     items,
			Operation: operationFunc,
			Test:      testFunc,
		}

		monkeys[monkey.ID] = monkey
	}

	rounds := 20
	for round := 0; round < rounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := monkeys[i]
			for _, item := range monkey.Items {
				newItem := monkey.Operation(item)
				newItem = newItem / 3
				targetMonkeyID := monkey.Test(newItem)
				monkeys[targetMonkeyID].Items = append(monkeys[targetMonkeyID].Items, newItem)
				monkey.InspectionCount++
			}
			monkey.Items = []int64{}
		}
	}

	inspections := []int64{}

	for i := 0; i < len(monkeys); i++ {
		monkey := monkeys[i]
		inspections = append(inspections, monkey.InspectionCount)
		//fmt.Printf("Monkey %d inspected items %d times.\n", monkey.ID, monkey.InspectionCount)
	}

	sort.Slice(inspections, func(i, j int) bool { return inspections[i] < inspections[j] })
	top := inspections[len(inspections)-2:]

	monkeyBusiness := int64(0)
	for _, val := range top {
		if monkeyBusiness == 0 {
			monkeyBusiness += val
		} else {
			monkeyBusiness *= val
		}
	}

	fmt.Println(monkeyBusiness)
	//fmt.Println()

	return nil
}
