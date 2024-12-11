package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

type Day struct {
	*problem.DayTemplate

	stones []int
	cache  map[cacheKey]int
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2024", "day11_plutonian_pebbles", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template, cache: make(map[cacheKey]int)}, nil
}

func (d *Day) PrepareInput(input string) {
	d.stones = tools.AtoiSlice(strings.Split(input, " "))
}

type cacheKey struct {
	stone  int
	blinks int
}

func (d *Day) calculateStone(stone, blinks int) int {
	if blinks == 0 {
		return 1
	}

	key := cacheKey{stone: stone, blinks: blinks}
	if result, ok := d.cache[key]; ok {
		return result
	}

	digitsCount := int(math.Log10(float64(stone))) + 1

	var result int
	switch {
	case stone == 0:
		return d.calculateStone(1, blinks-1)
	case digitsCount%2 == 0:
		power := int(math.Pow(10, float64(digitsCount/2)))
		result = d.calculateStone(stone/power, blinks-1) + d.calculateStone(stone%power, blinks-1)
	default:
		result = d.calculateStone(stone*2024, blinks-1)
	}

	d.cache[key] = result
	return result
}

func (d *Day) FirstProblem() int {
	var result int

	for _, stone := range d.stones {
		result += d.calculateStone(stone, 25)
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	for _, stone := range d.stones {
		result += d.calculateStone(stone, 75)
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
