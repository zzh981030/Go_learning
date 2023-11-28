/*
以下实例中将整型转化为浮点型，并计算结果，将结果赋值给浮点型变量：
*/

package main

import "fmt"

func main() {
	var sum uint32 = 17.00000
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)
}
