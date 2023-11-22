/*有一种情况，我们可能需要保存数组，这样我们就需要使用到指针。

以下声明了整型指针数组：

var ptr [MAX]*int;
ptr 为整型指针数组。因此每个元素都指向了一个值。以下实例的三个整数将存储在指针数组中：
*/

package main

import "fmt"

const MAX int = 3 //定义一个MAX的常量

func main() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int //声明整形指针数组

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i]) //上一步中将数组中元素的指针地址赋值给ptr，这里通过读取ptr存储的指针地址直接将数据显示出来
	}
}
