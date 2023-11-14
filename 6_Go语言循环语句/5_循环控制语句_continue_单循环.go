/*
在变量 a 等于 15 的时候跳过本次循环执行下一次循环：
*/
package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 10

	/* for 循环 */
	for a < 20 {
		if a == 15 { //当满足条件时，执行此代码块的内容
			/* 跳出此代码块 */
			a = a + 1
			continue //遇到continue，直接跳过此代码块
			//break //遇到break，直接结束此运行程序
		}
		fmt.Printf("a 的值为 : %d\n", a)
		a++
	}
}
