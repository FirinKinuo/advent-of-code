package main

import (
	"fmt"
	"github.com/FirinKinuo/advent-of-code"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Day struct {
	*problem.DayTemplate
	games []string
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2023", "day02_cube_conundrum", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %s", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	d.games = strings.Split(input, "\r\n")
}

func (d *Day) FirstProblem() int {
	maxCubesForGame := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	possibleGamesSum := 0

	for i, gameString := range d.games {
		bags := strings.TrimPrefix(gameString, fmt.Sprintf("Game %d:", i+1))
		if d.isGamePossible(bags, maxCubesForGame) {
			possibleGamesSum += i + 1
		}
	}

	return possibleGamesSum
}

func (d *Day) isGamePossible(bags string, maxCubes map[string]int) bool {
	var compRegEx = regexp.MustCompile(`(?P<count>\d+)\s(?P<color>\w+)`)
	for _, cube := range compRegEx.FindAllStringSubmatch(bags, -1) {
		count, _ := strconv.Atoi(cube[1])
		if maxCubes[cube[2]] < count {
			return false
		}
	}

	return true
}

func (d *Day) SecondProblem() int {
	sumPowGamesSets := 0

	for i, gameString := range d.games {
		bags := strings.TrimPrefix(gameString, fmt.Sprintf("Game %d:", i+1))
		set := d.findMinimalPossibleCubesForGame(bags)

		setsPow := 1
		for _, count := range set {
			setsPow *= count
		}

		sumPowGamesSets += setsPow
	}

	return sumPowGamesSets
}

func (d *Day) findMinimalPossibleCubesForGame(bags string) map[string]int {
	minimalCubes := make(map[string]int)

	var compRegEx = regexp.MustCompile(`(?P<count>\d+)\s(?P<color>\w+)`)
	for _, cube := range compRegEx.FindAllStringSubmatch(bags, -1) {
		count, _ := strconv.Atoi(cube[1])
		if minimalCubes[cube[2]] < count {
			minimalCubes[cube[2]] = count
		}
	}
	return minimalCubes
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
