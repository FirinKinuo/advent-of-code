package problem

import (
	"fmt"
	"path"
)

type DayTemplate struct {
	year    string
	name    string
	Problem *Problem
}

func NewDayTemplate(year string, name string, inputType InputType) (*DayTemplate, error) {
	day := &DayTemplate{
		year: year,
		name: name,
	}
	p, err := NewProblem(path.Join(year, day.name, string(inputType)))
	if err != nil {
		return nil, fmt.Errorf("new problem init: %s", err)
	}

	day.Problem = p

	return day, nil
}

func (d *DayTemplate) DayName() string {
	return d.name
}

func (d *DayTemplate) Year() string {
	return d.year
}
