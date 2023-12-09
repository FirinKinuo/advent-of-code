package main

import (
	"fmt"
	"github.com/FirinKinuo/advent-of-code"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type Day struct {
	*problem.DayTemplate

	seeds       []int
	seedsRanges [][]int
	maps        [][][]int
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2023", "day05_if_you_give_a_seed_a_fertilizer", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %s", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	for i, s := range strings.Split(input, "\r\n\r\n") {
		switch i {
		case 0:
			d.seeds = d.parseValues(s)
		default:
			d.maps = append(d.maps, d.parseMap(s))
		}
	}

	for i, seed := range d.seeds {
		if i%2 == 0 {
			d.seedsRanges = append(d.seedsRanges, []int{seed})
		} else {
			d.seedsRanges[i/2] = append(d.seedsRanges[i/2], seed)
		}
	}
}

func (d *Day) parseValues(values string) []int {
	var seeds []int
	for _, s := range strings.Split(values, " ") {
		convertedValue, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		seeds = append(seeds, convertedValue)
	}
	return seeds
}

func (d *Day) parseMap(input string) [][]int {
	var mapSlice [][]int
	for _, s := range strings.Split(input, "\r\n") {
		parsedValues := d.parseValues(s)
		if parsedValues != nil {
			mapSlice = append(mapSlice, parsedValues)
		}
	}
	return mapSlice
}

func (d *Day) correspond(src int, maps [][]int) int {
	value := src
	for _, m := range maps {
		if src >= m[1] && src < m[1]+m[2] {
			value = src - m[1] + m[0]
		}
	}

	return value
}

func (d *Day) findLocations() []int {
	locations := make([]int, 0, len(d.seeds))
	seedPathMap := map[int][]int{}
	for _, seed := range d.seeds {
		seedPathMap[seed] = append(seedPathMap[seed], d.correspond(seed, d.maps[0]))

		for i := 1; i < len(d.maps); i++ {
			seedPathMap[seed] = append(seedPathMap[seed], d.correspond(seedPathMap[seed][len(seedPathMap[seed])-1], d.maps[i]))
		}
		locations = append(locations, seedPathMap[seed][len(seedPathMap[seed])-1])
	}

	return locations
}

func (d *Day) FirstProblem() int {
	locations := d.findLocations()

	return slices.Min(locations)
}

func (d *Day) findMinLocationsInRanges() int {
	minLocation := math.MaxInt

	wg := &sync.WaitGroup{}
	wg.Add(len(d.seedsRanges))

	for _, seedsRange := range d.seedsRanges {
		go func(seedsRange []int) {
			for seed := seedsRange[0]; seed < seedsRange[0]+seedsRange[1]; seed++ {
				seedPaths := make([]int, 0, len(d.maps))
				seedPaths = append(seedPaths, d.correspond(seed, d.maps[0]))

				for i := 1; i < len(d.maps); i++ {
					seedPaths = append(seedPaths, d.correspond(seedPaths[len(seedPaths)-1], d.maps[i]))
				}

				if seedPaths[len(seedPaths)-1] < minLocation {
					minLocation = seedPaths[len(seedPaths)-1]
				}
			}
			wg.Done()
		}(seedsRange)
	}

	wg.Wait()

	return minLocation
}

func (d *Day) SecondProblem() int {
	return d.findMinLocationsInRanges()
}

func main() {
	day, err := NewDay(problem.ProblemInput)
	if err != nil {
		log.Fatalf("new day: %s", err)
	}

	day.Problem.Solve(day)
}
