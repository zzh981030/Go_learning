# Go 语言范围(Range)

----
Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
```go
for key, value := range oldMap {
    newMap[key] = value
}
```
以上代码中的 key 和 value 是可以省略。
如果只想读取 key，格式如下：
```go
    for key := range oldMap
```
或者这样：
```go
    for key, _ := range oldMap
```
如果只想读取 value，格式如下：
```go
    for _, value := range oldMap
```
