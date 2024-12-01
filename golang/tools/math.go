package tools

func Abs[T int | int32 | int64 | float32 | float64](x T) T {
	var zero T
	if x < zero {
		return -x
	}

	return x
}
