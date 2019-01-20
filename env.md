# GoLang 环境搭建
Go是一门全新的静态类型开发语言，具有自动垃圾回收，丰富的内置类型,函数多返回值，错误处理，匿名函数,并发编程，反射等特性.
## 1. 下载 GoLang软件包
Windows / Linux / MacOS 下载地址 : [https://golang.org/dl/](https://golang.org/dl/)

## 2. 设置环境变量
### 2.1 环境变量说明
go命令依赖一个重要的环境变量 `$GOPATH`   
GOPATH允许多个目录，当有多个目录时，请注意分隔符，多个目录的时候Windows是分号;
当有多个GOPATH时默认将 `go get` 获取的包存放在第一个目录下。

1. `$GOROOT`
即golang 的安装路径

1. `$GOPATH` 目录约定有三个子目录
    + `src` 存放源代码(比如：.go .c .h .s等)，按照golang默认约定。`go run` 、 `go install` 等命令的当前工作路径（即在此路径下执行上述命令）。
    + `pkg` 编译时生成的中间文件（比如：.a），golang编译包时。
    + `bin` 编译后生成的可执行文件（为了方便，可以把此目录加入到 `$PATH` 变量中，如果有多个gopath，那么使用 `${GOPATH//://bin:}/bin` 添加所有的bin目录）

1. `$GOBIN`
执行 `go install` 会在 `$GOBIN` 目录下生成项目的二进制可执行文件。

1. go get   
`go install` 会做两件事：1. 从远程下载需要用到的包 ; 2. 执行go install

1. go build   
在project目录下执行 `go build` 会在当前目录下生成项目的二进制可执行文件。

1. go install   
`go install` 会生成可执行文件直接放到bin目录下。如果是一个普通的包，会被编译生成到pkg目录下该文件是.a结尾。   
注意: 如果设置了 `$GOBIN` 环境变量，生成的可执行文件就会放在 `$GOBIN` 目录下。如果没有设置 `$GOBIN` 环境变量,生成的可执行文件就会放在 `$GOPATH\bin` 目录下(此时默认 `$GOPATH\bin` 为 `$GOBIN`)。如果需要执行这些可执行文件,就要把 `$GOBIN` 或 `$GOPATH\bin` 加入到 `$PATH` 中。

1. go clean -i  
`go clean -i` 会清理 `$GOBIN` 和 项目当前目录下的二进制可执行文件。

1. go env   
    查看 go 环境
    ```
    PS E:\hsdocker\golang_faq> go env
    set GOARCH=amd64
    set GOBIN=
    set GOEXE=.exe
    set GOHOSTARCH=amd64
    set GOHOSTOS=windows
    set GOOS=windows
    set GOPATH=E:\Go_WorkSpace
    set GORACE=
    set GOROOT=D:\Go
    set GOTOOLDIR=D:\Go\pkg\tool\windows_amd64
    set GCCGO=gccgo
    set CC=gcc
    set GOGCCFLAGS=-m64 -mthreads -fmessage-length=0
    set CXX=g++
    set CGO_ENABLED=1
    set CGO_CFLAGS=-g -O2
    set CGO_CPPFLAGS=
    set CGO_CXXFLAGS=-g -O2
    set CGO_FFLAGS=-g -O2
    set CGO_LDFLAGS=-g -O2
    set PKG_CONFIG=pkg-config
    ```

1. 代码目录结构  
GOPATH 下的 src 目录就是接下来开发程序的主要目录，所有的源码都是放在这个目录下面，那么一般的做法就是一个目录一个项目，例如: `$GOPATH/src/hello` 表示 hello 这个应用包或可执行应用，这个根据package是main还是其他来决定，main就是可执行应用，其他就是应用包。
    ```
    go_project     // go_project为GOPATH目录
      -- bin
         -- myApp1  // 编译生成
         -- myApp2  // 编译生成
         -- myApp3  // 编译生成
      -- pkg
      -- src
         -- myApp1     // project1
            -- models
            -- controllers
            -- others
            -- main.go 
         -- myApp2     // project2
            -- models
            -- controllers
            -- others
            -- main.go 
         -- myApp3     // project3
            -- models
            -- controllers
            -- others
            -- main.go
    ```

### 2.2 Windows安装配置go环境
1. 首先下载linux下的go包：[https://golang.org/dl/](https://golang.org/dl/)
2. 下载安装包后直接安装即可。
3. 添加 `$GOROOT` 、`$GOPATH` 、`$GOBIN` 环境变量
4. 查看 go 环境
    ```
    go env
    ```

### 2.3 Linux安装配置go环境
1. 首先下载linux下的go包：[https://golang.org/dl/](https://golang.org/dl/)
2. 下载之后解压并移动到 /usr/local/go 也就是GOROOT
    ```
    tar -zxvf go1.9.5.linux-amd64.tar.gz -C /usr/local/
    ```
3. 设置GOPATH，添加PATH环境变量 `vi /etc/profile`
    ```
export GOROOT="/usr/local/go"
export GOPATH="$HOME/gopath"
export PATH="$PATH:$GOROOT/bin:$GOPATH/bin"
    ```

4. 生效 sbt 环境变量
```
source /etc/profile
```

5. 查看 go 环境
    ```
    go env
    ```

# 参考
1. Go语言之讲解GOROOT、GOPATH、GOBIN . https://www.cnblogs.com/pyyu/p/8032257.html


