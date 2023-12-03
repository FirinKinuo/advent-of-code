package main

import (
	"fmt"
	"github.com/FirinKinuo/advent-of-code"
	"log"
	"path"
	"strconv"
	"strings"
)

type position struct {
	xStart int
	xEnd   int
}

type gear struct {
	value    string
	position position
}

type Day struct {
	problem      *problem.Problem
	engineScheme []string
	rowLen       int
	gearsRow     [][]gear
}

func NewDay(inputType problem.InputType) (*Day, error) {
	p, err := problem.NewProblem(path.Join("2023", "day03_gear_ratios", string(inputType)))
	if err != nil {
		return nil, fmt.Errorf("new problem init: %s", err)
	}

	return &Day{problem: p}, nil
}

func (d *Day) firstProblem(input string) int {
	if d.gearsRow == nil {
		d.scanGears(input)
	}

	gearsSum := 0

	for rowIndex, gearsRow := range d.gearsRow {
		for _, currentGear := range gearsRow {
			countAdjacentNeighbors := d.countAdjacentNeighbors(rowIndex, currentGear)

			if len(countAdjacentNeighbors) > 0 {
				gearValue, _ := strconv.Atoi(currentGear.value)
				gearsSum += gearValue
			}

		}
	}

	return gearsSum
}

func (d *Day) scanGears(input string) {
	d.engineScheme = strings.Split(input, "\r\n")
	d.rowLen = len(d.engineScheme)
	d.gearsRow = make([][]gear, d.rowLen)

	for y, row := range d.engineScheme {
		gearDigits := ""
		for x, char := range row {
			if char >= '0' && char <= '9' {
				gearDigits += string(char)
				if x == d.rowLen-1 {
					d.appendToGearsRow(y, gear{
						value: gearDigits,
						position: position{
							xStart: x - len([]rune(gearDigits)) + 1,
							xEnd:   x,
						},
					})
					gearDigits = ""
				}

			} else {
				if gearDigits != "" {
					d.appendToGearsRow(y, gear{
						value: gearDigits,
						position: position{
							xStart: x - len([]rune(gearDigits)),
							xEnd:   x - 1,
						},
					})
					gearDigits = ""
				}

				if char != '.' {
					d.appendToGearsRow(y, gear{
						value: string(char),
						position: position{
							xStart: x,
							xEnd:   x,
						},
					})
				}
			}
		}
	}
}

func (d *Day) appendToGearsRow(row int, g gear) {
	d.gearsRow[row] = append(d.gearsRow[row], g)
}

func (d *Day) countAdjacentNeighbors(gearRowIndex int, g gear) []gear {
	neighbors := make([]gear, 0, 8)
	for i := gearRowIndex - 1; i <= gearRowIndex+1; i++ {
		if i < 0 || i == d.rowLen {
			continue
		}

		neighborsGears := d.gearsRow[i]
		for _, neighbor := range neighborsGears {
			if neighbor == g {
				continue
			}

			if (g.position.xStart >= neighbor.position.xStart-1 && neighbor.position.xStart <= g.position.xEnd+1) &&
				(g.position.xEnd <= neighbor.position.xEnd+1 && neighbor.position.xEnd >= g.position.xStart-1) ||
				neighbor.position.xStart >= g.position.xStart-1 && neighbor.position.xEnd <= g.position.xEnd+1 {
				neighbors = append(neighbors, neighbor)
			}
		}
	}

	return neighbors
}

func (d *Day) secondProblem(input string) int {
	if d.gearsRow == nil {
		d.scanGears(input)
	}

	gearsPowSum := 0

	for rowIndex, gearsRow := range d.gearsRow {
		for _, currentGear := range gearsRow {
			if currentGear.value == "*" {
				countAdjacentNeighbors := d.countAdjacentNeighbors(rowIndex, currentGear)

				if len(countAdjacentNeighbors) == 2 {
					firstPartValue, err := strconv.Atoi(countAdjacentNeighbors[0].value)
					secondPartValue, err := strconv.Atoi(countAdjacentNeighbors[1].value)
					if err != nil {
						continue
					}

					gearsPowSum += firstPartValue * secondPartValue
				}
			}
		}
	}

	return gearsPowSum
}

func (d *Day) SolveProblems() {
	solvers := []func(input string) int{
		d.firstProblem,
		d.secondProblem,
	}

	d.problem.Solve(solvers)
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.SolveProblems()
}
