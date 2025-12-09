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

	tiles []image.Point
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2025", "day09_movie_theater", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	lines := tools.SplitNewLines(input)

	d.tiles = make([]image.Point, len(lines))

	for i, line := range lines {
		d.tiles[i] = tools.PointFromSlice(tools.AtoiSlice(strings.Split(line, ",")))
	}
}

func (d *Day) FirstProblem() int {
	var result int

	tilesLen := len(d.tiles)

	for i := 0; i < tilesLen; i++ {
		for j := i + 1; j < tilesLen; j++ {
			point1 := d.tiles[i]
			point2 := d.tiles[j]

			area := (point1.X - point2.X + 1) * (point1.Y - point2.Y + 1)

			if area > result {
				result = area
			}
		}
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	tilesLen := len(d.tiles)

	for i := 0; i < tilesLen; i++ {
		for j := i + 1; j < tilesLen; j++ {
			point1 := d.tiles[i]
			point2 := d.tiles[j]

			area := (point1.X - point2.X + 1) * (point1.Y - point2.Y + 1)

			rectMin := image.Point{
				X: min(point1.X, point2.X),
				Y: min(point1.Y, point2.Y),
			}
			rectMax := image.Point{
				X: max(point1.X, point2.X),
				Y: max(point1.Y, point2.Y),
			}

			valid := true

			for k := 0; k < tilesLen; k++ {
				point1 = d.tiles[k]
				point2 = d.tiles[(k+1)%tilesLen]

				segMin := image.Point{
					X: min(point1.X, point2.X),
					Y: min(point1.Y, point2.Y),
				}

				segMax := image.Point{
					X: max(point1.X, point2.X),
					Y: max(point1.Y, point2.Y),
				}

				if rectMin.X < segMax.X && rectMax.X > segMin.X && rectMin.Y < segMax.Y && rectMax.Y > segMin.Y {
					valid = false
					break
				}
			}

			if valid && area > result {
				result = area
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
