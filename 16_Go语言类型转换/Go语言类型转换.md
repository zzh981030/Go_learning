# Go 语言类型转换

---
类型转换用于将一种数据类型的变量转换为另外一种类型的变量。

Go 语言类型转换基本格式如下：
```go
type_name(expression)
```
type_name 为类型，expression 为表达式。

---
## 数值类型转换
将整型转换为浮点型：
```go
var a int = 10
var b float64 = float64(a)
```

---
## 字符串类型转换

* 注意：涉及到字符串类型转换的情况下，通过**strconv**函数实现

将一个字符串转换成另一个类型，可以使用以下语法：
```go
var str string = "10"
var num int
num, _ = strconv.Atoi(str)
```

以上代码将字符串变量 str 转换为整型变量 num。

注意，**strconv.Atoi** 函数返回两个值，第一个是转换后的整型值，第二个是可能发生的错误，我们可以使用空白标识符 **_** 来忽略这个错误

---
## 接口类型转换
接口类型转换有两种情况：**类型断言**和**类型转换**。

类型断言用于将接口类型转换为指定类型，其语法为：
```go
value.(type)
或者
value.(T)
```

其中 value 是接口类型的变量，type 或 T 是要转换成的类型。

如果类型断言成功，它将返回转换后的值和一个布尔值，表示转换是否成功。

类型转换用于将一个接口类型的值转换为另一个接口类型，其语法为：
```go
T(value)
T 是目标接口类型，value 是要转换的值。
```
