package main

import (
	"fmt"
	"log"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

const (
	maxAdjacent = 4
)

type Day struct {
	*problem.DayTemplate
	grid [][]rune
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2025", "day04_printing_department", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	lines := tools.SplitNewLines(input)
	d.grid = make([][]rune, len(lines))

	for i, line := range lines {
		d.grid[i] = []rune(line)
	}
}

func (d *Day) findAdjacentRolls(grid [][]rune, lineIndex, rowIndex int) int {
	rolls := 0

	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			neighborLine := lineIndex + y
			neighborRow := rowIndex + x
			if x == 0 && y == 0 || neighborLine < 0 || neighborLine > len(grid)-1 || neighborRow < 0 || neighborRow > len(grid[0])-1 {
				continue
			}
			if grid[neighborLine][neighborRow] == '@' {
				rolls++
			}

			if rolls == maxAdjacent {
				break
			}
		}
	}

	return rolls
}

func (d *Day) FirstProblem() int {
	var result int

	for lineIndex, line := range d.grid {
		for rowIndex, row := range line {
			if row != '@' {
				continue
			}
			rolls := d.findAdjacentRolls(d.grid, lineIndex, rowIndex)

			if rolls < maxAdjacent {
				result++
			}
		}
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	for {
		removedRolls := 0
		for lineIndex, line := range d.grid {
			for rowIndex, row := range line {
				if row != '@' {
					continue
				}

				rolls := d.findAdjacentRolls(d.grid, lineIndex, rowIndex)

				if rolls < maxAdjacent {
					d.grid[lineIndex][rowIndex] = 'x'
					removedRolls++
					result++
				}
			}
		}

		if removedRolls == 0 {
			break
		}
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
