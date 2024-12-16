package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

type Robot struct {
	position image.Point
	velocity image.Point
}

type Day struct {
	*problem.DayTemplate
	inputType problem.InputType
	robotsP1  []Robot
	robotsP2  []Robot
	gridSize  image.Rectangle
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2024", "day14_restroom_redoubt", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template, inputType: inputType}, nil
}

func (d *Day) PrepareInput(input string) {
	robotsStrings := strings.Split(input, "\n")

	for _, robotString := range robotsStrings {
		parts := strings.Split(robotString, " ")
		robot := Robot{
			position: tools.PointFromSlice(tools.AtoiSlice(strings.Split(parts[0][2:], ","))),
			velocity: tools.PointFromSlice(tools.AtoiSlice(strings.Split(parts[1][2:], ","))),
		}

		d.robotsP1 = append(d.robotsP1, robot)
		d.robotsP2 = append(d.robotsP2, robot)

	}

	if d.inputType == problem.ProblemInput {
		d.gridSize = image.Rect(0, 0, 101, 103)
	} else {
		d.gridSize = image.Rect(0, 0, 11, 7)
	}
}

func (d *Day) FirstProblem() int {
	result := 1
	for i := 0; i < 100; i++ {
		for i, robot := range d.robotsP1 {
			robot.position = robot.position.Add(robot.velocity)
			if robot.position.X >= d.gridSize.Dx() {
				robot.position.X -= d.gridSize.Dx()
			}

			if robot.position.X < 0 {
				robot.position.X += d.gridSize.Dx()
			}

			if robot.position.Y >= d.gridSize.Dy() {
				robot.position.Y -= d.gridSize.Dy()
			}

			if robot.position.Y < 0 {
				robot.position.Y += d.gridSize.Dy()
			}

			d.robotsP1[i] = robot
		}
	}

	separator := image.Point{
		X: d.gridSize.Dx() / 2,
		Y: d.gridSize.Dy() / 2,
	}

	quadrants := [4]image.Rectangle{
		image.Rect(0, 0, separator.X, separator.Y),
		image.Rect(separator.X+1, 0, d.gridSize.Dx(), separator.Y),
		image.Rect(0, separator.Y+1, separator.X, d.gridSize.Dy()),
		image.Rect(separator.X+1, separator.Y+1, d.gridSize.Dx(), d.gridSize.Dy()),
	}

	robotsInQuadrant := make(map[image.Rectangle]int, len(quadrants))

	for _, robot := range d.robotsP1 {
		for _, quadrant := range quadrants {
			if robot.position.In(quadrant) {
				robotsInQuadrant[quadrant]++
			}
		}
	}

	for _, robotsCount := range robotsInQuadrant {
		result *= robotsCount
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	for i := 1; i <= 10000; i++ {
		img := image.NewRGBA(d.gridSize)
		draw.Draw(img, img.Bounds(), &image.Uniform{C: color.RGBA{A: 255}}, image.Point{}, draw.Src)
		pointColor := color.RGBA{R: 255, G: 255, B: 255, A: 255}

		for i, robot := range d.robotsP2 {
			robot.position = robot.position.Add(robot.velocity)
			if robot.position.X >= d.gridSize.Dx() {
				robot.position.X -= d.gridSize.Dx()
			}

			if robot.position.X < 0 {
				robot.position.X += d.gridSize.Dx()
			}

			if robot.position.Y >= d.gridSize.Dy() {
				robot.position.Y -= d.gridSize.Dy()
			}

			if robot.position.Y < 0 {
				robot.position.Y += d.gridSize.Dy()
			}

			img.Set(robot.position.X, robot.position.Y, pointColor)
			d.robotsP2[i] = robot
		}

		file, err := os.Create(fmt.Sprintf("images/output_%d.png", i))
		if err != nil {
			panic(fmt.Sprintf("Cannot create output file: %v", err))
		}
		_ = png.Encode(file, img)
		_ = file.Close()
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
