//在类型转换中，我们必须保证要转换的值和目标接口类型之间是兼容的，否则编译器会报错。
/*
以上实例中，我们定义了一个 Writer 接口和一个实现了该接口的结构体 StringWriter。
然后，我们将 StringWriter 类型的指针赋值给 Writer 接口类型的变量 w。接着，我们使用类型转换将 w 转换为 StringWriter 类型，并将转换后的值赋值给变量 sw。
最后，我们使用 sw 访问 StringWriter 结构体中的字段 str，并打印出它的值。
*/

package main

import "fmt"

type Writer interface { //定义一个接口，名为Writer
	Write([]byte) (int, error)
}

type StringWriter struct { //定义一个结构体
	str string
}

func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}

func main() {
	var w Writer = &StringWriter{} //定义一个实现Writer接口的结构体StringWriter，并将其指针赋值为变量w
	sw := w.(*StringWriter)        //使用类型转换将w转换为StringWriter类型，并将结果赋值给变量sw
	sw.str = "Hello, World"        //使用sw访问StringWriter结构体中的dtr字段
	fmt.Println(sw.str)
}
