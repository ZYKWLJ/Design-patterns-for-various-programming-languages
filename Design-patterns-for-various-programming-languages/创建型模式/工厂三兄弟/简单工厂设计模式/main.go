package main

import "fmt"

func main() {
	var str string
	for {
		fmt.Print("请输入图表类型 (pie, histogram, line): ")
		fmt.Scan(&str)
		chart := Chart(str)
		if chart != nil {
			chart.display()
		} else {
			fmt.Println("chart error!")
		}
	}
}
