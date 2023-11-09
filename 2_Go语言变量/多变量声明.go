// //类型相同多个变量, 非全局变量
// var vname1, vname2, vname3 type
// vname1, vname2, vname3 = v1, v2, v3
//
// var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断
// vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误
//
// // 这种因式分解关键字的写法一般用于声明《全局变量》
// var (
//	vname1 v_type1
//	vname2 v_type2
// )

// 《局部变量》声明后必须使用。否则会得到编译错误
// 《全局变量》允许声明但不使用。 同一类型的多个变量可以声明在同一行
// 空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。
package main

import "fmt"

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main() {
	g, h := 123, "hello"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)
}
