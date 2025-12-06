package tools

func FindLCM[T Integer](a, b T) T {
	return a * b / FindGCD(a, b)
}

func FindLCMOfSlice[T Integer](numbers []T) T {
	if len(numbers) < 2 {
		panic("slice must have at least 2 elements")
	}

	lcm := numbers[0]
	for i := 1; i < len(numbers); i++ {
		lcm = FindLCM(lcm, numbers[i])
	}

	return lcm
}

func FindGCD[T Integer](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
