# range

Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。


### slice
```
//这是我们使用range去求一个slice的和。使用数组跟这个很类似
nums := []int{2, 3, 4}
sum := 0
// 空白符"_"省略索引
for _, num := range nums {
    sum += num
}
fmt.Println("sum:", sum)
// sum: 9
```


### 字符串
```
//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
for i, c := range "go" {
    fmt.Println(i, c)
}
// 0 103
// 1 111
```


### 数组
* range表达式只会在 for 语句开始执行时被求值一次，无论后边会有多少次迭代
* range表达式的求值结果会被复制，也就是说，被迭代的对象是range表达式结果值的副本而不是原值。
```go
arys := [...]int{1, 2, 3, 4, 5}
maxa := len(arys) - 1
for i, ary := range arys {
	if i == maxa {
		arys[0] = 2 * ary
	} else {
		arys[i + 1] = 2 * ary
	}
}
// arys: [10 2 4 6 8]
```


### 切片
```go
slice1 := []int{1, 2, 3, 4, 5}
maxs := len(slice1) - 1
for i, ele := range slice1 {
	if i == maxs {
		slice1[0] = 2 * ele
	} else {
		slice1[i + 1] = 2 * ele
	}
}
// slice1:  [32 2 4 8 16]
```
















