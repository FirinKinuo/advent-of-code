package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

type Day struct {
	*problem.DayTemplate

	freshIDRanges [][2]int
	available     []int
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2025", "day05_cafeteria", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	lists := tools.SplitSeparatedLines(input)

	for _, freshIDRange := range tools.SplitNewLines(lists[0]) {
		edges := strings.Split(freshIDRange, "-")
		d.freshIDRanges = append(d.freshIDRanges, [2]int{tools.MustAtoi(edges[0]), tools.MustAtoi(edges[1])})
	}

	d.available = tools.AtoiSlice(tools.SplitNewLines(lists[1]))
}

func (d *Day) mergeRanges(ranges [][2]int) [][2]int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	mergedRanges := make([][2]int, 0, len(ranges)/2)

	mergedRanges = append(mergedRanges, ranges[0])

	for i := 1; i < len(ranges); i++ {
		last := &mergedRanges[len(mergedRanges)-1]
		current := ranges[i]

		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			mergedRanges = append(mergedRanges, current)
		}
	}

	return mergedRanges
}

func (d *Day) FirstProblem() int {
	var result int

	mergedRanges := d.mergeRanges(d.freshIDRanges)

	for _, availableID := range d.available {
		for _, idRange := range mergedRanges {
			if availableID >= idRange[0] && availableID <= idRange[1] {
				result++
				break
			}
		}
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	mergedRanges := d.mergeRanges(d.freshIDRanges)

	for _, idRange := range mergedRanges {
		result += idRange[1] - idRange[0] + 1
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
