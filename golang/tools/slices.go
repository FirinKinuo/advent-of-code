package tools

import (
	"cmp"
	"image"
)

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

func CountInSlice[T comparable](slice []T, value T) int {
	count := 0
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}

func IsMonotonic[T cmp.Ordered](nums []T) (index int, ok bool) {
	increasing, decreasing := false, false

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			increasing = true
		} else if nums[i] < nums[i-1] {
			decreasing = true
		}

		if increasing && decreasing {
			index = i
			return index, false
		}
	}

	return -1, true
}

func Copy2DSlice[T any](src [][]T) [][]T {
	dest := make([][]T, len(src))

	for i := range src {
		dest[i] = make([]T, len(src[i]))
		copy(dest[i], src[i])
	}

	return dest
}

func PointFromSlice(nums []int) image.Point {
	return image.Point{X: nums[0], Y: nums[1]}
}
