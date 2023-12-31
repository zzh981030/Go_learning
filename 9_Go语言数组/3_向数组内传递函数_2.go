/*以上实例中我们使用的形参并未设定数组大小。

浮点数计算输出有一定的偏差，你也可以转整型来设置精度。
*/

package main

import (
	"fmt"
)

func main() {
	a := 1690                         // 表示1.69
	b := 1700                         // 表示1.70
	c := a * b                        // 结果应该是2873000表示 2.873
	fmt.Println(c)                    // 内部编码
	fmt.Println(float64(c) / 1000000) // 显示
}
