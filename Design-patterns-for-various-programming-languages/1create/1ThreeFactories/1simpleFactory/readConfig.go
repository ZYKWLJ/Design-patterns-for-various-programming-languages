package simplechartTypes

import (
	"bufio"
	"fmt"
	"os"
)

type ret struct {
	info []string
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

	str := ""
	// 逐个字符读取文件内容
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			// 如果文件结束，且当前行有内容，添加到结果切片
			if len(str) > 0 {
				ans.info = append(ans.info, str)
			}
			return ans
		}
		if char == '\n' || char == '\r' { // 处理换行符和回车符
			// 忽略空行
			if len(str) > 0 {
				ans.info = append(ans.info, str)
			}
			str = ""
		} else {
			// 拼接字符
			str += string(char)
		}
	}
	return ret{nil, nil}
}

// 测试
//func main() {
//	filePath := "D:\\1code\\DesignPatterns\\Design-patterns-for-various-programming-languages\\1create\\1ThreeFactories\\config.txt"
//	ans := ReadFileCharByChar(filePath)
//	strs := ans.info
//	for _, str := range strs {
//		fmt.Println(str, " len:", len(str))
//	}
//	fmt.Println(ans)
//}
