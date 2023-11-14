/*
Go 函数可以返回多个值，例如：
*/

package main

import "fmt"

func swap_test(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap_test("Google", "Runoob")
	fmt.Println(a, b)
}
