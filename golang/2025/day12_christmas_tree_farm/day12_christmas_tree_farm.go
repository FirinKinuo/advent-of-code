package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

type Day struct {
	*problem.DayTemplate

	shapeSizes []int
	regions    []Region
}

type Region struct {
	width  int
	length int
	counts []int
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2025", "day12_christmas_tree_farm", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	parts := tools.SplitSeparatedLines(input)

	d.shapeSizes = make([]int, len(parts)-1)

	for i, part := range parts[:len(parts)-1] {
		d.shapeSizes[i] = strings.Count(part, "#")
	}

	for _, region := range tools.SplitNewLines(parts[len(parts)-1]) {
		regionParts := strings.Split(region, ": ")
		size := strings.Split(regionParts[0], "x")

		d.regions = append(d.regions, Region{
			width:  tools.MustAtoi(size[0]),
			length: tools.MustAtoi(size[1]),
			counts: tools.AtoiSlice(strings.Split(regionParts[1], " ")),
		})
	}
}

func (d *Day) FirstProblem() int {
	var result int

	for _, region := range d.regions {
		shapesArea := 0
		for i, count := range region.counts {
			shapesArea += count * d.shapeSizes[i]
		}

		if shapesArea <= region.width*region.length {
			result++
		}
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	return result
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
