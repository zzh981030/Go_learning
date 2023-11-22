/*
二维数组通过指定坐标来访问。如数组中的行索引与列索引，例如：

val := a[2][3]
或
var value int = a[2][3]
*/

// 二维数组可以使用循环嵌套来输出元素：
package main

import "fmt"

func main() {
	/* 数组 - 5 行 2 列*/
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
