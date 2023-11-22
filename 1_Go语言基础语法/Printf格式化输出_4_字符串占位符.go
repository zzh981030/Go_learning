//https://blog.csdn.net/gudongkun1121/article/details/131697332

package main

import "fmt"

func main() {
	var str = "今天是个好日子"
	fmt.Printf("描述一下今天：%s \n", str)
	fmt.Printf("描述一下今天(带引号)：%q \n", str)
	fmt.Printf("字符串的十六进制表示，每两个字母为一个byte：%x \n", str)
	fmt.Printf("字符串的十六进制表示，空格作为分隔符：% x \n", str)

}
