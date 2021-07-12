# if语句

Go语言的流程控制主要包括条件分支、循环和并发。


if语句一般会由关键字if、条件表达式和由花括号包裹的代码块组成。代码块是包含了若干表达式和语句的序列。在Go语言中，代码块必须由花括号包裹。
```
if 100 > number {
    number += 3
} else if 100 < number {
    number -= 2
} else {
    fmt.Println("OK!")
}   
```

```
if number := 4; 100 > number {
    number += 3
} else if 100 < number {
    number -= 2
} else {
    fmt.Println("OK!")
}  
```
































































