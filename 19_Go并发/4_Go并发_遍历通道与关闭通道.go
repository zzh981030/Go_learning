package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1 //初始化x,y值
	for i := 0; i < n; i++ {
		c <- x        //把每次迭代的值传进通道
		x, y = y, x+y //本次循环的结果作为下一次循环的初始值
	}
	close(c)
}

func main() {
	c := make(chan int, 10) //创建一个包含十个数据的通道
	go fibonacci(cap(c), c) //cap(c)表示通道c的容量
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。

	//运行完上述函数，结果已保存到通道中，
	for i := range c {
		fmt.Println(i)
	}
}
