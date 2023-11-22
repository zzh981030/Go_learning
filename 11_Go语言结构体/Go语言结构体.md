# Go 语言结构体

---
Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。

结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。

结构体表示一项记录，比如保存图书馆的书籍记录，每本书有以下属性：

* Title ：标题
* Author ： 作者
* Subject：学科
* ID：书籍ID
---

## 定义结构体

结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：
```go
type struct_variable_type struct {
    member definition
    member definition
    ...
    member definition
}
```

一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：
```go
variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
```
---
## 访问结构体成员
如果要访问结构体成员，需要使用点号 . 操作符，格式为：
```go
结构体.成员名
```
---
## 结构体作为函数参数
你可以像其他数据类型一样将结构体类型作为参数传递给函数。并以以上实例的方式访问结构体变量：

---

## 结构体指针
你可以定义指向结构体的指针类似于其他指针变量，格式如下：
```go
var struct_pointer *Books
```
以上定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前：
```go
struct_pointer = &Book1
```
使用结构体指针访问结构体成员，使用 "." 操作符：
```go
struct_pointer.title
```
