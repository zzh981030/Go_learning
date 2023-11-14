// 以下实例有多重循环，演示了使用标记和不使用标记的区别：
package main

import "fmt"

func main() {

	// 不使用标记
	fmt.Println("---- break ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
	}

	// 使用标记
	fmt.Println("---- break label ----")
end_mark: //这是一个标记的变量
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			//break end_mark //结束当前代码块，并跳到标记点的后方
		}
		break end_mark //结束当前代码块，并跳到标记点的后方
	}
}
