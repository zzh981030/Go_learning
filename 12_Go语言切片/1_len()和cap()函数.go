/*
len() 和 cap() 函数

切片是可索引的，并且可以由 len() 方法获取长度。

切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
*/
package main

import (
	"fmt"
)

func main() {

	var numbers = make([]int, 3, 5) //定义数组时指定长度和容量(切片长度)

	printSlice_1(numbers)
}

// 定义子函数，显示数组的<长度，容量和数组内容>
func printSlice_1(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
