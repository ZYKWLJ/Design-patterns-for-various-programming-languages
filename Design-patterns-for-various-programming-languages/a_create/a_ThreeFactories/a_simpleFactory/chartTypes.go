package a_simpleFactory

import (
	"fmt"
)

type chartDisplay interface {
	Display()
}
type Chart struct {
	name string
}

func (c *Chart) Display() {
	fmt.Println("display Chart...display", c.name)
}

type pie struct {
	Chart
}

type histogram struct {
	Chart
}
type line struct {
	Chart
}

func newPie(name string) *Chart {
	fmt.Println("Initiated a new pie")
	return &Chart{name: name}
}

func newHistogram(name string) *Chart {
	fmt.Println("Initiated a new histogram")
	return &Chart{name: name}
}
func newLine(name string) *Chart {
	fmt.Println("Initiated a new line")
	return &Chart{name: name}
}
