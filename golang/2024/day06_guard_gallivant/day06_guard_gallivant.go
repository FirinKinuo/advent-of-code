package main

import (
	"fmt"
	"github.com/FirinKinuo/advent-of-code/tools"
	"log"
	"maps"
	"strconv"
	"strings"
	"sync"

	"github.com/FirinKinuo/advent-of-code"
)

type Day struct {
	*problem.DayTemplate

	grid [][]rune
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2024", "day06_guard_gallivant", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	gridStrings := strings.Split(input, "\n")

	for _, str := range gridStrings {
		d.grid = append(d.grid, []rune(str))
	}
}

func (d *Day) findStartPosition() [2]int {
	for rowIndex, row := range d.grid {
		for columnIndex, symbol := range row {
			if symbol == '^' {
				return [2]int{rowIndex, columnIndex}
			}
		}
	}

	panic("cannot find start position, check input")
}

const (
	axisX = 0
	axisY = 1
)

func (d *Day) goToObstacle(grid [][]rune, direction rune, currentPosition [2]int) (map[string]struct{}, [2]int, bool) {
	visitedTemp := make(map[string]struct{})

	var axis int
	var moveForward bool
	switch direction {
	case '^':
		axis = axisX
		moveForward = true
	case '>':
		axis = axisY
		moveForward = false
	case 'v':
		axis = axisX
		moveForward = false
	case '<':
		axis = axisY
		moveForward = true
	}

	var end, step int
	if moveForward {
		end = 0
		step = -1
	} else {
		end = len(grid) - 1
		step = 1
	}

	for i := currentPosition[axis]; i != end; i += step {
		if axis == axisX {
			visitedTemp[strconv.Itoa(i)+":"+strconv.Itoa(currentPosition[axisY])] = struct{}{}
			if grid[i+step][currentPosition[axisY]] == '#' {
				return visitedTemp, [2]int{i, currentPosition[axisY]}, false
			}
		} else {
			visitedTemp[strconv.Itoa(currentPosition[axisX])+":"+strconv.Itoa(i)] = struct{}{}
			if grid[currentPosition[axisX]][i+step] == '#' {
				return visitedTemp, [2]int{currentPosition[axisX], i}, false
			}
		}
	}

	return visitedTemp, currentPosition, true
}

var directionRotate = map[rune]rune{
	'^': '>',
	'>': 'v',
	'v': '<',
	'<': '^',
}

func (d *Day) FirstProblem() int {
	currentPosition := d.findStartPosition()
	visited := map[string]struct{}{strconv.Itoa(currentPosition[0]) + ":" + strconv.Itoa(currentPosition[1]): {}}
	direction := d.grid[currentPosition[0]][currentPosition[1]]

	for {
		visitedPoints, newPosition, leaves := d.goToObstacle(d.grid, direction, currentPosition)
		maps.Copy(visited, visitedPoints)
		currentPosition = newPosition
		direction = directionRotate[direction]
		if leaves {
			break
		}
	}

	return len(visited) + 1
}

func (d *Day) SecondProblem() int {
	var result int

	startPosition := d.findStartPosition()

	wg := &sync.WaitGroup{}
	wg.Add(len(d.grid))
	for rowIndex := range d.grid {
		go func(rowIndex int) {
			defer wg.Done()

			for columnIndex := range d.grid[rowIndex] {
				currentPosition := startPosition
				direction := d.grid[currentPosition[0]][currentPosition[1]]
				copiedGrid := tools.Copy2DSlice(d.grid)
				if copiedGrid[rowIndex][columnIndex] == '#' {
					continue
				}
				copiedGrid[rowIndex][columnIndex] = '#'
				visitedPosition := make(map[string]struct{})

				for {
					visitedPositionKey := fmt.Sprintf("%d:%d@%d", currentPosition[0], currentPosition[1], direction)

					if _, ok := visitedPosition[visitedPositionKey]; ok {
						result++
						break
					}

					visitedPosition[visitedPositionKey] = struct{}{}
					_, newPosition, leaves := d.goToObstacle(copiedGrid, direction, currentPosition)
					currentPosition = newPosition
					direction = directionRotate[direction]
					if leaves {
						break
					}
				}
			}
		}(rowIndex)
	}

	wg.Wait()
	return result
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
