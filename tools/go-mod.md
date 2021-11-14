# go mod


### 环境变量配置
```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOSUMDB=off
```


### 初始化
```shell
# 将会生成 go.mod 文件
go mod init <项目名>
```
go.mod 文件一旦创建后，它的内容将会被 go toolchain 全面掌控。go toolchain 会在各类命令执行时，比如 go get、go build、go mod 等修改和维护 go.mod 文件


### go mod 命令
* init :在当前目录初始化mod
* download : 下载依赖包
* edit : 编辑go.mod
* graph : 打印模块依赖图
* tidy : 拉取缺少的模块，移除不用的模块
* vendor :将依赖复制到vendor下
* verify : 验证依赖是否正确
* why : 解释为什么需要依赖


### 四个命令模块
* module : 语句指定包的名字（路径）
* require : 语句指定的依赖项模块
* replace : 语句可以替换依赖项模块
* exclude : 语句可以忽略依赖项模块


### 升级所有依赖
```shell
go get -u 
```


### 查看可升级的 package
```shell
go list -m -u all
```



