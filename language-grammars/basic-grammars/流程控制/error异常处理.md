# error

Go语言的函数可以一次返回多个结果。

error是Go语言内置的一个接口类型。只要一个类型的方法集合包含了名为Error、无参数声明且仅声明了一个string类型的结果的方法，就相当于实现了error接口。
```
type error interface { 
    Error() string
}
```

创造error，只需调用标准库代码包errors的New函数即可。errors.New是一个很常用的函数。在Go语言标准库的代码包中有很多由此函数创建出来的错误值，比如os.ErrPermission、io.EOF等变量的值。
```
if path == "" {
    return nil, errors.New("The parameter is invalid!")
}  
```
























