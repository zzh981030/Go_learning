package main

import "fmt"

// 定义一个结构体，包含以下几个变量，并声明变量类型
type Books_1 struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {

	// 创建一个新的结构体
	fmt.Println(Books_1{"Go_语言", "www.runoob.com", "Go_语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books_1{title: "Go_语言", author: "www.runoob.com", subject: "Go_语言教程", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books_1{title: "Go_语言", author: "www.runoob.com"})
}
