// 一个切片在未初始化之前默认为 nil，长度为 0，实例如下：

package main

import "fmt"

func main() {
	var numbers []int

	printSlice_2(numbers)

	if numbers != nil {
		fmt.Printf("切片不是空的")
	} else {
		fmt.Printf("切片 is 空的")
	}
}

// 定义子函数，显示数组的<长度，容量和数组内容>
func printSlice_2(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
