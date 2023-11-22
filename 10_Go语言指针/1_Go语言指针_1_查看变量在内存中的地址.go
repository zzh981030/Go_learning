// 显示变量在内存中的地址
package main

import "fmt"

func main() {
	var a int = 10

	fmt.Printf("变量的地址: %b\n", &a) //二进制显示
	fmt.Printf("变量的地址: %o\n", &a) //八进制显示
	fmt.Printf("变量的地址: %d\n", &a) //十进制显示
	fmt.Printf("变量的地址: %x\n", &a) //十六进制显示
}
