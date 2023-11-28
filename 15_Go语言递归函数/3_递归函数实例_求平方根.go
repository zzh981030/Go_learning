/*
以下实例通过 Go 语言使用递归方法实现求平方根的代码：
*/

package main

import (
	"fmt"
)

func sqrtRecursive(x, guess, prevGuess, epsilon float64) float64 {
	diff := guess*guess - x //计算猜测值与输入值的均方误差
	fmt.Println("误差精度= ", diff)
	if diff < epsilon && -diff < epsilon { //判断猜测值与输入值的均方误差，误差小于精度的情况下，返回猜测值
		return guess
	}

	newGuess := (guess + x/guess) / 2 //多次迭代，逼近新的猜测值
	if newGuess == prevGuess {
		return guess
	}

	return sqrtRecursive(x, newGuess, guess, epsilon)
}

//func sqrt(x float64) float64 {
//	return sqrtRecursive(x, 1.0, 0.0, 1e-9)
//}

func main() {
	x := 99.0
	//result := sqrt(x)
	result := sqrtRecursive(x, 1.0, 0.0, 1e-9)
	fmt.Printf("%.2f 的平方根为 %.6f\n", x, result)
}

/*
以上实例中，sqrtRecursive 函数使用递归方式实现平方根的计算。

sqrtRecursive 函数接受四个参数：

	x 表示待求平方根的数
	guess 表示当前猜测的平方根值
	prevGuess 表示上一次的猜测值
	epsilon 表示精度要求（即接近平方根的程度）
递归的终止条件是当前猜测的平方根与上一次猜测的平方根非常接近，差值小于给定的精度 epsilon。

在 sqrt 函数中，我们调用 sqrtRecursive 来计算平方根，并传入初始值和精度要求，然后在 main 函数中，我们调用 sqrt 函数来求解平方根，并将结果打印出来。
*/
