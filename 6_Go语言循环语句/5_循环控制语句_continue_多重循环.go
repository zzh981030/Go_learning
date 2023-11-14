/*
以下实例有多重循环，演示了使用标记和不使用标记的区别：
*/

package main

import "fmt"

func main() {

	// 不使用标记
	fmt.Println("---- continue ---- ")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue
		}
	}

	// 使用标记
	fmt.Println("---- continue label ----")
end_mark:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		continue end_mark //跳到标记点的后方

		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			//continue end_mark //跳到标记点的后方
		}
	}
}
