package main

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

type Day struct {
	*problem.DayTemplate

	calibrations [][]int
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2024", "day07_bridge_repair", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	calibrations := strings.Split(input, "\n")

	d.calibrations = make([][]int, len(calibrations))

	for i, calibration := range calibrations {
		parts := strings.Split(calibration, ": ")

		d.calibrations[i] = append(d.calibrations[i], tools.MustAtoi(parts[0]))
		d.calibrations[i] = append(d.calibrations[i], tools.AtoiSlice(strings.Split(parts[1], " "))...)
	}

}

func (d *Day) evaluate(target int, nums []int, withConcatenationOperator bool) bool {
	for _, operations := range d.generateAllOperationsSet(len(nums), withConcatenationOperator) {
		result := nums[0]
		for i, num := range nums[1:] {
			switch operations[i] {
			case '+':
				result += num
			case '*':
				result *= num
			case '|':
				result = d.concatenate(result, num)
			}

			if result > target {
				break
			}
		}

		if result == target {
			return true
		}
	}

	return false
}

func (d *Day) generateAllOperationsSet(calibrationsLen int, withConcatenationOperator bool) [][]rune {
	operators := []rune{'+', '*'}
	if withConcatenationOperator {
		operators = append(operators, '|')
	}

	var result [][]rune
	var currentCombination []rune

	d.generateNextOperationSet(0, calibrationsLen, currentCombination, operators, &result)

	return result
}

func (d *Day) generateNextOperationSet(
	position, calibrationsLen int,
	currentCombination, operators []rune,
	result *[][]rune,
) {
	if position == calibrationsLen-1 {
		*result = append(*result, append([]rune{}, currentCombination...))
		return
	}

	for _, operator := range operators {
		currentCombination = append(currentCombination, operator)
		d.generateNextOperationSet(position+1, calibrationsLen, currentCombination, operators, result)
		currentCombination = currentCombination[:len(currentCombination)-1]
	}
}

func (d *Day) concatenate(a, b int) int {
	numDigits := 0
	temp := b
	for temp > 0 {
		numDigits++
		temp /= 10
	}

	result := 1
	for i := 0; i < numDigits; i++ {
		result *= 10
	}

	return a*result + b
}

func (d *Day) FirstProblem() int {
	var result int

	wg := sync.WaitGroup{}
	wg.Add(len(d.calibrations))

	for _, calibrations := range d.calibrations {
		go func() {
			defer wg.Done()
			if d.evaluate(calibrations[0], calibrations[1:], false) {
				result += calibrations[0]
			}
		}()
	}

	wg.Wait()

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	wg := sync.WaitGroup{}
	wg.Add(len(d.calibrations))

	for _, calibrations := range d.calibrations {
		go func() {
			defer wg.Done()
			if d.evaluate(calibrations[0], calibrations[1:], true) {
				result += calibrations[0]
			}
		}()
	}

	wg.Wait()

	return result
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
