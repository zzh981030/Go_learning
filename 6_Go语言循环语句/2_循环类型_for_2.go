/*
init 和 post 参数是可选的，我们可以直接省略它，类似 While 语句。
以下实例在 sum 小于 10 的时候计算 sum 自相加后的值：
*/

package main

import "fmt"

func main() {
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	// 这样写也可以，更像 While 语句形式
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)
}
