/*
在 Go 语言中，break 语句用于终止当前循环或者 switch 语句的执行，并跳出该循环或者 switch 语句的代码块。

break 语句可以用于以下几个方面：

用于循环语句中跳出循环，并开始执行循环之后的语句。
break 在 switch 语句中在执行一条 case 后跳出语句的作用。
break 可应用在 select 语句中。
在多重循环中，可以用标号 label 标出想 break 的循环。
*/
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 当 i 等于 5 时跳出循环
		}
		fmt.Println(i)
	}
}
