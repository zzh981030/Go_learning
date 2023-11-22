/*
如果你想要在函数内修改原始数组，可以通过传递数组的指针来实现。

以下实例演示如何向函数传递数组，函数接受一个数组和数组的指针作为参数：

总结：函数中直接修改数组值，不会改变原数组中的值；函数中通过指针修改数组，会改变原数组中的值
*/

package main

import "fmt"

// 函数接受一个数组作为参数。功能：将数组中的所有值*2
func modifyArray(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2

	}
}

// 函数接受一个数组的指针作为参数。功能，将数组中的所有值*2
func modifyArrayWithPointer(arr *[5]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = (*arr)[i] * 2
	}
}

func main() {
	// 创建一个包含5个元素的整数数组
	myArray := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Original Array:", myArray)

	// 传递数组给函数，但不会修改原始数组的值
	modifyArray(myArray)
	//myArray_2 = modifyArray(myArray)
	fmt.Println("Array after modifyArray:", myArray)

	// 传递数组的指针给函数，可以修改原始数组的值
	modifyArrayWithPointer(&myArray)
	fmt.Println("Array after modifyArrayWithPointer:", myArray)
}
