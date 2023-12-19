package tools

func TransposingStringsSlice(in []string) []string {
	numRows := len(in)
	numCols := len(in[0])

	transposedPatterns := make([]string, numCols)
	for i := 0; i < numCols; i++ {
		transposedPatterns[i] = ""
		for j := 0; j < numRows; j++ {
			transposedPatterns[i] += string(in[j][i])
		}
	}

	return transposedPatterns
}

func ReverseStringsSlice(in []string) []string {
	reversed := make([]string, len(in))
	for i, j := len(in)-1, 0; i >= 0; i, j = i-1, j+1 {
		reversed[j] = in[i]
	}
	return reversed
}
