# beego 使用
### 安装 beego
1. 下载并安装beego
```
go get -u github.com/astaxie/beego
go get -u github.com/beego/bee
```

2. 快速建立一个应用 hello
```
cd $GOPATH/src
bee new hello
cd hello
bee run
```

### 使用 beego
#### 1. [beego 命令](https://beego.me/docs/install/bee.md)
```
AVAILABLE COMMANDS
    version     Prints the current Bee version
    migrate     Runs database migrations
    api         Creates a Beego API application
    bale        Transforms non-Go files to Go source files
    fix         Fixes your application by making it compatible with newer versions of Beego
    dlv         Start a debugging session using Delve
    dockerize   Generates a Dockerfile for your Beego application
    generate    Source code generator
    hprose      Creates an RPC application based on Hprose and Beego frameworks
    new         Creates a Beego application
    pack        Compresses a Beego application into a single file
    rs          Run customized scripts
    run         Run the application by starting a local development server
    server      serving static content over HTTP on port
```

1. `version`  命令
    版本查看命令
    ```
    bee version
    ```

1. `new` 命令

    生成一个应用 hello
    ```
    bee new hello
    ```

1. `run` 命令

    在应用的根目录运行应用
    ```
    cd  $GOPATH/src/hello
    bee run
    ```

1. `api` 命令
    
    创建 API 应用
    ```
    bee api apiproject
    ```

    创建基于数据库表CURD的 API 应用
    ```
    # bee api [appname] [-tables=""] [-driver=mysql] [-conn="root:<password>@tcp(127.0.0.1:3306)/test"]
    bee api apidb -driver=mysql -conn="root:Hnbd@501-25@tcp(202.116.46.215:3306)/Service"

    # 运行并生成 swager
    cd apiproject
    bee run -gendoc=true -downdoc=true
    ```

1. `pack` 命令

    发布应用的时候打包，会把项目打包成 tar.gz 包
    ```
    cd  $GOPATH/src/hello
    bee run
    ```

1. [`migrate` 命令](https://beego.me/docs/install/bee.md#migrate-%E5%91%BD%E4%BB%A4)

    应用的数据库迁移命令，主要是用来每次应用升级，降级的SQL管理
    ```
    bee migrate rollback [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
    ```

1. `dockerize`命令

    生成Dockerfile文件来实现docker化应用
    ```
     bee dockerize -image="library/golang:1.6.4" -expose=9000
    ```

# 参考
1. beego 官网 . https://beego.me
2. beego github . https://github.com/astaxie/beego