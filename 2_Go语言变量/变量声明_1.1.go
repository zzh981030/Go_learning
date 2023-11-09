//变量声明
//第一种，指定变量类型，如果没有初始化，则变量默认为零值。
//var v_name v_type
//v_name = value

package main

import "fmt"

func main() {

	// 声明一个变量并初始化
	var a = "RUNOOB"
	fmt.Println(a)

	// 没有初始化就为零值
	var b int
	fmt.Println(b)

	// bool 零值为 false
	var c bool
	fmt.Println(c)
}
