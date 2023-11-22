//https://blog.csdn.net/gudongkun1121/article/details/131697332

package main

import "fmt"

func main() {
	//整数占位符
	v1 := 10
	v2 := 20170 //今字的码点值
	fmt.Printf("二进制:%b \n", v1)
	fmt.Printf("uincode码点值转字符:%c \n", v2)
	fmt.Printf("十进制:%d \n", v1)
	fmt.Printf("八进制:%o \n", v1)
	fmt.Printf("0o为前缀的八进制: %O \n", v1)
	fmt.Printf("用单引号把字面值包起来: %q \n", v2)
	fmt.Printf("十六进制: %x \n", v1)
	fmt.Printf("大写十六进制: %X \n", v1)
	fmt.Printf("Unicode格式: %U \n", v2)
	//宽度设置
	fmt.Printf("%v的二进制表示为：%b;go 语法表示为：%#b;指定二进制的宽度为8位，不足的补0：%08b \n", v1, v1, v1, v1)
	fmt.Printf("%v的十六进制表示为:%x; go语法表示，定宽度为8，不足的补0：%#08x \n", v1, v1, v1)
	fmt.Printf("%v的字符为：%c；指定宽度为5位，不足5位的补空格：%5c\n", v2, v2, v2)

}
