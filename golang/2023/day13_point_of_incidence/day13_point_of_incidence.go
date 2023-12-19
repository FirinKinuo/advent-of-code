package main

import (
	"fmt"
	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
	"log"
)

type note struct {
	patterns []string
}

type Day struct {
	*problem.DayTemplate
	notes []note
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2023", "day13_point_of_incidence", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %s", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	for _, s := range tools.SplitSeparatedLines(input) {
		d.notes = append(d.notes, note{patterns: tools.SplitNewLines(s)})
	}
}

func (d *Day) reflect(pattern []string, smudge int) int {
	for i := 1; i < len(pattern); i++ {
		a := tools.ReverseStringsSlice(pattern[:i])
		b := pattern[i:]
		trim := min(len(a), len(b))
		if d.countMismatch(a[:trim], b[:trim]) == smudge {
			return i
		}
	}
	return 0
}

func (d *Day) countMismatch(a, b []string) int {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count
}

func (d *Day) FirstProblem() int {
	reflections := 0

	for _, n := range d.notes {
		reflections += 100*d.reflect(n.patterns, 0) + d.reflect(tools.TransposingStringsSlice(n.patterns), 0)
	}

	return reflections
}

func (d *Day) SecondProblem() int {
	reflections := 0

	for _, n := range d.notes {
		reflections += 100*d.reflect(n.patterns, 1) + d.reflect(tools.TransposingStringsSlice(n.patterns), 1)
	}

	return reflections
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
