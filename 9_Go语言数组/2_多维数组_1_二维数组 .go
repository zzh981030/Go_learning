/*二维数组
二维数组是最简单的多维数组，二维数组本质上是由一维数组组成的。二维数组定义方式如下：
var arrayName [ x ][ y ] variable_type

variable_type 为 Go 语言的数据类型，arrayName 为数组名，二维数组可认为是一个表格，x 为行，y 为列，下图演示了一个二维数组 a 为三行四列：
*/

//二维数组中的元素可通过 a[ i ][ j ] 来访问。

package main

import "fmt"

func main() {
	// Step 1: 创建数组
	values := [][]int{} //初始化二维数组，定义一个空二维数组

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3} //定义一维数组row1
	row2 := []int{4, 5, 6} //定义一维数组row2
	values = append(values, row1)
	values = append(values, row2)

	//Step 3: 显示两行数据
	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	// Step 4: 访问第一个元素
	fmt.Println("第一个元素为：")
	fmt.Println(len(values))
	fmt.Println(values[0][0])
}
