package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)     //make函数创建一个动态内存
	go sum(s[:len(s)/2], c) //c=17
	go sum(s[len(s)/2:], c) //c=-5
	x, y := <-c, <-c        // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}
