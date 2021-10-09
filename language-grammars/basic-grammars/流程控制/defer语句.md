# defer语句

defer 语句仅能被放置在函数或方法中。它由关键字defer和一个调用表达式组成。

defer 语句就是被用来延迟执行代码的。这要延迟到该语句所在的函数即将执行结束的那一刻，无论结束执行的原因是什么。

这里的调用表达式所表示的既不能是对Go语言内建函数的调用，也不能是对Go语言标准库代码包unsafe中的那些函数的调用。实际上，满足上述条件的调用表达式被称为表达式语句。
```
// 当这条defer语句被执行的时候，其中的这条表达式语句并不会被立即执行。它的确切的执行时机是在其所属的函数（这里是readFile）的执行即将结束的那个时刻。
func readFile(path string) ([]byte, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    return ioutil.ReadAll(file)
}
```
无论readFile函数正常地返回了结果还是由于在其执行期间有运行时恐慌发生而被剥夺了流程控制权，其中的file.Close()都会在该函数即将退出那一刻被执行。这就更进一步地保证了资源的及时释放。

当一个函数中存在多个defer语句时，它们携带的表达式语句的执行顺序一定是它们的出现顺序的倒序。
* 当一个函数即将结束执行时，其中的写在最下边的defer函数调用会最先执行，其次是写在它上边、与它的距离最近的那个defer函数调用，以此类推，最上边的defer函数调用会最后一个执行。
```
func deferIt() {
    defer func() {
        fmt.Print(1)
    }()
    defer func() {
        fmt.Print(2)
    }()
    defer func() {
        fmt.Print(3)
    }()
    fmt.Print(4)
}
// 4321
```

defer携带的表达式语句代表的是对某个函数或方法的调用。这个调用可能会有参数传入，比如：fmt.Print(i + 1)。如果代表传入参数的是一个表达式，那么在defer语句被执行的时候该表达式就会被求值了.
```
func deferIt3() {
    f := func(i int) int {
        fmt.Printf("%d ",i)
        return i * 10
    }
    for i := 1; i < 5; i++ {
        defer fmt.Printf("%d ", f(i))
    }
}
// 1 2 3 4 40 30 20 10
```


```
func deferIt4() {
    for i := 1; i < 5; i++ {
        defer func() {
            fmt.Print(i)
        }()
    }
}   
// 5555
```




















