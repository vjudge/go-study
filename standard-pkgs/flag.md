# flag

用于接收和解析命令参数。


### 导入
```go
import "flag"
```



### flag.StringVar()
```go
flag.StringVar(&name, "name", "everyone", "The greeting object.")
// 第一个参数: 存储该命令参数值的地址
// 第二个参数: 指定该命令参数值的名称
// 第三个参数: 指定在未追加该命令参数时的默认值
// 第四个参数: 该命令参数的简短说明
```



### flag.String()
返回一个已经分配好用于存储命令参数值的地址
```go
name := flag.String("name", "everyone", "The greeting object.")
```



### flag.BoolVar()
```go
flag.BoolVar(&h, "h", false, "this help")
```


### flag.Bool()
```go
h := flag.BoolVar("h", false, "this help")
```



### flag.IntVar()
```go
flag.BoolVar(&num, "num", 123, "this help")
```


### flag.Int()
```go
num := flag.BoolVar("num", 123, "this help")
```



### flag.Var()
```go
flag.Var(&flagVal, "name", "help message")
```



### flag.Parse()
解析命令参数，把他们的值赋给相应的变量。  
该函数的调用，必须在所有命令参数存储载体的声明，和设置之后，并在读取任何命令参数值之前。



### flag.Usage = func () { ... }
自定义命令源码文件的参数使用说明
```go
func usage () {
    fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`)
    flag.PrintDefaults()
}
```



### flag.CommandLine
默认情况下的命令参数容器。通过对其重新赋值，可以更深层次的定制当前命令源码文件的参数使用说明。







