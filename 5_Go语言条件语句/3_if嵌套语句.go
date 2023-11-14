/*
你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
Go 编程语言中 if...else 语句的语法如下：
if 布尔表达式 1 {
	 在布尔表达式 1 为 true 时执行
	if 布尔表达式 2 {
		 在布尔表达式 2 为 true 时执行
	}
}
*/

package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			//fmt.Printf("a 的值为 100 ， b 的值为 200\n")
			fmt.Printf("a 的值为 %d ， b 的值为 %d\n", a, b)
		}
	}
	fmt.Printf("a 值为 : %d\n", a)
	fmt.Printf("b 值为 : %d\n", b)
}
