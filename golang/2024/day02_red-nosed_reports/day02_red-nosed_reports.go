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
	reports [][]int
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2024", "day02_red-nosed_reports", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	reports := strings.Split(input, "\n")

	d.reports = make([][]int, len(reports))

	for i, report := range reports {
		for _, levelString := range strings.Split(report, " ") {
			d.reports[i] = append(d.reports[i], tools.MustAtoi(levelString))
		}
	}
}

func (d *Day) FirstProblem() int {
	var safeReportCount int
	for _, report := range d.reports {
		if d.isSafe(report) {
			safeReportCount++
			continue
		}
	}

	return safeReportCount
}

func (d *Day) isSafe(report []int) bool {
	_, isMonotonic := tools.IsMonotonic(report)
	if !isMonotonic {
		return false
	}

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]

		if tools.Abs(diff) > 3 || diff == 0 {
			return false
		}
	}

	return true
}

func (d *Day) SecondProblem() int {
	var safeReportCount int
	for _, report := range d.reports {
		if d.isSafe(report) {
			safeReportCount++
			continue
		}

		for i := 0; i < len(report); i++ {
			tryReport := make([]int, 0, len(report)-1)
			tryReport = append(tryReport, report[:i]...)
			tryReport = append(tryReport, report[i+1:]...)

			if d.isSafe(tryReport) {
				safeReportCount++
				break
			}
		}
	}

	return safeReportCount
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
