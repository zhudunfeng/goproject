package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//定义一个结构体,用于保存统计结果
type CharCount struct {
	ChCount    int //记录英文个数
	NumCount   int //记录数字个数
	SpaceCount int //记录空格个数
	OtherCount int //记录其他字符个数
}

func main() {
	//思路:打开一个文件,创一个Reader
	//每读取一行，就去统计该行有多少个英文、数字、空格和其他字符
	//然后将结果保存到一个结构体
	file, err := os.Open("../test.txt")
	if err != nil {
		fmt.Println("open file error:", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var count CharCount
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//为了兼容中文字符。可以将str转成[]rune
		str1 := []rune(str)
		//每一行统计一次
		for _, v := range str1 {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			default:
				count.OtherCount++
			}
		}
	}

	fmt.Printf("英文字符%v 数字%v 空格%v 其他字符%v\n",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
