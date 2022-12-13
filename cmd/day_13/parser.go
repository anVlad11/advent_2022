package main

import (
	"fmt"
)

type Sign string

const (
	Eq = Sign("eq")
	Gt = Sign("gt")
	Lt = Sign("lt")
)

func isFloat(input interface{}) bool {
	switch input.(type) {
	case float64:
		return true
	default:
		return false
	}
}

func isSlice(input interface{}) bool {
	switch input.(type) {
	case []interface{}:
		return true
	default:
		return false
	}
}

func toSlice(input interface{}) []interface{} {
	switch typed := input.(type) {
	case []interface{}:
		return typed
	default:
		return nil
	}
}

func toFloat(input interface{}) float64 {
	switch typed := input.(type) {
	case float64:
		return typed
	default:
		return 0
	}
}

func parse(left, right interface{}) Sign {
	logger.AddLevel()
	defer logger.SubLevel()

	logger.Println(fmt.Sprintf("- Compare %v vs %v", left, right))

	var leftSlice []interface{}
	var rightSlice []interface{}

	if isFloat(left) == isFloat(right) {
		if isFloat(left) {
			rightTyped := toFloat(right)
			leftTyped := toFloat(left)

			res := Eq
			if leftTyped < rightTyped {
				logger.AddLevel()
				logger.Println("- Left side is smaller, so inputs are in the right order")
				logger.SubLevel()
				res = Lt
			} else if leftTyped > rightTyped {
				logger.AddLevel()
				logger.Println("- Right side is smaller, so inputs are not in the right order")
				logger.SubLevel()
				res = Gt
			}

			return res

		}

		leftSlice = toSlice(left)
		rightSlice = toSlice(right)

	} else {
		logger.AddLevel()
		if isFloat(left) {
			leftSlice = []interface{}{toFloat(left)}
			rightSlice = toSlice(right)

			logger.Println(fmt.Sprintf("- Mixed types; convert left to %v and retry comparison", leftSlice))
			logger.Println(fmt.Sprintf("- Compare %v vs %v", leftSlice, rightSlice))

		} else {
			leftSlice = toSlice(left)
			rightSlice = []interface{}{toFloat(right)}

			logger.Println(fmt.Sprintf("- Mixed types; convert right to %v and retry comparison", rightSlice))
			logger.Println(fmt.Sprintf("- Compare %v vs %v", leftSlice, rightSlice))
		}
	}

	maxLen := len(leftSlice)
	if len(rightSlice) > maxLen {
		maxLen = len(rightSlice)
	}

	if maxLen == 0 {
		logger.AddLevel()
		logger.Println("- Both sides ran out of items, so i don't know what to do")
		return Eq
	}

	for i := 0; i < maxLen; i++ {
		if len(leftSlice) < i+1 {
			logger.AddLevel()
			logger.Println("- Left side ran out of items, so inputs are in the right order")
			return Lt
		}

		if len(rightSlice) < i+1 {
			logger.AddLevel()
			logger.Println("- Right side ran out of items, so inputs are not in the right order")
			return Gt
		}

		res := parse(leftSlice[i], rightSlice[i])

		if res != Eq {
			return res
		}

	}

	return Eq

}
