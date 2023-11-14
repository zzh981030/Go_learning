/*实例
计算 1 到 10 的数字之和：
*/

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println(sum)
}
