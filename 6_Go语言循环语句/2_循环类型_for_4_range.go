/*
For-each range 循环
这种格式的循环可以对字符串、数组、切片等进行迭代输出元素。
*/
package main

import "fmt"

func main() {
	strings := []string{"google", "runoob"}
	for i, s := range strings { //i,s分别对应字符串中元素的位置和值
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5} //numbers为一个长度6的空数组，只指定了前四个元素，因此为{1，2，3，5，0，0}。
	for i, x := range numbers {   //i,x分别对应数组中元素的位置和值
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}
