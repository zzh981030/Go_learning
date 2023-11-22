//如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。

//当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址：
/*
指向指针的指针变量声明格式如下：

	var ptr **int;
以上指向指针的指针变量为整型。

访问指向指针的指针变量值需要使用两个 * 号，如下所示：
*/

package main

import "fmt"

func main() {

	var a int      //声明变量
	var ptr *int   //指向变量的指针，需要用到一个*声明
	var pptr **int //指向指针的指针，需要用到两个*声明

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}
