package main

import "fmt"

func main() {
	//var str string
	filePath := "D:\\1code\\各种语言的设计模式\\Design-patterns-for-various-programming-languages\\1创建型模式\\1工厂三兄弟\\config.txt"
	ans := ReadFileCharByChar(filePath)
	info := ans.info
	for _, charts := range info {
		fmt.Printf("chart In this turn : %s \t\n", charts)
		chart := Chart(charts)
		if chart != nil {
			chart.display()
		} else {
			fmt.Println("chart error!")
		}
	}
	fmt.Println("graphic test over")
}
