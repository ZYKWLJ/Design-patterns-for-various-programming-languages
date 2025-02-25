package a_simpleFactory

import "fmt"

// 优点：简单工厂模式的核心在将创建各种实例的机会全部封装起来了，统一使用一个接口!
// 缺点：不具备开闭原则(对修改关闭，对扩展开放)，扩展起来需要修改if-else代码，带来不便！
func creatChart(chartTypes string) *Chart {
	if chartTypes == "pie" {
		return newPie(chartTypes)
	} else if chartTypes == "histogram" {
		return newHistogram(chartTypes)
	} else if chartTypes == "line" {
		return newLine(chartTypes)
	}
	fmt.Println("There is no current Chart in the graphics library")
	return &Chart{name: "nil"}
}
