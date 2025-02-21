package main

import (
	"fmt"
)

type chart struct {
	name string
}
type pie struct {
	chart
}

type histogram struct {
	chart
}
type line struct {
	chart
}

func newPie(name string) *chart {
	fmt.Println("Initiated a new pie")
	return &chart{name: name}
}

func newHistogram(name string) *chart {
	fmt.Println("Initiated a new histogram")
	return &chart{name: name}
}
func newLine(name string) *chart {
	fmt.Println("Initiated a new line")
	return &chart{name: name}
}

func (c *chart) display() {
	fmt.Println("display chart...display", c.name)
}
