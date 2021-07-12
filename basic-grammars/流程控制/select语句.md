# select语句

select语句属于条件分支流程控制方法，只能用于通道。它可以包含若干条case语句，并根据条件选择其中的一个执行。select语句中的case关键字只能后跟用于通道的发送操作的表达式以及接收操作的表达式或语句。
```
// 如果该select语句被执行时通道ch1和ch2中都没有任何数据，那么只有default case会被执行。
// 如果在当时有数据的通道多于一个，那么Go语言会通过一种伪随机的算法来决定哪一个case将被执行。
ch1 := make(chan int, 1)
ch2 := make(chan int, 1)
select {
    case e1 := <-ch1:
        fmt.Printf("1th case is selected. e1=%v.\n", e1)
    case e2 := <-ch2:
        fmt.Printf("2th case is selected. e2=%v.\n", e2)
    default:
        fmt.Println("No data!")
} 
```

如果一条select语句中不存在default case， 并且在被执行时其中的所有case都不满足执行条件，那么它的执行将会被阻塞！当前流程的进行也会因此而停滞。直到其中一个case满足了执行条件，执行才会继续。

未被初始化的通道会使操作它的case永远满足不了执行条件。对于针对它的发送操作和接收操作来说都是如此。

break语句也可以被包含在select语句中的case语句中。它的作用是立即结束当前的select语句的执行，不论其所属的case语句中是否还有未被执行的语句。



















