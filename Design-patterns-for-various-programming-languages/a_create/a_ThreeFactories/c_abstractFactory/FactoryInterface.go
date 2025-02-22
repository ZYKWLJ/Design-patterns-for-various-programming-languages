package Factory

import "fmt"

// 定义工厂接口=>作用：解耦子类的方法实现
type FactoryInterface struct {
	chartType string
	size      string
}

func newFactoryInterface(chartType string, size string) *FactoryInterface {
	return &FactoryInterface{chartType: chartType, size: size}
}

const filePathTheChartWeHave = "D:\\1code\\DesignPatterns\\Design-patterns-for-various-programming-languages\\a_create\\a_ThreeFactories\\c_abstractFactory\\charts.txt"

func (factory *FactoryInterface) createChart(set map[string]struct{}) *Chart {
	//set := chartWeHave(filePathTheChartWeHave)
	if _, ok := set[factory.chartType+factory.size]; ok {
		fmt.Println("Initiated a new ", factory.chartType)
		return &Chart{name: factory.chartType, size: factory.size}
	} else {
		fmt.Println("There is no current Chart in the graphics library")
		return &Chart{name: "nil", size: "nil"}
	}
}

func chartWeHave(filePath string) map[string]struct{} {
	ans := ReadFileCharByChar(filePath)
	set := make(map[string]struct{})
	for _, str := range ans.info {
		temp := str[0] + str[1]
		//fmt.Println("temp=", temp)
		set[temp] = struct{}{}
		//fmt.Println("我们本来提供的的类型", str)
	}
	return set
}
