package tools

import "strconv"

func MustAtoi(in string) int {
	value, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}

	return value
}

func AtoiSlice(in []string) []int {
	out := make([]int, len(in))

	for i := range in {
		out[i] = MustAtoi(in[i])
	}

	return out
}
