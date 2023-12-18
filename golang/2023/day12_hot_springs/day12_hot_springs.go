package main

import (
	"fmt"
	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
	"log"
	"strconv"
	"strings"
)

type condition struct {
	springs string
	states  []int
}

type Day struct {
	*problem.DayTemplate
	conditions             []condition
	countArrangementsCache map[string]int
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2023", "day12_hot_springs", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %s", err)
	}
	return &Day{DayTemplate: template, countArrangementsCache: make(map[string]int)}, nil
}

func (d *Day) PrepareInput(input string) {
	for _, line := range tools.SplitNewLines(input) {
		parts := strings.Split(line, " ")
		d.conditions = append(d.conditions, condition{
			springs: parts[0],
			states:  d.parseStates(parts[1]),
		})
	}
}

func (d *Day) parseStates(in string) []int {
	state := make([]int, 0, len(in))
	for _, num := range strings.Split(in, ",") {
		state = append(state, tools.MustAtoi(num))
	}

	return state
}

func (d *Day) createCacheKey(springs string, states []int) string {
	var buf strings.Builder
	buf.WriteString(springs)
	for _, s := range states {
		buf.WriteString(" " + strconv.Itoa(s))
	}

	return buf.String()
}

func (d *Day) countArrangements(springs string, states []int) int {
	cacheKey := d.createCacheKey(springs, states)
	value, ok := d.countArrangementsCache[cacheKey]
	if ok {
		return value
	}

	if len(states) == 0 {
		if strings.Contains(springs, "#") {
			return 0
		}
		return 1
	}
	if len(springs) == 0 {
		if len(states) == 0 {
			return 1
		}
		return 0
	}

	result := 0

	if strings.ContainsAny(springs[:1], ".?") {
		result += d.countArrangements(springs[1:], states)
	}

	if strings.ContainsAny(springs[:1], "#?") {
		if states[0] <= len(springs) &&
			!strings.Contains(springs[:states[0]], ".") &&
			(states[0] == len(springs) || springs[states[0]] != '#') {
			nextSprings := ""
			if states[0]+1 < len(springs) {
				nextSprings = springs[states[0]+1:]
			}
			result += d.countArrangements(nextSprings, states[1:])
		}
	}

	d.countArrangementsCache[cacheKey] = result
	return result
}

func (d *Day) FirstProblem() int {
	totalArrangements := 0

	for _, c := range d.conditions {
		totalArrangements += d.countArrangements(c.springs, c.states)
	}

	return totalArrangements
}

func (d *Day) SecondProblem() int {
	totalArrangements := 0

	for _, c := range d.conditions {
		foldedUpSprings := strings.Join([]string{c.springs, c.springs, c.springs, c.springs, c.springs}, "?")
		foldedUpStates := make([]int, 0, len(c.states)*5)

		for i := 0; i < 5; i++ {
			foldedUpStates = append(foldedUpStates, c.states...)
		}

		totalArrangements += d.countArrangements(foldedUpSprings, foldedUpStates)
	}

	return totalArrangements
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
