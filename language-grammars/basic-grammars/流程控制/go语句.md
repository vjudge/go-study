# go 语句

go语句可以携带一条表达式语句。go语句的执行会很快结束，并不会对当前流程的进行造成阻塞或明显的延迟。
```
go fmt.Println("Go!")
```
go语句的执行与其携带的表达式语句的执行在时间上没有必然联系。能够确定的仅仅是后者会在前者完成之后发生。
在go语句被执行时，其携带的函数（也被称为go函数）以及要传给它的若干参数（如果有的话）会被封装成一个实体（即Goroutine），并被放入到相应的待运行队列中。Go语言的运行时系统会适时的从队列中取出待运行的Goroutine并执行相应的函数调用操作。

go函数的执行时间的不确定性。


runtime.Gosched函数的作用是让当前正在运行的Goroutine（这里是运行main函数的那个Goroutine）暂时“休息”一下，而让Go运行时系统转去运行其它的Goroutine。
```
package main

import (
    "fmt"
    "runtime"
)

func main() {
    go fmt.Println("Go!")
    runtime.Gosched()
}
```

sync.WaitGroup类型有三个方法可用——Add、Done和Wait。
* Add会使其所属值的一个内置整数得到相应增加。
* Done会使那个整数减1。
* Wait方法会使当前Goroutine阻塞直到那个整数为0。



















