package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/FirinKinuo/advent-of-code"
	"github.com/FirinKinuo/advent-of-code/tools"
)

type Day struct {
	*problem.DayTemplate
	diskMap []int
}

func NewDay(inputType problem.InputType) (*Day, error) {
	template, err := problem.NewDayTemplate("2024", "day09_disk_fragmenter", inputType)
	if err != nil {
		return nil, fmt.Errorf("new day template: %w", err)
	}
	return &Day{DayTemplate: template}, nil
}

func (d *Day) PrepareInput(input string) {
	d.diskMap = tools.AtoiSlice(strings.Split(input, ""))

}

func (d *Day) FirstProblem() int {
	var result int

	var blocks []int
	blockId := 0
	for i, data := range d.diskMap {
		if i%2 == 0 {
			for n := 0; n < data; n++ {
				blocks = append(blocks, blockId)
			}
			blockId++
		} else {
			for n := 0; n < data; n++ {
				blocks = append(blocks, -1)
			}
		}
	}

	for i := len(blocks) - 1; i != 0; i-- {
		freeSpaceIndex := slices.Index(blocks, -1)
		if freeSpaceIndex == -1 {
			break
		}
		blocks[freeSpaceIndex] = blocks[i]
		blocks = blocks[:i]
	}

	for i, block := range blocks {
		result += block * i
	}
	return result
}

func (d *Day) findFreeSpaceIndex(blocks []int, insertSpace int) int {
	for i := 0; i < len(blocks); i++ {
		freeSpaceIndex := slices.IndexFunc(blocks[i:], func(i int) bool {
			return i < 1000 && i >= 100+insertSpace
		})

		if freeSpaceIndex != -1 {
			return freeSpaceIndex
		}
	}
	return -1
}

func (d *Day) SecondProblem() int {
	var result int
	blocks := make([]int, 0, len(d.diskMap)*2)
	blockId := 0

	for i, data := range d.diskMap {
		if i%2 == 0 {
			blocks = append(blocks, 1_000_000+data*100_000+blockId)
			blockId++
		} else {
			blocks = append(blocks, 100+data)
		}
	}

	for i := len(blocks) - 1; i >= 1; i-- {
		if blocks[i]%100_000 == blocks[i] {
			continue
		}

		count := (blocks[i] / 100_000) % 10
		freeSpaceIndex := d.findFreeSpaceIndex(blocks, count)

		if freeSpaceIndex == -1 || freeSpaceIndex > i {
			continue
		}

		freeSpaceCount := blocks[freeSpaceIndex] % 10
		blocks[freeSpaceIndex] = 100 + (freeSpaceCount - count)
		blocks = append(blocks[:freeSpaceIndex], append([]int{blocks[i]}, blocks[freeSpaceIndex:]...)...)
		blocks[i+1] = 100 + count
	}

	var formatedData []int
	for _, block := range blocks {
		var count, id int
		if block%100_000 == block {
			count = block % 100
		} else {
			count = (block / 100_000) % 10
			id = block % 100_000
		}

		for j := 0; j < count; j++ {
			formatedData = append(formatedData, id)
		}
	}

	for i, value := range formatedData {
		result += i * value
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
