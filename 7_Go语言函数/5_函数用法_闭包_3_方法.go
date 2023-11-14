//方法就是一个包含了接受者的函数
/*
Go 语言中同时有函数和方法。
一个方法就是一个包含了接受者的函数，接受者可以是<命名类型>或者<结构体类型>的<<一个值或者是一个指针>>。
所有给定类型的方法属于该类型的方法集。语法格式如下：
func (variable_name variable_data_type) function_name() [return_type]{
	// 函数体
}
*/
package main

import (
	"fmt"
	"math"
)

/* 定义结构体 */
//用来指定使用此结构体的函数中的数据类型
type Circle struct {
	radius float64
}

func main() {
	var c1 Circle     //将Circle变量实例化给c1
	c1.radius = 10.00 //c1就有了一些方法

	fmt.Println("圆的面积 = ", c1.getArea())
	fmt.Println("圆的圆的周长 = ", c1.getCircumference())
}

// 该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 { //求圆面积的函数
	//c.radius 即为 Circle 类型对象中的属性
	return math.Pi * c.radius * c.radius
}

func (d Circle) getCircumference() float64 { //求圆周长的函数
	return math.Pi * d.radius * 2
}
