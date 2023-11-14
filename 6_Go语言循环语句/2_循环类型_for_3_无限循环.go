//无限循环:

package main

import "fmt"

func main() {
	sum := 0
	for {
		sum++ // 无限循环下去
		fmt.Println(sum)
	}
	fmt.Println(sum) // 无法输出
}

//要停止无限循环，可以在命令窗口按下ctrl-c 。
