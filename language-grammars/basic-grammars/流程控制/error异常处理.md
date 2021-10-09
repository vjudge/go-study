# error

error 类型其实是一个接口类型，也是一个 Go 语言的内建类型。



error是Go语言内置的一个接口类型。只要一个类型的方法集合包含了名为Error、无参数声明且仅声明了一个string类型的结果的方法，就相当于实现了error接口。
```
type error interface { 
    Error() string
}
```


### errors.New()
errors.New是一个很常用的创建 error 的函数。它会返回一个包含了错误信息的 error 类型值。  
在Go语言标准库的代码包中有很多由此函数创建出来的错误值，比如os.ErrPermission、io.EOF等变量的值。
```go
if path == "" {
    return nil, errors.New("The parameter is invalid!")
}  
// fmt.Printf 函数如果发现被打印的值是一个error类型的值，那么就会去调用它的Error方法。
```


### errors使用
在函数声明的结果列表的最后，声明一个该类型的结果，同时在调用这个函数之后，先判断它返回的最后一个结果值是否“不为nil”。如果这个值不是 nil，那么就进入错误处理流程，否则就继续正常流程。
```go

```



### 卫述语句
卫述语句就是被用来检查后续操作的前置条件，并进行相应处理的语句。




















