# for语句


### fori
for语句代表着循环。一条语句通常由关键字for、初始化子句、条件表达式、后置子句和以花括号包裹的代码块组成。
```
for i := 0; i < 10; i++ {
    fmt.Print(i, " ")
}  
```
可以省略掉初始化子句、条件表达式、后置子句中的任何一个或多个，不过起到分隔作用的分号一般需要被保留下来，除非在仅有条件表达式或三者全被省略时分号才可以被一同省略。



### forr
for语句可以用range子句替换掉for子句。range子句包含一个或两个迭代变量（用于与迭代出的值绑定）、特殊标记:=或=、关键字range以及range表达式。其中，range表达式的结果值的类型应该是能够被迭代的，包括：字符串类型、数组类型、数组的指针类型、切片类型、字典类型和通道类型。
```
// 每次迭代出的第一个值所代表的是第二个字符值经编码后的第一个字节在该字符串经编码后的字节数组中的索引值。
for i, v := range "Go语言" {
    fmt.Printf("%d: %c\n", i, v)
}
// 一个中文字符在经过UTF-8编码之后会表现为三个字节
// 0: G
// 1: o
// 2: 语
// 5: 言 
```

对于数组值、数组的指针值和切片之来说，range子句每次也会迭代出两个值。其中，第一个值会是第二个值在被迭代值中的索引，而第二个值则是被迭代值中的某一个元素。同样的，迭代是以索引递增的顺序进行的。

对于字典值来说，range子句每次仍然会迭代出两个值。显然，第一个值是字典中的某一个键，而第二个值则是该键对应的那个值。注意，对字典值上的迭代，Go语言是不保证其顺序的。

携带range子句的for语句还可以应用于一个通道值之上。其作用是不断地从该通道值中接收数据，不过每次只会接收一个值。注意，如果通道值中没有数据，那么for语句的执行会处于阻塞状态。无论怎样，这样的循环会一直进行下去。直至该通道值被关闭，for语句的执行才会结束。


break：被执行时会使其所属的for语句的执行立即结束。
continue：被执行时会使当次迭代被中止（当次迭代的后续语句会被忽略）而直接进入到下一次迭代。























