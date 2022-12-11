package main

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

type Gorilla struct {
	ID              int
	Items           []*big.Int
	InspectionCount int64
	Operation       func(*big.Int) *big.Int
	Test            func(*big.Int) int
}

func part2(file []string) error {
	if len(file) == 0 {
		return nil
	}

	if file[len(file)-1] != "" {
		file = append(file, "")
	}

	monkeys := map[int]*Gorilla{}

	minimumCommonDivision := big.NewInt(1)

	for i := 0; i < len(file); i += 7 {
		monkeyRaw := file[i : i+7]

		id, _ := strconv.Atoi(strings.Split(strings.Split(monkeyRaw[0], ":")[0], " ")[1])

		itemsRaw := strings.Split(strings.Join(strings.Split(strings.TrimSpace(monkeyRaw[1]), " ")[2:], ""), ",")
		items := []*big.Int{}
		for _, s := range itemsRaw {
			item, _ := strconv.ParseInt(s, 10, 64)
			items = append(items, big.NewInt(item))
		}

		operationFunc := func(line string) func(old *big.Int) *big.Int {
			operationRaw := strings.Split(strings.TrimSpace(line), " ")[1:]

			leftOperand := operationRaw[2]
			leftOperandInt, _ := strconv.ParseInt(leftOperand, 10, 64)
			leftOperandBigInt := big.NewInt(leftOperandInt)
			rightOperand := operationRaw[4]
			rightOperandInt, _ := strconv.ParseInt(rightOperand, 10, 64)
			rightOperandBigInt := big.NewInt(rightOperandInt)
			operator := operationRaw[3]

			return func(old *big.Int) *big.Int {
				left := leftOperandBigInt
				if leftOperand == "old" {
					left = old
				}
				right := rightOperandBigInt
				if rightOperand == "old" {
					right = old
				}

				result := big.NewInt(0)

				switch operator {
				case "*":
					result = result.Mul(left, right)
				case "+":
					result = result.Add(left, right)
				case "-":
					result = result.Sub(left, right)
				case "/":
					result = result.Div(left, right)
				}

				return result
			}
		}(monkeyRaw[2])

		testFunc := func(testRaw, ifTrueRaw, ifFalseRaw string) func(int2 *big.Int) int {
			testDivisibleBy, _ := strconv.ParseInt(strings.Split(strings.TrimSpace(testRaw), " ")[3], 10, 64)
			testDivisibleByBig := big.NewInt(testDivisibleBy)
			ifTrueMonkeyID, _ := strconv.Atoi(strings.Split(strings.TrimSpace(ifTrueRaw), " ")[5])
			ifFalseMonkeyID, _ := strconv.Atoi(strings.Split(strings.TrimSpace(ifFalseRaw), " ")[5])

			minimumCommonDivision = minimumCommonDivision.Mul(minimumCommonDivision, testDivisibleByBig)

			return func(value *big.Int) int {
				if big.NewInt(0).Mod(value, testDivisibleByBig).Cmp(big.NewInt(0)) == 0 {
					return ifTrueMonkeyID
				}

				return ifFalseMonkeyID
			}
		}(monkeyRaw[3], monkeyRaw[4], monkeyRaw[5])

		monkey := &Gorilla{
			ID:        id,
			Items:     items,
			Operation: operationFunc,
			Test:      testFunc,
		}

		monkeys[monkey.ID] = monkey
	}

	rounds := 10000

	/*
		roundsToPrint := map[int]bool{
			1:  true,
			20: true,
		}

	*/

	for round := 1; round <= rounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := monkeys[i]
			for _, item := range monkey.Items {
				newItem := monkey.Operation(item)
				newItem = newItem.Mod(newItem, minimumCommonDivision)
				targetMonkeyID := monkey.Test(newItem)
				monkeys[targetMonkeyID].Items = append(monkeys[targetMonkeyID].Items, newItem)
				monkey.InspectionCount++
			}
			monkey.Items = []*big.Int{}
		}
		/*
			if _, exist := roundsToPrint[round]; exist || round%1000 == 0 {
				fmt.Printf("== After round %d ==\n", round)
				for i := 0; i < len(monkeys); i++ {
					monkey := monkeys[i]
					fmt.Printf("Monkey %d inspected items %d times.\n", monkey.ID, monkey.InspectionCount)
				}
				fmt.Println()
			}

		*/
	}

	inspections := []int64{}

	for i := 0; i < len(monkeys); i++ {
		monkey := monkeys[i]
		inspections = append(inspections, monkey.InspectionCount)
	}

	sort.Slice(inspections, func(i, j int) bool { return inspections[i] < inspections[j] })
	top := inspections[len(inspections)-2:]

	monkeyBusiness := int64(1)
	for _, val := range top {
		monkeyBusiness *= val
	}

	fmt.Println(monkeyBusiness)

	return nil
}
