package simplechartTypes

import (
	"fmt"
	"testing"
)

// 普通测试函数，包含具体的测试逻辑
func TestChartCreation(t *testing.T) {
	filePath := "D:\\1code\\DesignPatterns\\Design-patterns-for-various-programming-languages\\1create\\1ThreeFactories\\config.txt"
	ans := ReadFileCharByChar(filePath)
	if ans.err != nil {
		t.Fatalf("读取文件出错: %v", ans.err)
	}
	info := ans.info
	for _, chartType := range info {
		t.Logf("Chart In this turn : %s \t\n", chartType)
		chart := creatChart(chartType)
		if chart == nil {
			fmt.Println("Chart error!")
		}
	}
	t.Log("graphic test over")
}
