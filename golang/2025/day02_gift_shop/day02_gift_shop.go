package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

type Day struct {
	*problem.DayTemplate
	idRanges [][2]int
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2025", "day02_gift_shop", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	ranges := strings.Split(input, ",")

	d.idRanges = make([][2]int, len(ranges))

	for i, r := range ranges {
		edges := strings.Split(r, "-")
		start := tools.MustAtoi(edges[0])
		end := tools.MustAtoi(edges[1])

		d.idRanges[i] = [2]int{start, end}
	}
}

func (d *Day) FirstProblem() int {
	var result int

	for _, idRange := range d.idRanges {
		for i := idRange[0]; i <= idRange[1]; i++ {
			str := strconv.Itoa(i)
			n := len(str)
			if n%2 != 0 {
				continue
			}

			if str[0:n/2] == str[n/2:] {
				result += i
			}
		}
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	for _, idRange := range d.idRanges {
		for i := idRange[0]; i <= idRange[1]; i++ {
			str := strconv.Itoa(i)
			n := len(str)

			for patternLen := 1; patternLen <= n/2; patternLen++ {
				if n%patternLen != 0 {
					continue
				}

				pattern := str[:patternLen]
				repeats := n / patternLen
				constructed := ""

				for j := 0; j < repeats; j++ {
					constructed += pattern
				}

				if constructed == str && repeats > 1 {
					result += i
					break
				}
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
