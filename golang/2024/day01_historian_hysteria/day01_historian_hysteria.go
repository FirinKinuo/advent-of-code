package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

type locationIdsList struct {
	left  []int
	right []int
}

type Day struct {
	*problem.DayTemplate

	locationIds locationIdsList
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2024", "day01_historian_hysteria", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	locationIdsRows := tools.SplitNewLines(input)

	d.locationIds.left = make([]int, 0, len(locationIdsRows))
	d.locationIds.right = make([]int, 0, len(locationIdsRows))

	for _, row := range locationIdsRows {
		locationIds := strings.Split(row, "   ")
		d.locationIds.left = append(d.locationIds.left, tools.MustAtoi(locationIds[0]))
		d.locationIds.right = append(d.locationIds.right, tools.MustAtoi(locationIds[1]))
	}
}

func (d *Day) FirstProblem() int {
	var result int

	slices.Sort(d.locationIds.left)
	slices.Sort(d.locationIds.right)

	for i := 0; i < len(d.locationIds.left); i++ {
		result += tools.Abs(d.locationIds.left[i] - d.locationIds.right[i])
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	for _, id := range d.locationIds.left {
		countInRight := tools.CountInSlice(d.locationIds.right, id)

		result += id * countInRight
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
