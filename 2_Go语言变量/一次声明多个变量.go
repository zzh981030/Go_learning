//声明变量的一般形式是使用 var 关键字：
//var identifier type

//可以一次声明多个变量：
//var identifier1, identifier2 type

package main

import "fmt"

func main() {
	var a string = "Runoob" //一次定义一个变量
	fmt.Println(a)

	var b, c int = 1, 2 //一次定义多个变量
	fmt.Println(b, c)

	var d int //一次定义多个变量
	fmt.Println(d)
}
