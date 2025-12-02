package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

const (
	initialPosition = 50
)

const (
	rotateRight = 'R'
)

type Day struct {
	*problem.DayTemplate

	instructions []int
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2025", "day01_secret_entrance", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	inputLines := tools.SplitNewLines(input)

	d.instructions = make([]int, 0, len(inputLines))

	for _, line := range inputLines {
		direction := line[0]

		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(fmt.Sprintf("failed to parse steps: %v", err))
		}

		if direction == rotateRight {
			d.instructions = append(d.instructions, steps)
		} else {
			d.instructions = append(d.instructions, -steps)
		}
	}

}

func (d *Day) FirstProblem() int {
	var result int

	position := initialPosition

	for _, delta := range d.instructions {
		position = (position + delta) % 100

		if position < 0 {
			position += 100
		}

		if position == 0 {
			result++
		}
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	position := initialPosition

	// I tried to come up with a smarter solution for about 15 minutes,
	// then gave up on everything and just stupidly brute forced it xD
	for _, delta := range d.instructions {
		for i := 0; i < tools.Abs(delta); i++ {
			position = position + tools.Sign(delta)
			if position%100 == 0 {
				result++
			}
			if position < 0 {
				position += 100
			} else if position > 99 {
				position -= 100
			}
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
