# 安装go


### Mac 安装
```shell
brew install go
brew install golang
# 查看go
brew info go
brew info golang
```


### CentOS 安装
```shell
# 预安装
$ yum -y install make autoconf automake cmake perl-CPAN libcurl-devel libtool gcc gcc-c++ glibc-headers zlib-devel git-lfs telnet ctags lrzsz jq expat-devel openssl-devel
# uname -a 查看系统，确定下载二进制文件包
# 下载安装包
$ wget https://golang.org/dl/go1.17.1.linux-amd64.tar.gz -O /tmp/go1.17.1.linux-amd64.tar.gz
# 如果以上命令被墙，也可以执行以下命令
$ wget https://marmotedu-1254073058.cos.ap-beijing.myqcloud.com/tools/go1.17.1.linux-amd64.tar.gz -O /tmp/go1.17.1.linux-amd64.tar.gz
$ mkdir -p $HOME/go
$ tar -xvzf /tmp/go1.17.1.linux-amd64.tar.gz -C $HOME/go
$ mv $HOME/go/go $HOME/go/go1.17.1

# 将下列变量追加到 $HOME/.bashrc 文件中
tee -a $HOME/.bashrc <<'EOF'
# Go envs
export GOVERSION=go1.17.1 # Go 版本设置
export GO_INSTALL_DIR=$HOME/go # Go 安装目录
export GOROOT=$GO_INSTALL_DIR/$GOVERSION # GOROOT 设置
export GOPATH=$WORKSPACE/golang # GOPATH 设置
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH # 将 Go 语言自带的和通过 go install 安装的二进制文件加入到 PATH 路径中
export GO111MODULE="on" # 开启 Go moudles 特性
export GOPROXY=https://goproxy.cn,direct # 安装 Go 模块时，代理服务器设置
export GOPRIVATE=
export GOSUMDB=off # 关闭校验 Go 依赖包的哈希值
EOF

# 将以上配置写入以下文件，并执行下
vim /etc/profile
source /etc/profile
```



### 官网安装
官网地址：https://golang.google.cn/dl/


### 查看go版本
```shell
go version
```











