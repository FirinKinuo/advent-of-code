package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/FirinKinuo/advent-of-code"
)

type Day struct {
	*problem.DayTemplate
	wordSearch []string
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2024", "day04_ceres_search", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	d.wordSearch = strings.Split(input, "\n")
}

func (d *Day) rotate(matrix []string) []string {
	rotated := make([]string, len(matrix[0]))

	for i := 0; i < len(matrix[0]); i++ {
		var sb strings.Builder
		for j := len(matrix) - 1; j >= 0; j-- {
			sb.WriteByte(matrix[j][i])
		}
		rotated[i] = sb.String()
	}

	return rotated
}

func (d *Day) match(matrix []string, pattern string, gap int, width int) int {
	matchCount := 0

	for i := 0; i < len(matrix)-width+1; i++ {
		for j := 0; j < len(matrix[i])-width+1; j++ {
			block := ""
			for d := 0; d < width; d++ {
				block += matrix[i+d][j : j+width]
			}

			if d.checkPattern(block, pattern, gap) {
				matchCount++
			}
		}
	}
	return matchCount
}

func (d *Day) checkPattern(block string, pattern string, gap int) bool {
	if len(block) < len(pattern) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		if pattern[i] != block[i*gap+i] {
			return false
		}
	}
	return true
}

func (d *Day) FirstProblem() int {
	var result int

	for rotation := 0; rotation < 4; rotation++ {
		for _, row := range d.wordSearch {
			result += strings.Count(row, "XMAS")
		}
		result += d.match(d.wordSearch, "XMAS", 4, 4)
		d.wordSearch = d.rotate(d.wordSearch)
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	for rotation := 0; rotation < 4; rotation++ {
		result += d.match(d.wordSearch, `MMASS`, 1, 3)
		d.wordSearch = d.rotate(d.wordSearch)
	}

	return result
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
