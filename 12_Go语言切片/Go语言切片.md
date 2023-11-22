# Go 语言切片(Slice)
Go 语言切片是对数组的抽象。

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

---
## 定义切片
你可以声明一个未指定大小的数组来定义切片：
```go
var identifier []type
```
切片不需要说明长度。
或使用 ***make()*** 函数来创建切片:
```go
var slice1 []type = make([]type, len)

也可以简写为

slice1 := make([]type, len)
```
也可以指定容量，其中***capacity***为可选参数。
```go
make([]T, length, capacity)
```
这里 len 是数组的长度并且也是切片的初始长度。

---
## 切片初始化
```go
s :=[] int {1,2,3 }
```
直接初始化切片，***[]*** 表示是切片类型，***{1,2,3}*** 初始化值依次是 ***1,2,3***，其 ***cap=len=3***。
```go
s := arr[:]
```
初始化切片***s***，是数组 arr 的引用。
```go
s := arr[startIndex:endIndex]
```
将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片。
```go
s := arr[startIndex:]
```
默认 endIndex 时将表示一直到arr的最后一个元素。
```go
s := arr[:endIndex]
```

默认 startIndex 时将表示从 arr 的第一个元素开始。
```go
s1 := s[startIndex:endIndex]
```

通过切片 s 初始化切片 s1。
```go
s :=make([]int,len,cap)
```

通过内置函数***make()*** 初始化切片 ***s***，***[]int*** 标识为其元素类型为 int 的切片。
