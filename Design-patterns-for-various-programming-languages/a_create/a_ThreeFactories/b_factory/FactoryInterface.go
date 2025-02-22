package Factory

import "fmt"

// 定义工厂接口=>作用：解耦子类的方法实现
type FactoryInterface struct {
	chartType string
}

func newFactoryInterface(chartType string) *FactoryInterface {
	return &FactoryInterface{chartType: chartType}
}

// 这个抽象方法统一创建对象
//
//	func (factory *FactoryInterface) createChart() *Chart {
//		fmt.Println("Initiated a new ", factory.chartType)
//		return &Chart{factory.chartType}
//	}
const filePathTheChartWeHave = "D:\\1code\\DesignPatterns\\Design-patterns-for-various-programming-languages\\a_create\\a_ThreeFactories\\b_factory\\charts.txt"

//set = chartWeHave(filePathTheChartWeHave)

func (factory *FactoryInterface) createChart(set map[string]struct{}) *Chart {
	//set := chartWeHave(filePathTheChartWeHave)
	if _, ok := set[factory.chartType]; ok {
		fmt.Println("Initiated a new ", factory.chartType)
		return &Chart{name: factory.chartType}
	} else {
		fmt.Println("There is no current Chart in the graphics library")
		return &Chart{name: "nil"}
	}
}

func chartWeHave(filePath string) map[string]struct{} {
	ans := ReadFileCharByChar(filePath)
	set := make(map[string]struct{})
	for _, str := range ans.info {
		set[str] = struct{}{}
		//fmt.Println("我们本来提供的的类型", str)
	}
	return set
}
