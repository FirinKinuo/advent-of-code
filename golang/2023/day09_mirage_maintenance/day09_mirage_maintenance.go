package main

import (
	"fmt"
	"github.com/FirinKinuo/advent-of-code"
	"log"
	"strconv"
	"strings"
)

type Day struct {
	*problem.DayTemplate
	dataset [][]int
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2023", "day09_mirage_maintenance", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %s", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	dataset := strings.Split(input, "\r\n")
	d.dataset = make([][]int, len(dataset))
	for i, data := range dataset {
		d.dataset[i] = make([]int, 0, len(data))
		for _, num := range strings.Split(data, " ") {
			value, _ := strconv.Atoi(num)

			d.dataset[i] = append(d.dataset[i], value)
		}
	}
}

func (d *Day) increasings(data []int) []int {
	increasings := make([]int, 0, len(data)-1)
	for i := 0; i < len(data)-1; i++ {
		increasings = append(increasings, data[i+1]-data[i])
	}

	return increasings
}

func (d *Day) extrapolateRight(data []int) int {
	if len(data) == 0 {
		return 0
	}

	return data[len(data)-1] + d.extrapolateRight(d.increasings(data))
}

func (d *Day) FirstProblem() int {
	extrapolatedValuesSum := 0
	for _, data := range d.dataset {
		extrapolatedValuesSum += d.extrapolateRight(data)
	}

	return extrapolatedValuesSum
}

func (d *Day) extrapolateLeft(data []int) int {
	if len(data) == 0 {
		return 0
	}

	return data[0] - d.extrapolateLeft(d.increasings(data))
}

func (d *Day) SecondProblem() int {
	extrapolatedValuesSum := 0
	for _, data := range d.dataset {
		extrapolatedValuesSum += d.extrapolateLeft(data)
	}

	return extrapolatedValuesSum
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
