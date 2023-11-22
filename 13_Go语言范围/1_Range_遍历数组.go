// 遍历简单的数组，2**%d 的结果为 2 对应的次方数：

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
