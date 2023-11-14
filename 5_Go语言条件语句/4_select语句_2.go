/*
以下实例中，我们定义了两个通道，并启动了两个协程（Goroutine）从这两个通道中获取数据。
在 main 函数中，我们使用 select 语句在这两个通道中进行非阻塞的选择，如果两个通道都没有可用的数据，就执行 default 子句中的语句。

以下实例执行后会不断地从两个通道中获取到的数据，当两个通道都没有可用的数据时，会输出 "no message received"。
*/

package main

import "fmt"

func main() {
	// 定义两个通道
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个 goroutine，分别从两个通道中获取数据
	go func() {
		for {
			ch1 <- "from 1"
		}
	}()
	go func() {
		for {
			ch2 <- "from 2"
		}
	}()

	// 使用 select 语句非阻塞地从两个通道中获取数据
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		default:
			// 如果两个通道都没有可用的数据，则执行这里的语句
			fmt.Println("no message received")
		}
	}
}
