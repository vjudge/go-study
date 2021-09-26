# init() 函数
用于包（package）的初始化。
对同一个package中不同文件是按文件名字符串比较“从小到大”顺序调用各文件中的init()函数。


### init函数的主要作用
* 初始化不能采用初始化表达式初始化的变量
* 程序运行前的注册
* 实现sync.Once功能
* 其他


### init函数的主要特点
* init函数没有输入参数、返回值
* 每个包可以有多个init函数，不能被其他函数调用
* 包的每个源文件也可以有多个init函数
* 同一个包的init执行顺序，Go没有明确定义，编程时要注意程序不要依赖这个执行顺序
* 不同包的init函数按照包导入的依赖关系决定执行顺序
* init函数先于main函数自动执行：所有被编辑器识别到的init函数都会在main函数执行前被调用
* 其他包的init是在main包的init调用之前被执行


### Go程序初始化
Go程序初始化先于main函数执行，由runtime进行初始化，初始化顺序如下：
1. 初始化导入的包（包的初始化顺序并不是按导入顺序（“从上到下”）执行的，runtime需要解析包依赖关系，没有依赖的包最先初始化，与变量初始化依赖关系类似，参见golang变量的初始化）
2. 初始化包作用域的变量（该作用域的变量的初始化也并非按照“从上到下、从左到右”的顺序，runtime解析变量依赖关系，没有依赖的变量最先初始化，参见golang变量的初始化）
3. 执行包的init函数


```
package main
                                                                                                        
import "fmt" 

var T int64 = a()

func init() {
   fmt.Println("init in main.go ")
}

func a() int64 {
   fmt.Println("calling a()")
   return 2
}
func main() {                  
   fmt.Println("calling main")     
}
// 初始化顺序：变量初始化->init()->main()
// calling a()
// init in main.go
// calling main
```

对于不同的package，如果不相互依赖的话，按照main包中”先import的后调用”的顺序调用其包中的init()，如果package存在依赖，则先调用最早被依赖的package中的init()，最后调用main函数。
![包引用关系](https://github.com/vjudge/go-study/blob/master/images/package-relation.png)









