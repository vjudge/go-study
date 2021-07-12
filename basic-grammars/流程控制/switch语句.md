# switch语句

在switch语句中，每个case会携带一个表达式。case表达式的结果类型并不一定是bool，它的结果类型需要与switch表达式的结果类型一致。switch语句会依据该表达式的结果与各个case表达式的结果是否相同来决定执行哪个分支。
当没有一个常规的case被选中的时候，default 就会被选中。
```
var name string
// 省略若干条语句
switch name {
case "Golang":
    fmt.Println("A programming language from Google.")
case "Rust":
    fmt.Println("A programming language from Mozilla.")
default:
    fmt.Println("Unknown!")
}  
```





























