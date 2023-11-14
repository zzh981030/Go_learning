/*
在变量 a 大于 15 的时候跳出循环：
*/
package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 10

	/* for 循环 */
	for a < 20 {
		fmt.Printf("a 的值为 : %d\n", a)
		a++
		if a > 15 {
			/* a 大于 15 时使用 break 语句跳出循环 */
			break
		}
	}
}
