package main

import (
	"container/list"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

type Day struct {
	*problem.DayTemplate

	rules map[string]struct{}
	pages [][]int
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2024", "day05_print_queue", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template, rules: make(map[string]struct{})}, nil
}

func (d *Day) PrepareInput(input string) {
	parts := strings.Split(input, "\n\n")

	for _, rules := range strings.Split(parts[0], "\n") {
		rulesSlice := tools.AtoiSlice(strings.Split(rules, "|"))
		d.rules[strconv.Itoa(rulesSlice[0])+":"+strconv.Itoa(rulesSlice[1])] = struct{}{}
	}

	for _, page := range strings.Split(parts[1], "\n") {
		d.pages = append(d.pages, tools.AtoiSlice(strings.Split(page, ",")))
	}
}

func (d *Day) check(page []int) bool {
	for i := 0; i < len(page)-1; i++ {
		_, ok := d.rules[strconv.Itoa(page[i])+":"+strconv.Itoa(page[i+1])]
		if !ok {
			return false
		}
	}

	return true
}

func (d *Day) fixPage(page []int) []int {
	graph := make(map[int][]int)
	indegree := make(map[int]int)

	for rule := range d.rules {
		first := tools.MustAtoi(rule[0:2])
		second := tools.MustAtoi(rule[3:5])
		if slices.Contains(page, first) && slices.Contains(page, second) {
			graph[first] = append(graph[first], second)
			indegree[second]++

			if _, ok := indegree[first]; !ok {
				indegree[first] = 0
			}
			if _, ok := indegree[second]; !ok {
				indegree[second] = 0
			}
		}
	}

	queue := list.New()
	for number, count := range indegree {
		if count == 0 {
			queue.PushBack(number)
		}
	}

	var fixedPage []int
	for queue.Len() > 0 {
		number := queue.Remove(queue.Front()).(int)
		fixedPage = append(fixedPage, number)

		for _, neighbor := range graph[number] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue.PushBack(neighbor)
			}
		}
	}

	return fixedPage
}

func (d *Day) FirstProblem() int {
	var result int

	for _, page := range d.pages {
		if d.check(page) {
			result += page[len(page)/2]
		}
	}

	return result
}

func (d *Day) SecondProblem() int {
	var result int

	for _, page := range d.pages {
		if !d.check(page) {
			fixed := d.fixPage(page)
			result += fixed[len(fixed)/2]
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
