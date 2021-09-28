# switch语句

在switch语句中，每个case会携带一个表达式。  
switch语句会依据该表达式的结果与各个case表达式的结果是否相同来决定执行哪个分支。
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



### case 表达式
* case 表达式一般由 case 关键字和一个表达式列表组成，表达式列表中的多个表达式之间需要有英文逗号,分割。
* case 表达式的结果类型并不一定是 bool，它的结果类型需要与 switch 表达式的结果类型一致。
    * 只要 switch 表达式的结果值与某个 case 表达式中的任意一个子表达式的结果值相等，该 case 表达式所属的 case 子句就会被选中。
    * 一旦某个 case 子句被选中，其中的附带在 case 表达式后边的那些语句就会被执行。与此同时，其他的所有case子句都会被忽略。
* 如果被选中的 case 子句附带的语句列表中包含了 fallthrough 语句，那么紧挨在它下边的那个 case 子句附带的语句也会被执行。
* switch语句对switch表达式的结果类型，以及各个case表达式中子表达式的结果类型要求相同。
    * 如果case表达式中子表达式的结果值是无类型的常量，那么它的类型会被自动地转换为switch表达式的结果类型。



### switch 对 case 表达式的约束
* switch语句不允许case表达式中的子表达式结果值存在相等的情况，不论这些结果值相等的子表达式，是否存在于不同的case表达式中。


























