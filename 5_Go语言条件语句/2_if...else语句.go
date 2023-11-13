/*
if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。

语法
Go 编程语言中 if...else 语句的语法如下：

	if 布尔表达式 {
		在布尔表达式为 true 时执行
	} else {

		 在布尔表达式为 false 时执行
	}
*/
package main

import "fmt"

func main() {
	/* 局部变量定义 */
	var a int = 100
	/* 判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	} else {
		/* 如果条件为 false 则执行以下语句 */
		fmt.Printf("a 不小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)

}
