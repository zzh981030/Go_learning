参考链接：https://blog.csdn.net/yilovexing/article/details/121172745

len指的长度，cap指的容量

**make() 是 Go 语言内存分配的内置函数，默认有三个参数。**
```go
make(Type, len, cap)
```

* Type：数据类型，必要参数，Type 的值只能是 slice、 map、 channel 这三种数据类型。
* len：数据类型实际占用的内存空间长度，map、 channel 是可选参数，slice 是必要参数。
* cap：为数据类型提前预留的内存空间长度，可选参数。所谓的提前预留是当前为数据类型申请内存空间的时候，提前申请好额外的内存空间，这样可以避免二次分配内存带来的开销，大大提高程序的性能。

---
**为了能更好的理解这些参数的含义，我们先来看下 make() 的三种不同用法：**
```go
make(map[string]string)

make([]int, 2)

make([]int, 2, 4)
```

* 第一种，只传类型，不指定实际占用的内存空间和提前预留的内存空间，适用于 map 和 channel 。
* 第二种，指定实际占用的内存空间为 2，不指定提前预留的内存空间。
* 第三种，指定实际占用的内存空间为 2，指定提前预留的内存空间是 4。

---
看到这里你有没有这样的疑惑，既然在初始化的时候已经指定数据的大小了，那为什么还要指定预留的大小呢？这是因为 make() 使用的是一种动态数组算法，一开始先向操作系统申请一小块内存，这个就是 cap，等 cap 被 len 占用满以后就需要扩容，扩容就是动态数组再去向操作系统申请当前长度的两倍的内存，然后将旧数据复制到新内存空间中。
