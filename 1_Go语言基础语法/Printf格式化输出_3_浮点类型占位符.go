//https://blog.csdn.net/gudongkun1121/article/details/131697332

package main

import "fmt"

func main() {
	var f1 float64 = 123.789
	var f2 float64 = 12345678910.78999

	fmt.Printf("指数为二的幂的无小数科学计数法：%b \n", f1)
	fmt.Printf("科学计数法：%e \n", f1)
	fmt.Printf("科学计数法大写：%E \n", f1)
	fmt.Printf("有小数点无指数，也就是常规的浮点数格式。默认宽度和精度：%f \n", f1)
	fmt.Printf("宽度9默认精度：%9f \n", f1)
	fmt.Printf("默认宽度，精度保留2位小数%.2f \n", f1)
	fmt.Printf("宽度为9，精度保留2位小数%9.2f \n", f1)
	fmt.Printf("宽度为9，精度保留0位小数%9.f \n", f1)
	fmt.Printf("根据情况自动选择 %%e 或者 %%f 来输出，来产生更紧凑的输出（末位无0）：%g %g \n", f1, f2)
	fmt.Printf("根据情况自动选择 %%E 或者 %%f 来输出，来产生更紧凑的输出（末位无0）：%G %G\n", f1, f2)
	fmt.Printf("用十六进制的方式来表示：%x \n", f1)
	fmt.Printf("用十六进制的方式来表示，大写：%X \n", f1)

}
