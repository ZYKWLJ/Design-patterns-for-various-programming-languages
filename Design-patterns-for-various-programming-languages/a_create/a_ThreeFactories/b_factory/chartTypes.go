package Factory

import (
	"fmt"
)

type chartInterface interface {
	Display()
}
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

// 假设新增一个类型，直接添加这一处就好了~~很简洁
// (注意，添加代码不同于传统意义上的破坏修改性，还是对开闭原则友好的！因为没有打破原有代码逻辑，只是增加)
type square struct {
	Chart
}

func (c *Chart) Display() {
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
