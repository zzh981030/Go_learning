package main

import (
	"fmt"
)

func main() {
	var i, j, k int
	// 声明数组的同时快速初始化数组
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0} //定义不定长数组
	/* 输出每个数组元素的值 */
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0} //定义一个长度为5的数组，默认全部元素为0；并指定索引为1和3的9元素的值
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}

	//for key_1, value_1 := range balance {
	//	fmt.Printf("第%d个元素是 %f\n", key_1, value_1)
	//}
}
