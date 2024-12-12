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

	plantsGrid [][]rune
	gridSize   image.Rectangle
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2024", "day12_garden_groups", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	plantsStrings := strings.Split(input, "\n")

	for _, plants := range plantsStrings {
		d.plantsGrid = append(d.plantsGrid, []rune(plants))
	}

	d.gridSize = image.Rectangle{Max: image.Point{X: len(d.plantsGrid), Y: len(d.plantsGrid[0])}}
}

var offsets = [4]image.Point{
	{X: -1},
	{X: 1},
	{Y: -1},
	{Y: 1},
}

func (d *Day) visitRegion(plantCoords image.Point, visited tools.UniqueAny) (int, tools.UniqueAny) {
	area := 0
	borders := make(tools.UniqueAny)

	d.visit(plantCoords, borders, visited, &area)
	return area, borders
}

func (d *Day) visit(plantCoords image.Point, borders tools.UniqueAny, visited tools.UniqueAny, area *int) {
	if visited.Exists(plantCoords) {
		return
	}
	visited.Add(plantCoords)
	*area++
	for _, offset := range offsets {
		nextPlantCoords := plantCoords.Add(offset)
		if nextPlantCoords.In(d.gridSize) &&
			d.plantsGrid[nextPlantCoords.X][nextPlantCoords.Y] == d.plantsGrid[plantCoords.X][plantCoords.Y] {
			d.visit(nextPlantCoords, borders, visited, area)
		} else {
			borders.Add([2]image.Point{plantCoords, nextPlantCoords})
		}
	}
}

func (d *Day) countSides(borders tools.UniqueAny) int {
	visited := make(tools.UniqueAny)

	numSides := 0
	for side := range borders {
		if visited.Exists(side) {
			continue
		}
		numSides++
		d.visitSide(side.([2]image.Point), borders, visited)
	}
	return numSides
}

func (d *Day) visitSide(sides [2]image.Point, borders tools.UniqueAny, visited tools.UniqueAny) {
	if visited.Exists(sides) || !borders.Exists(sides) {
		return
	}

	visited.Add(sides)

	if sides[0].X == sides[1].X {
		d.visitSide([2]image.Point{sides[0].Add(offsets[0]), sides[1].Add(offsets[0])}, borders, visited)
		d.visitSide([2]image.Point{sides[0].Add(offsets[1]), sides[1].Add(offsets[1])}, borders, visited)
	} else {
		d.visitSide([2]image.Point{sides[0].Add(offsets[2]), sides[1].Add(offsets[2])}, borders, visited)
		d.visitSide([2]image.Point{sides[0].Add(offsets[3]), sides[1].Add(offsets[3])}, borders, visited)
	}
}

func (d *Day) FirstProblem() int {
	var result int
	visited := make(tools.UniqueAny)

	for x := 0; x < d.gridSize.Max.X; x++ {
		for y := 0; y < d.gridSize.Max.Y; y++ {
			plant := image.Point{X: x, Y: y}
			if !visited.Exists(plant) {
				area, borders := d.visitRegion(plant, visited)
				result += area * len(borders)
			}
		}
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int
	visited := make(tools.UniqueAny)

	for x := 0; x < d.gridSize.Max.X; x++ {
		for y := 0; y < d.gridSize.Max.Y; y++ {
			plant := image.Point{X: x, Y: y}
			if !visited.Exists(plant) {
				area, borders := d.visitRegion(plant, visited)
				sides := d.countSides(borders)
				result += area * sides
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
