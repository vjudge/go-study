# go env 环境变量

* $GOROOT : Go 语言安装根目录的路径，也就是 Go 在电脑上的安装位置。
* $GOPATH : 工作区目录的路径。包含 src、pkg、bin，用于存放源码文件，包文件，和可执行文件。
  - Go 语言的工作目录，可以是多个，每个目录代表 Go 语言的一个工作区。
  - 放置 Go 语言的源码文件，以及安装后的归档文件和可执行文件。
* $GOBIN : Go 程序生成的可执行文件的路径。
* $GOOS : 表示目标机器的操作系统，值可以是 darwin、freebsd、linux 或windows。
* $GO111MODULE
  * off : go 命令行将不会支持 module 功能，寻找依赖包的方式将会沿用旧版本那种通过 vendor 目录或者GOPATH模式来查找
  * on : go命令行会使用 modules，而一点也不会去GOPATH目录下查找
  * auto : 默认值，go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种情形：
    * 当前目录在 GOPATH/src 之外且该目录包含 go.mod 文件
    * 当前文件在包含 go.mod 文件的目录下面





















