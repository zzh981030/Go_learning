//函数作为另一个函数的实参——即函数嵌套函数，一个函数引用另一个函数
/*Go 语言可以很灵活的创建函数，并作为另外一个函数的实参。以下实例中我们在定义的函数中初始化一个变量，该函数仅仅是为了使用内置函数 math.sqrt()，实例为：
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	/* 声明函数变量 */
	//定义一个函数，这个函数仅仅是为了使用内置函数math.Sqrt()
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(9)) //使用自定义函数求开方
}
