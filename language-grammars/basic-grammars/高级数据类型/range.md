# range


Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
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


```
//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
for i, c := range "go" {
    fmt.Println(i, c)
}
// 0 103
// 1 111
```























