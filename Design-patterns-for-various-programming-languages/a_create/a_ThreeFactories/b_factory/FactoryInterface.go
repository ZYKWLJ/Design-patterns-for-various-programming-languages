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
func (factory *FactoryInterface) createChart() *Chart {
	fmt.Println("Initiated a new ", factory.chartType)
	return &Chart{factory.chartType}
}
