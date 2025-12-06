package tools

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Integer | Float
}

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
