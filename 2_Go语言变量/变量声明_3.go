//第三种，如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误，格式：
//var intVal int
//intVal :=1  这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明

//直接使用<intVal := 1>即可

//intVal := 1 相等于：
//var intVal int
//intVal =1

// 可以将 var f string = "Runoob" 简写为 f := "Runoob"
package main

import "fmt"

func main() {
	f := "Runoob" // var f string = "Runoob"

	fmt.Println(f)
}
