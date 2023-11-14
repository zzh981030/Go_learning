/*
当创建函数时，你定义了函数需要做什么，通过调用该函数来执行指定任务。
调用函数，向函数传递参数，并返回值，例如：
*/

package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max1(a, b)

	fmt.Printf("最大值是 : %d\n", ret)
}

/* 函数返回两个数的最大值 */
func max1(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
