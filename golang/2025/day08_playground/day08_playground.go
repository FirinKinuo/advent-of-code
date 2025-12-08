package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strings"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

type Day struct {
	*problem.DayTemplate

	points []tools.Point3D
	edges  []Edge
}

type Edge struct {
	from     int
	to       int
	distance float64
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2025", "day08_playground", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	lines := tools.SplitNewLines(input)
	d.points = make([]tools.Point3D, len(lines))

	for i, line := range lines {
		coords := tools.AtoiSlice(strings.Split(line, ","))

		d.points[i] = tools.Point3D{X: coords[0], Y: coords[1], Z: coords[2]}
	}

	// yeah, it's cheat, I know :)
	d.buildEdges()
}

func (d *Day) getDistance(p1, p2 tools.Point3D) float64 {
	dx := float64(p1.X - p2.X)
	dy := float64(p1.Y - p2.Y)
	dz := float64(p1.Z - p2.Z)

	return math.Pow(dx, 2) + math.Pow(dy, 2) + math.Pow(dz, 2)
}

func (d *Day) buildEdges() {
	n := len(d.points)
	edges := make([]Edge, 0, n*(n-1)/2)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			distance := d.getDistance(d.points[i], d.points[j])
			edges = append(edges, Edge{i, j, distance})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	d.edges = edges
}

func (d *Day) FirstProblem() int {
	result := 1

	dsu := tools.NewDSU(len(d.points))

	edgeIndex := 0

	for conn := 0; conn < 1000 && edgeIndex < len(d.edges); conn++ {
		edge := d.edges[edgeIndex]
		dsu.Union(edge.from, edge.to)
		edgeIndex++
	}

	sizes := dsu.GetComponentSizes()

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	for i := 0; i < int(math.Min(3, float64(len(sizes)))); i++ {
		result *= sizes[i]
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	if len(d.points) == 0 {
		return 0
	}

	dsu := tools.NewDSU(len(d.points))

	var lastUnionEdge Edge

	for _, edge := range d.edges {
		if dsu.Union(edge.from, edge.to) {
			lastUnionEdge = edge

			if dsu.CountComponents() == 1 {
				break
			}
		}
	}

	result = d.points[lastUnionEdge.from].X * d.points[lastUnionEdge.to].X

	return result
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
