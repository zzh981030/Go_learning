/*
以下实例我们定义了多个匿名函数，并展示了如何将匿名函数赋值给变量、在函数内部使用匿名函数以及将匿名函数作为参数传递给其他函数。
*/

package main

import (
	"fmt"
)

func main() {
	/* demo——1 */
	// 定义一个匿名函数并将其赋值给变量add
	add := func(a, b int) int {
		return a + b
	}

	// 调用匿名函数
	result := add(3, 5)
	fmt.Println("3 + 5 =", result)

	/* demo——2 */
	// 在函数内部使用匿名函数
	multiply := func(x, y int) int {
		return x * y
	}

	product := multiply(4, 6)
	fmt.Println("4 * 6 =", product)

	/* demo——3.1 */
	// 将匿名函数作为参数传递给其他函数
	calculate := func(operation func(int, int) int, x, y int) int {
		return operation(x, y)
	}

	sum := calculate(add, 2, 8) //实例化calculate对象后，传入参数add和x,y；calculate中的operation操作即为add。且calculate会将后边的两个函数传入到add函数中
	fmt.Println("2 + 8 =", sum)

	/* demo——3.2 */
	// 也可以直接在函数调用中定义匿名函数
	difference := calculate(func(a, b int) int {
		return a - b
	}, 10, 4)
	fmt.Println("10 - 4 =", difference)
}
