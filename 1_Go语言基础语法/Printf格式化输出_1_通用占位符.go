//https://blog.csdn.net/gudongkun1121/article/details/131697332

package main

import "fmt"

type simple struct {
	value int
}

func main() {
	var vv simple

	vvv := new(simple)

	// 通用占位符
	fmt.Println("-- 通用占位符 --")
	fmt.Printf("默认格式的值：%v \n", vv)
	fmt.Printf("包含字段名的值：%+v \n", vv)
	fmt.Printf("go语法表示的值：%#v \n", vv)
	fmt.Printf("go语法表示的类型：%T \n", vv)
	fmt.Printf("输出字面上的百分号：%%10 \n")

	// 通用占位符打印指针
	fmt.Println("-- 通用占位符打印指针 --")
	fmt.Printf("默认格式的值：%v \n", vvv)
	fmt.Printf("包含字段名的值：%+v \n", vvv)
	fmt.Printf("go语法表示的值：%#v \n", vvv)
	fmt.Printf("go语法表示的类型：%T \n", vvv)
	fmt.Printf("输出字面上的百分号：%%10 \n")

}
