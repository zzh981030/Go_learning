/*
Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。实例如下：
*/
package main

import "fmt"

/* 声明全局变量 */
//这里不因该注释的，但是解释器一直报错，就先把这一句注释掉
//var g int = 20

func main() {
	/* 声明局部变量 */
	var g int = 10

	fmt.Printf("结果： g = %d\n", g)
}
