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

	devices map[string][]string
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2025", "day11_reactor", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template, devices: make(map[string][]string)}, nil
}

func (d *Day) PrepareInput(input string) {
	for _, line := range tools.SplitNewLines(input) {
		parts := strings.Split(line, ": ")

		d.devices[parts[0]] = strings.Split(parts[1], " ")
	}
}

func (d *Day) FirstProblem() int {
	var result int

	memo := make(map[string]int)

	var dfs func(node string) int
	dfs = func(node string) int {
		if node == "out" {
			return 1
		}

		if val, ok := memo[node]; ok {
			return val
		}

		total := 0
		for _, neighbor := range d.devices[node] {
			total += dfs(neighbor)
		}

		memo[node] = total
		return total
	}

	result = dfs("you")

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	memo := make(map[string]map[int]int)

	var dfs func(node string, state int) int
	dfs = func(node string, state int) int {
		if node == "out" {
			if state == 0b11 {
				return 1
			}

			return 0
		}

		if states, ok := memo[node]; ok {
			if val, ok := states[state]; ok {
				return val
			}
		}

		total := 0
		for _, neighbor := range d.devices[node] {
			newState := state

			if neighbor == "dac" {
				newState |= 0b1
			}
			if neighbor == "fft" {
				newState |= 0b10
			}

			total += dfs(neighbor, newState)
		}

		if _, ok := memo[node]; !ok {
			memo[node] = make(map[int]int)
		}

		memo[node][state] = total

		return total
	}

	result = dfs("svr", 0)

	return result
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
