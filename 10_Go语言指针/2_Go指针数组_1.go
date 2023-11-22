/*在我们了解指针数组前，先看个实例，定义了长度为 3 的整型数组：
 */

package main

import "fmt"

const MAX1 int = 3

func main() {

	a := []int{10, 100, 200}
	var i int

	for i = 0; i < MAX1; i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
}
