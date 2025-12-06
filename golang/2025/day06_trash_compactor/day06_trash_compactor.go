package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

type expression struct {
	operands []int
	operator string
}

type Day struct {
	*problem.DayTemplate

	worksheet []string
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2025", "day06_trash_compactor", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	lines := tools.SplitNewLines(input)
	linesLen := make([]int, 0, len(lines))

	for _, line := range lines {
		linesLen = append(linesLen, len(line))
	}

	maxLen := slices.Max(linesLen)

	for i := 0; i < len(lines); i++ {
		lines[i] = lines[i] + strings.Repeat(" ", maxLen-len(lines[i])+1)
	}

	d.worksheet = lines
}

func (d *Day) FirstProblem() int {
	var result int

	worksheetLen := len(d.worksheet)
	operandMatrix := make([][]int, worksheetLen-1)

	for rowIdx, row := range d.worksheet[:worksheetLen-1] {
		var operandRunes []rune

		for _, sym := range row {
			if sym == ' ' && len(operandRunes) == 0 {
				continue
			} else if sym == ' ' {
				operandMatrix[rowIdx] = append(operandMatrix[rowIdx], tools.MustAtoi(string(operandRunes)))
				operandRunes = []rune{}
			} else {
				operandRunes = append(operandRunes, sym)
			}
		}
	}

	operandMatrix = tools.Transpose(operandMatrix)

	operatorCount := 0
	for _, sym := range d.worksheet[worksheetLen-1] {
		switch sym {
		case '+':
			result += tools.Sum(operandMatrix[operatorCount]...)
		case '*':
			result += tools.Multiply(operandMatrix[operatorCount]...)
		default:
			continue
		}

		operatorCount++
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	worksheetLen := len(d.worksheet)
	operandMatrix := make([][]int, len(d.worksheet[0])/3)

	rowIdx := 0

	for i := len(d.worksheet[0]) - 2; i >= 0; i-- {
		var operandBytes []byte

		for j := 0; j < worksheetLen-1; j++ {
			if d.worksheet[j][i] == ' ' && len(operandBytes) == 0 {
				continue
			} else if d.worksheet[j][i] != ' ' {
				operandBytes = append(operandBytes, d.worksheet[j][i])
			}
		}

		if len(operandBytes) == 0 {
			rowIdx++
		} else {
			operandMatrix[rowIdx] = append(operandMatrix[rowIdx], tools.MustAtoi(string(operandBytes)))
		}
	}

	operatorCount := 0
	for i := len(d.worksheet[0]) - 1; i >= 0; i-- {
		switch d.worksheet[worksheetLen-1][i] {
		case '+':
			result += tools.Sum(operandMatrix[operatorCount]...)
		case '*':
			result += tools.Multiply(operandMatrix[operatorCount]...)
		default:
			continue
		}

		operatorCount++
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
