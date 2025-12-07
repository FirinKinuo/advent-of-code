package main

import (
	"fmt"
	"image"
	"log"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

type Day struct {
	*problem.DayTemplate

	grid  [][]rune
	start image.Point
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2025", "day07_laboratories", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	lines := tools.SplitNewLines(input)

	d.grid = make([][]rune, len(lines))

	for i, line := range lines {
		d.grid[i] = []rune(line)
	}

	for x, symbol := range d.grid[0] {
		if symbol == 'S' {
			d.start = image.Point{X: x, Y: 0}
			break
		}
	}
}

func (d *Day) FirstProblem() int {
	var result int
	columnLength := len(d.grid)
	rowLength := len(d.grid[0])

	queue := make([]image.Point, 0, columnLength*rowLength)
	queue = append(queue, d.start)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Y == columnLength-1 {
			continue
		}

		nextY := current.Y + 1

		switch d.grid[nextY][current.X] {
		case '.':
			d.grid[nextY][current.X] = '|'
			queue = append(queue, image.Point{X: current.X, Y: nextY})
		case '^':
			leftX := current.X - 1
			if leftX >= 0 {
				if d.grid[nextY][leftX] == '.' {
					d.grid[nextY][leftX] = '|'
					queue = append(queue, image.Point{X: leftX, Y: nextY})
				}
			}

			rightX := current.X + 1
			if rightX < rowLength {
				if d.grid[nextY][rightX] == '.' {
					d.grid[nextY][rightX] = '|'
					queue = append(queue, image.Point{X: rightX, Y: nextY})
				}
			}

			result++
		}
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int
	columnLength := len(d.grid)
	rowLength := len(d.grid[0])

	ways := make([][]int, columnLength)
	for i := range ways {
		ways[i] = make([]int, rowLength)
	}

	ways[d.start.Y][d.start.X] = 1

	for y := d.start.Y; y < columnLength-1; y++ {
		for x := 0; x < rowLength; x++ {
			count := ways[y][x]
			if count == 0 {
				continue
			}

			switch d.grid[y+1][x] {
			case '^':
				if x > 0 {
					ways[y+1][x-1] += count
				}

				if x < rowLength-1 {
					ways[y+1][x+1] += count
				}
			case '.', '|':
				ways[y+1][x] += count
			}
		}
	}

	result = tools.Sum(ways[columnLength-1]...)

	return result
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
