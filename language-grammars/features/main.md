# 入口函数（主函数）

要生成Go语言可执行程序，必须有main和package包，且必须该包下有main()函数。

### 必须是main包
```
package main
```


### 必须是main方法
```
func main () {
   ...
   os.Exit(0) // -1
}
```
1. main函数不支持任何返回值，可以通过 os.Exit 来返回状态。
2. main函数不支持传入参数。在程序中直接通过 os.Args 获取命令行参数。


注：文件名不一定是main.go