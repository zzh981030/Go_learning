/*
以下实例将整数转换为字符串：
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 123
	str := strconv.Itoa(num)
	fmt.Printf("整数 %d  转换为字符串为：'%s'\n", num, str)
}
