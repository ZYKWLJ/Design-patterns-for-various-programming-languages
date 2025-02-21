package main

import "fmt"

func Chart(chartTypes string) *chart {
	if chartTypes == "pie" {
		return newPie(chartTypes)
	} else if chartTypes == "histogram" {
		return newHistogram(chartTypes)
	} else if chartTypes == "line" {
		return newLine(chartTypes)
	}
	fmt.Println("There is no current chart in the graphics library")
	return nil
}
