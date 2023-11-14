/*
在 switch 语句中使用 break：
*/
package main

import "fmt"

func main() {
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday.")
	case "Tuesday":
		fmt.Println("It's Tuesday.")
		break // 跳出 switch 语句
	case "Wednesday":
		fmt.Println("It's Wednesday.")
	}
}
