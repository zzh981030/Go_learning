package main

import "fmt"

const (
	i = 1 << iota
	j = 3 << iota //3对应的二进制为11，iota为1，因此左移1位，成为110，即6=3*2^1次方.注意：第二行重新引用了iota，则只与第几行有关，与上一行的值无关
	k             //在上一行基础上，11往左移2位，成为1100，即12=3*3=3*2^3次方
	l             //在上一行基础上，11往左移3位，成为11000，即24=3*8=3*2^3次方
)

//注意： <<n == *(2^n)

func main() {
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}
