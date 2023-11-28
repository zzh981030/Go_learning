//以下实例将字符串转换为整数

package main

import (
	"fmt"
	"strconv" //strconv.Atoi 函数返回两个值，第一个是转换后的整型值，第二个是可能发生的错误，我们可以使用空白标识符 _ 来忽略这个错误
)

func main() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
	}
}
