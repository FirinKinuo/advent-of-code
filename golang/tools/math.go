package tools

func Abs[T Signed | Float](x T) T {
	var zero T
	if x < zero {
		return -x
	}

	return x
}

func Sign[T Signed | Float](x T) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

func Sum[T Number](operand ...T) T {
	var result T
	for _, op := range operand {
		result += op
	}

	return result
}

func Multiply[T Number](operand ...T) T {
	result := T(1)
	for _, op := range operand {
		result *= op
	}

	return result
}
