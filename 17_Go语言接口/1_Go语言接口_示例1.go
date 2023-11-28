/*
在上面的例子中，我们定义了一个接口 Phone，接口里面有一个方法 call()。

然后我们在 main 函数里面定义了一个 Phone 类型变量，并分别为之赋值为 NokiaPhone 和 IPhone。

然后调用 call() 方法，
*/

package main

import (
	"fmt"
)

/* 定义接口 */
type Phone interface { //定义一个接口，命名为phone
	call()
}

/* 定义结构体 */
type NokiaPhone struct { //定义一个结构体，命名为NokiaPhone
}

/* 实现接口方法 */
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

/*------------------------------------*/

/* 定义结构体 */
type IPhone struct { //定义一个结构体，命名为IPhone
}

/* 实现接口方法 */
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone //定义接口变量

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}
