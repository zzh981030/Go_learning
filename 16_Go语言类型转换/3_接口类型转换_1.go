/*
以下实例中，我们定义了一个接口类型变量 i，并将它赋值为字符串 "Hello, World"。
然后，我们使用类型断言将 i 转换为字符串类型，并将转换后的值赋值给变量 str。
最后，我们使用 ok 变量检查类型转换是否成功，如果成功，我们打印转换后的字符串；否则，我们打印转换失败的消息。
*/

package main

import "fmt"

func main() {
	var i interface{} = "Hello, World"
	str, ok := i.(string)
	fmt.Println(str, ok)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
}
