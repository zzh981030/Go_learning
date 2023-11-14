/*在 select 语句中使用 break：
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Millisecond)
		ch1 <- 0
	}()

	go func() {
		time.Sleep(1 * time.Millisecond)
		ch2 <- 1
	}()

	select { //随机选一个
	case <-ch1:
		fmt.Println("Received from ch1.")
	case <-ch2:
		fmt.Println("Received from ch2.")
		break // 跳出 select 语句
	}
}
