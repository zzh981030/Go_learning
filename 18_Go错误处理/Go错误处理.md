# Go 错误处理

---
Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
error 类型是一个接口类型，这是它的定义：
```go
type error interface {
    Error() string
}
```

我们可以在编码中通过实现 error 接口类型来生成错误信息。
函数通常在最后的返回值中返回错误信息。使用 errors.New 可返回一个错误信息：
```go
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // 实现
}
```

在下面的例子中，我们在调用 Sqrt 的时候传递的一个负数，然后就得到了 non-nil 的 error 对象，将此对象与 nil 比较，结果为 true，所以 fmt.Println(fmt 包在处理 error 时会调用 Error 方法)被调用，以输出错误，请看下面调用的示例代码：
```go
result, err:= Sqrt(-1)
if err != nil {
    fmt.Println(err)
}
```
