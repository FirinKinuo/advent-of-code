package main

import (
	"fmt"
	"github.com/FirinKinuo/advent-of-code"
	"log"
	"math"
	"strconv"
	"strings"
)

type Day struct {
	*problem.DayTemplate
	time     []int
	distance []int
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2023", "day_06_wait_for_it", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %s", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	for i, line := range strings.Split(input, "\r\n") {
		switch i {
		case 0:
			d.time = d.parseInts(line[5:])
		case 1:
			d.distance = d.parseInts(line[9:])
		}
	}
}

func (d *Day) parseInts(input string) []int {
	parsedValues := make([]int, 0, len(input)/2)
	for _, s := range strings.Split(input, " ") {
		value, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		parsedValues = append(parsedValues, value)
	}

	return parsedValues
}

func (d *Day) countWin(time, distance float64) int {
	a := math.Ceil((time / 2) - math.Sqrt(math.Pow(time/2.0, 2.0)-distance))
	b := math.Floor((time / 2) + math.Sqrt(math.Pow(time/2.0, 2.0)-distance))

	return int(b - a + 1)
}

func (d *Day) calcWins(time []int, distance []int) int {
	wins := 1

	for i := 0; i < len(time); i++ {
		wins *= d.countWin(float64(time[i]), float64(distance[i]+1))
	}

	return wins
}

func (d *Day) improveNumber(input []int) int {
	numberString := ""

	for _, num := range input {
		numberString += strconv.Itoa(num)
	}

	improvedNumber, _ := strconv.Atoi(numberString)

	return improvedNumber
}

func (d *Day) FirstProblem() int {
	return d.calcWins(d.time, d.distance)
}

func (d *Day) SecondProblem() int {
	return d.calcWins([]int{d.improveNumber(d.time)}, []int{d.improveNumber(d.distance)})
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
