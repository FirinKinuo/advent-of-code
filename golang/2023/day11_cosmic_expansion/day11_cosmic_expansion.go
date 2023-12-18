package main

import (
	"fmt"
	"github.com/FirinKinuo/advent-of-code"
	"log"
	"math"
	"strings"
)

type galaxyCoords struct {
	x int
	y int
}

type Day struct {
	*problem.DayTemplate
	galaxiesMap    []string
	galaxiesCoords []galaxyCoords
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2023", "day11_cosmic_expansion", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %s", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	for _, line := range strings.Split(input, "\r\n") {
		d.galaxiesMap = append(d.galaxiesMap, line)
	}
}

func (d *Day) findGalaxiesCoords(offset int) {
	d.galaxiesCoords = []galaxyCoords{}
	for i, row := range d.galaxiesMap {
		for j, point := range row {
			if point == '#' {
				d.galaxiesCoords = append(d.galaxiesCoords, galaxyCoords{x: i, y: j})
			}
		}
	}

	yInsertOffset := 0
	xInsertOffset := 0
	for i, _ := range d.galaxiesMap[0] {
		if d.isNoGalaxiesInColumn(i) {
			for j, _ := range d.galaxiesCoords {
				if d.galaxiesCoords[j].y-yInsertOffset > i {
					d.galaxiesCoords[j].y += offset
				}
			}
			yInsertOffset += offset
		}
	}

	for i, row := range d.galaxiesMap {
		if !strings.Contains(row, "#") {
			for j, _ := range d.galaxiesCoords {
				if d.galaxiesCoords[j].x-xInsertOffset > i {
					d.galaxiesCoords[j].x += offset
				}
			}
			xInsertOffset += offset
		}
	}
}

func (d *Day) isNoGalaxiesInColumn(columnIndex int) bool {
	for _, m := range d.galaxiesMap {
		if m[columnIndex] == '#' {
			return false
		}
	}

	return true
}

func (d *Day) sumPairsDistances() int {
	stepsSum := 0
	for i, coords := range d.galaxiesCoords {
		for _, pairCoords := range d.galaxiesCoords[i+1:] {
			steps := int(math.Abs(float64(coords.x-pairCoords.x)) + math.Abs(float64(coords.y-pairCoords.y)))
			stepsSum += steps
		}
	}

	return stepsSum
}

func (d *Day) FirstProblem() int {
	d.findGalaxiesCoords(1)
	return d.sumPairsDistances()
}

func (d *Day) SecondProblem() int {
	d.findGalaxiesCoords(999_999_999)
	return d.sumPairsDistances()
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
