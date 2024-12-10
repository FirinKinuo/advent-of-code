package main

import (
	"container/list"
	"fmt"
	"image"
	"log"
	"strings"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

type Day struct {
	*problem.DayTemplate
	topographicMap     [][]int
	topographicMapSize image.Rectangle
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2024", "day10_hoof_it", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	for _, line := range strings.Split(input, "\n") {
		d.topographicMap = append(d.topographicMap, tools.AtoiSlice(strings.Split(line, "")))
	}

	d.topographicMapSize = image.Rectangle{
		Max: image.Point{
			X: len(d.topographicMap),
			Y: len(d.topographicMap[0]),
		},
	}
}

var directions = []image.Point{
	{X: -1}, {X: 1}, {Y: -1}, {Y: 1},
}

func (d *Day) bfs(startPoint image.Point) int {
	visited := make(tools.UniquePoints)

	queue := list.New()
	queue.PushBack(startPoint)

	visited.Add(startPoint)
	var reachableNineCount int

	for queue.Len() > 0 {
		point := queue.Remove(queue.Front()).(image.Point)

		if d.topographicMap[point.X][point.Y] == 9 {
			reachableNineCount++
		}

		for _, dir := range directions {
			nextPoint := point.Add(dir)
			if _, ok := visited[nextPoint]; !ok &&
				nextPoint.In(d.topographicMapSize) &&
				d.topographicMap[nextPoint.X][nextPoint.Y] == d.topographicMap[point.X][point.Y]+1 {
				visited.Add(nextPoint)
				queue.PushBack(nextPoint)
			}
		}
	}

	return reachableNineCount
}

func (d *Day) dfs(point image.Point, memo map[image.Point]int) int {
	if count, ok := memo[point]; ok {
		return count
	}

	if d.topographicMap[point.X][point.Y] == 9 {
		return 1
	}

	trailCount := 0

	for _, dir := range directions {
		nextPoint := point.Add(dir)
		if nextPoint.In(d.topographicMapSize) &&
			d.topographicMap[nextPoint.X][nextPoint.Y] == d.topographicMap[point.X][point.Y]+1 {
			trailCount += d.dfs(nextPoint, memo)
		}
	}

	memo[point] = trailCount
	return trailCount
}

func (d *Day) FirstProblem() int {
	var result int

	for x := range d.topographicMap {
		for y, value := range d.topographicMap[x] {
			if value == 0 {
				result += d.bfs(image.Point{X: x, Y: y})
			}
		}
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	memory := make(map[image.Point]int)

	for x := range d.topographicMap {
		for y, value := range d.topographicMap[x] {
			if value == 0 {
				result += d.dfs(image.Point{X: x, Y: y}, memory)
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
