package Factory

import (
	"fmt"
	"testing"
)

const filePathWeTryToAyanlze = "D:\\1code\\DesignPatterns\\Design-patterns-for-various-programming-languages\\a_create\\a_ThreeFactories\\config.txt"

// 普通测试函数，包含具体的测试逻辑
func TestChartCreation1(t *testing.T) {
	ans := ReadFileCharByChar(filePathWeTryToAyanlze)
	//fmt.Println("here passed1")
	set := chartWeHave(filePathTheChartWeHave)
	//fmt.Println("here passed2")
	if ans.err != nil {
		t.Fatalf("读取文件出错: %v", ans.err)
	}
	for _, chartTypeAndSize := range ans.info {
		fmt.Println(chartTypeAndSize)
		chartType, size := chartTypeAndSize[0], chartTypeAndSize[1]
		fmt.Println("in main", chartType, " ", size)
		t.Logf("Chart In this turn : type=%s ,size=%s\n", chartType, size)
		//核心改动点
		chart := newFactoryInterface(chartType, size)
		var Interface chartInterface = chart.createChart(set) //use interface to invoke Method！
		Interface.Display()
		if chart == nil {
			fmt.Println("Chart error!")
		}
		fmt.Println()
	}
	t.Log("graphic test over")
}
