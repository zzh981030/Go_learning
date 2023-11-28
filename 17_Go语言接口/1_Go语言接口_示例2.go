/*
以上实例中，我们定义了一个 Shape 接口，它定义了一个方法 area()，该方法返回一个 float64 类型的面积值。
然后，我们定义了两个结构体 Rectangle 和 Circle，它们分别实现了 Shape 接口的 area() 方法。
在 main() 函数中，我们首先定义了一个 Shape 类型的变量 s，然后分别将 Rectangle 和 Circle 类型的实例赋值给它，并通过 area() 方法计算它们的面积并打印出来，
*/

package main

import "fmt"

type Shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var s Shape

	s = Rectangle{width: 10, height: 5}
	fmt.Printf("矩形面积: %f\n", s.area())

	s = Circle{radius: 3}
	fmt.Printf("圆形面积: %f\n", s.area())
}
