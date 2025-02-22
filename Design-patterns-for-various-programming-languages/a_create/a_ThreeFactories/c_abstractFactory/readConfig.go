package Factory

//package main

import (
	"bufio"
	"fmt"
	"os"
)

type ret struct {
	info [][]string //新增加一个维度
	err  error
}

// ReadFileCharByChar 逐个字符读取文件内容
func ReadFileCharByChar(filePath string) ret {
	// 打开文件
	var ans ret
	file, err := os.Open(filePath)
	if err != nil {
		return ret{nil, fmt.Errorf("打开文件出错: %w", err)}
	}
	// 确保文件在函数结束时关闭
	defer file.Close()

	// 创建一个带缓冲的读取器
	reader := bufio.NewReader(file)

	//str := ""
	var str []string
	tempStr := ""
	// 逐个字符读取文件内容
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			// 如果文件结束，且当前行有内容，添加到结果切片
			if len(str) > 0 {
				ans.info = append(ans.info, str)
			}
			//fmt.Println(ans)
			return ans
		}
		if char == '\n' || char == '\r' { // 处理换行符和回车符
			// 忽略空行
			if len(str) > 0 {
				str = append(str, tempStr)
				tempStr = ""
				ans.info = append(ans.info, str)
			}
			str = []string{}
		} else {
			if char == ' ' {
				str = append(str, tempStr)
				tempStr = ""
			} else {
				// 拼接字符
				tempStr += string(char)
			}
		}
	}
	return ret{nil, nil}
}

//// 测试
//func main() {
//	filePathTheChartWeHave := "D:\\1code\\DesignPatterns\\Design-patterns-for-various-programming-languages\\a_create\\a_ThreeFactories\\config.txt"
//	ans := ReadFileCharByChar(filePathTheChartWeHave)
//	strs := ans.info
//	for _, str := range strs {
//		fmt.Println(str, " len:", len(str[0]))
//	}
//	fmt.Println(ans)
//}
