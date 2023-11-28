/*
以下实例通过 Go 语言的递归函数实例阶乘：
*/

package main

import "fmt"

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		fmt.Println(result)
		return result
	}
	return 1
}

func main() {
	var i int = 5
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}
