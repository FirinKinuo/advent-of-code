package main

import (
	"fmt"
	"image"
	"log"
	"strings"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

type Day struct {
	*problem.DayTemplate

	antennasCoords map[rune][]image.Point
	grid           image.Rectangle
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2024", "day08_resonant_collinearity", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template, antennasCoords: make(map[rune][]image.Point)}, nil
}

func (d *Day) PrepareInput(input string) {
	grid := strings.Split(input, "\n")

	d.grid = image.Rectangle{Max: image.Point{X: len(grid), Y: len(grid[0])}}

	for rowIndex, row := range grid {
		for columnIndex, symbol := range row {
			if symbol != '.' {
				d.antennasCoords[symbol] = append(d.antennasCoords[symbol], image.Point{X: rowIndex, Y: columnIndex})
			}
		}
	}
}

func (d *Day) FirstProblem() int {
	antinodes := make(tools.UniquePoints)

	for _, antennaCoords := range d.antennasCoords {
		for i := range antennaCoords {
			for j := range antennaCoords {
				if i == j {
					continue
				}
				antinodeCoords := antennaCoords[j].Sub(antennaCoords[i]).Add(antennaCoords[j])

				if antinodeCoords.In(d.grid) {
					antinodes.Add(antinodeCoords)
				}
			}
		}
	}

	return len(antinodes)
}

func (d *Day) SecondProblem() int {
	antinodes := make(tools.UniquePoints)

	for _, antennaCoords := range d.antennasCoords {
		for i := range antennaCoords {
			antinodes.Add(antennaCoords[i])

			for j := range antennaCoords {
				if i == j {
					continue
				}

				vectors := antennaCoords[j].Sub(antennaCoords[i])

				for n := 0; ; n++ {
					antinodeCoords := antennaCoords[j].Add(vectors.Mul(n + 1))
					if !antinodeCoords.In(d.grid) {
						break
					}
					antinodes.Add(antinodeCoords)
				}
			}
		}
	}

	return len(antinodes)
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
