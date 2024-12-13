package main

import (
	"fmt"
	"github.com/FirinKinuo/advent-of-code/tools"
	"image"
	"log"
	"strings"

	"github.com/FirinKinuo/advent-of-code"
)

type machine struct {
	a image.Point
	b image.Point

	prize image.Point
}

type Day struct {
	*problem.DayTemplate

	machines []machine
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2024", "day13_claw_contraption", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	machines := strings.Split(input, "\n\n")

	for _, m := range machines {
		newMachine := machine{}
		lines := strings.Split(m, "\n")
		for i := 0; i < 2; i++ {
			button := tools.AtoiSlice(strings.Split(lines[i][12:], ", Y+"))
			if i == 0 {
				newMachine.a = image.Point{X: button[0], Y: button[1]}
			} else {
				newMachine.b = image.Point{X: button[0], Y: button[1]}
			}
		}
		prize := tools.AtoiSlice(strings.Split(lines[2][9:], ", Y="))
		newMachine.prize = image.Point{X: prize[0], Y: prize[1]}
		d.machines = append(d.machines, newMachine)
	}
}

func (d *Day) evaluateCost(buttonA, buttonB, prize image.Point, addPrize int) int {
	b := ((prize.Y+addPrize)*buttonA.X - (prize.X+addPrize)*buttonA.Y) / (buttonB.Y*buttonA.X - buttonB.X*buttonA.Y)
	a := ((prize.X + addPrize) - b*buttonB.X) / buttonA.X
	if a*buttonA.X+b*buttonB.X == (prize.X+addPrize) && a*buttonA.Y+b*buttonB.Y == (prize.Y+addPrize) {
		return 3*a + b
	}

	return 0
}

func (d *Day) FirstProblem() int {
	var result int

	for _, m := range d.machines {
		result += d.evaluateCost(m.a, m.b, m.prize, 0)
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	for _, m := range d.machines {
		result += d.evaluateCost(m.a, m.b, m.prize, 10000000000000)
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
