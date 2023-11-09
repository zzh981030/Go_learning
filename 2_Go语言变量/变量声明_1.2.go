//变量声明
//第一种，指定变量类型，如果没有初始化，则变量默认为零值。

package main

import "fmt"

func main() {
	var i int     //int类型默认为0
	var f float64 //float64默认为0
	var b bool    //bool类型默认为false
	var s string  //string类型默认为空
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
