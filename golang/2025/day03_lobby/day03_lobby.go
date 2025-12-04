package main

import (
	"fmt"
	"log"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

type Day struct {
	*problem.DayTemplate

	banks []string
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2025", "day03_lobby", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	d.banks = tools.SplitNewLines(input)
}

func findMaxJoltage(bank string, batteriesCount int) int {
	bankLen := len(bank)

	stack := make([]byte, 0, batteriesCount)
	toRemove := bankLen - batteriesCount

	for i := 0; i < bankLen; i++ {
		for len(stack) > 0 && toRemove > 0 && bank[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			toRemove--
		}

		stack = append(stack, bank[i])
	}

	if len(stack) > batteriesCount {
		stack = stack[:batteriesCount]
	}

	return tools.MustAtoi(string(stack))
}

func (d *Day) FirstProblem() int {
	var result int

	for _, bank := range d.banks {
		result += findMaxJoltage(bank, 2)
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	for _, bank := range d.banks {
		result += findMaxJoltage(bank, 12)
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
