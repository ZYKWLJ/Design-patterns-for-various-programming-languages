package a_simpleFactory

import "fmt"

type Chart struct {
	name string
}

// 为父类实现的方法(go中很有意思就是结构只能是方法的集合，不能有字段，那么我们就为父类提供方法就好了)
func (c *Chart) display() {
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
