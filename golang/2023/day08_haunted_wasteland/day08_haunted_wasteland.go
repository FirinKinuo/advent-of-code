package main

import (
	"fmt"
	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
	"log"
	"strings"
)

type Day struct {
	*problem.DayTemplate
	instructions []int
	network      map[string][]string
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2023", "day08_haunted_wasteland", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %s", err)
	}
	return &Day{DayTemplate: template, network: make(map[string][]string)}, nil
}

func (d *Day) PrepareInput(input string) {
	tabs := strings.Split(input, "\r\n\r\n")
	d.instructions = make([]int, 0, len(tabs[0]))

	for _, instruction := range tabs[0] {
		direction := 0
		if instruction == 'R' {
			direction = 1
		}

		d.instructions = append(d.instructions, direction)
	}

	for _, node := range strings.Split(tabs[1], "\r\n") {
		d.network[node[:3]] = strings.Split(node[7:len(node)-1], ", ")
	}
}

func (d *Day) FirstProblem() int {
	currentNode := "AAA"
	steps := 0
	for currentNode != "ZZZ" {
		for _, instruction := range d.instructions {
			currentNode = d.network[currentNode][instruction]
			steps++
		}
	}

	return steps
}

func (d *Day) SecondProblem() int {
	nodesStarts := make([]string, 0, 10)
	for s, _ := range d.network {
		if s[2] == 'A' {
			nodesStarts = append(nodesStarts, s)
		}
	}

	steps := make([]int, len(nodesStarts))

	for i, start := range nodesStarts {
		currentElement := start
		for currentElement[2] != 'Z' {
			for _, instruction := range d.instructions {
				currentElement = d.network[currentElement][instruction]
				steps[i]++
			}
		}
	}

	return tools.FindLCMOfSlice(steps)
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
