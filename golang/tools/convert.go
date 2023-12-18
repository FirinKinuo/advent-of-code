package tools

import "strconv"

func MustAtoi(in string) int {
	value, _ := strconv.Atoi(in)

	return value
}
