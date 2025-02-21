package Factory

import (
	"fmt"
)

type Chart struct {
	name string
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

func (c *Chart) display() {
	fmt.Println("display Chart...display", c.name)
}

// 上一个版本的这些重复度高的的代码既可以去除了！
//func newPie(name string) *Chart {
//	fmt.Println("Initiated a new pie")
//	return &Chart{name: name}
//}
//
//func newHistogram(name string) *Chart {
//	fmt.Println("Initiated a new histogram")
//	return &Chart{name: name}
//}
//func newLine(name string) *Chart {
//	fmt.Println("Initiated a new line")
//	return &Chart{name: name}
//}
