/*
可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]，实例如下：
*/

//注意：数组切片时为<左闭右开>

package main

import "fmt"

func main() {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4]) //左闭右开区间，即第1，2，3号元素

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3]) //左闭右开区间，即第0，，1，2号元素

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:]) //左闭右开区间，即第4号及之后所有元素

	numbers1 := make([]int, 0, 5) //创建一个长度为0，切片数为5的数组
	printSlice_3(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2] //子切片包含0，1号元素，长度为2，但容量仍然和原数组一致
	printSlice_3(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5] //子切片包含2，3，4号元素，长度为3，容量怎么只有7了？？？？
	printSlice_3(number3)

}

// 定义子函数，显示数组的<长度，容量和数组内容>
func printSlice_3(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
