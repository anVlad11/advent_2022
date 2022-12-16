package helpers

import "strconv"

func Atoi(in string) int {
	val, _ := strconv.Atoi(in)

	return val
}

func ParseInt(in string) int64 {
	val, _ := strconv.ParseInt(in, 10, 64)

	return val
}
