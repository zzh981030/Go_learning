//常量的定义格式：
//const identifier [type] = value

//你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。
//显式类型定义： const b string = "abc"
//隐式类型定义： const b = "abc"

//多个相同类型的声明可以简写为：
//const c_name1, c_name2 = value1, value2

package main

import "fmt"

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)
}
