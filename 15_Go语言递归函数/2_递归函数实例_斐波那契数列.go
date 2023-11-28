/*以下实例通过 Go 语言的递归函数实现斐波那契数列：
 */

package main

import "fmt"

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		result := fibonacci(i)
		//fmt.Printf("%d\t", fibonacci(i))
		fmt.Printf("%d\t", result)
	}
}
