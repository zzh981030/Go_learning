/*
在 Go 语言中，break 语句在 select 语句中的应用是相对特殊的。
由于 select 语句的特性，break 语句并不能直接用于跳出 select 语句本身，因为 select 语句是非阻塞的，它会一直等待所有的通信操作都准备就绪。
如果需要提前结束 select 语句的执行，可以使用 return 或者 goto 语句来达到相同的效果。

以下实例，展示了在 select 语句中使用 return 来提前结束执行的情况：
实例
*/
package main

import (
	"fmt"
	"time"
)

func process(ch chan int) { //将channel中的值作为参数导入
	for {
		select {
		case val := <-ch: //使用channel中穿的参数值
			fmt.Println("Received value:", val)
			// 执行一些逻辑，即通道传进来的值符合下列条件时，执行return语句
			if val == 1 {
				return // 提前结束 select 语句的执行
			}
		default:
			fmt.Println("No value received yet.")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ch := make(chan int)

	go process(ch)

	time.Sleep(2 * time.Second)
	ch <- 1 //将变量的值传输到chanel中
	time.Sleep(1 * time.Second)
	ch <- 3
	time.Sleep(1 * time.Second)
	ch <- 5
	time.Sleep(1 * time.Second)
	ch <- 7

	time.Sleep(2 * time.Second)
}

/*
以上实例中，process 函数在一个无限循环中使用 select 语句等待通道 ch 上的数据。当接收到数据时，会执行一些逻辑。
当接收到的值等于 5 时，使用 return 提前结束 select 语句的执行。
*/
