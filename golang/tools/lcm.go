package tools

func FindLCM(a, b int) int {
	return a * b / FindGCD(a, b)
}

func FindLCMOfSlice(numbers []int) int {
	if len(numbers) < 2 {
		panic("FindLCMOfSlice: slice must have at least 2 elements")
	}

	lcm := numbers[0]
	for i := 1; i < len(numbers); i++ {
		lcm = FindLCM(lcm, numbers[i])
	}

	return lcm
}

func FindGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
