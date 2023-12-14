package main

import (
	"fmt"
	"github.com/FirinKinuo/advent-of-code"
	"log"
	"strings"
)

var connections = map[rune][2][2]int{
	'-': {{0, -1}, {0, 1}},
	'|': {{-1, 0}, {1, 0}},
	'L': {{-1, 0}, {0, 1}},
	'J': {{-1, 0}, {0, -1}},
	'7': {{1, 0}, {0, -1}},
	'F': {{1, 0}, {0, 1}},
	'.': {{0, 0}, {0, 0}},
}

type Day struct {
	*problem.DayTemplate
	matrix     [][]rune
	matrixEdge int
	verticesX  []int
	verticesY  []int
	startPos   [2]int
	steps      int
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2023", "day10_pipe_maze", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %s", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	data := strings.Split(input, "\r\n")
	d.matrix = make([][]rune, 0, len(data))

	for i, line := range data {
		if startPosIndex := strings.Index(line, "S"); startPosIndex != -1 {
			d.startPos = [2]int{i, startPosIndex}
		}

		d.matrix = append(d.matrix, []rune(line))
	}

	d.matrixEdge = len(d.matrix) - 1
}

var neighbors = [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func (d *Day) getPointAtMatrix(coords [2]int) rune {
	return d.matrix[coords[0]][coords[1]]
}

func (d *Day) canMove(neighborCoords [2]int, connectionDirection [2]int) bool {
	return neighborCoords[0] == (-connectionDirection[0]) && neighborCoords[1] == (-connectionDirection[1])
}

func (d *Day) nextPoint(pos, previousPos [2]int) [2]int {
	connection := d.matrix[pos[0]][pos[1]]
	connectionDirections := connections[connection]
	for _, direction := range connectionDirections {
		nextPoint := [2]int{pos[0] + direction[0], pos[1] + direction[1]}
		if nextPoint[0] == previousPos[0] && nextPoint[1] == previousPos[1] {
			continue
		}

		if d.matrix[nextPoint[0]][nextPoint[1]] == 'S' {
			d.steps++
			break
		}

		d.steps++
		d.verticesX = append(d.verticesX, nextPoint[0])
		d.verticesY = append(d.verticesY, nextPoint[1])
		d.nextPoint(nextPoint, pos)
	}

	return [2]int{}
}

func (d *Day) FirstProblem() int {
	nextPoint := [2]int{}
	for _, neighbor := range neighbors {
		nextPoint = [2]int{d.startPos[0] + neighbor[0], d.startPos[1] + neighbor[1]}
		if nextPoint[0] < 0 || nextPoint[0] > d.matrixEdge || nextPoint[1] < 0 || nextPoint[1] > d.matrixEdge {
			continue
		}

		pointConnection := d.getPointAtMatrix(nextPoint)
		if pointConnection == '.' {
			continue
		}

		connectionDirections := connections[pointConnection]
		for _, direction := range connectionDirections {
			if !d.canMove(neighbor, direction) {
				continue
			}
			d.verticesX = append(d.verticesX, d.startPos[0])
			d.verticesY = append(d.verticesY, d.startPos[1])
			d.steps++
			d.nextPoint(nextPoint, d.startPos)
		}
		break
	}

	return d.steps / 2
}

func (d *Day) SecondProblem() int {
	area := 0
	for i := 0; i < len(d.verticesX); i++ {
		if i == len(d.verticesX)-1 {
			area += (d.verticesY[i] + d.verticesY[0]) * (d.verticesX[i] - d.verticesX[0])
		} else {
			area += (d.verticesY[i] + d.verticesY[i+1]) * (d.verticesX[i] - d.verticesX[i+1])
		}
	}
	area = area / 2

	return area - len(d.verticesX)/2
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
