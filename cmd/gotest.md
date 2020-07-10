# Go Test

go 语言的test代码写起来和运行起来都相对简单。

- 文件名要以 '_test' 结尾
- 测试函数以 'Test' 开头


```
//执行所有测试函数
go test ./main/testMe.go ./main/testMe_test.go -v 
//执行指定函数的测试函数
go test ./main/testMe.go ./main/testMe_test.go -run='s1' -v  
```


## 例子

下面我们看一个例子


被测函数 testMe.go 如下



```go
package main

func f1(n int) int{
    if n == 0{
        return 0
    }

    if n == 1{
        return 1
    }
    return f1(n - 1) +f1(n - 2)
}

func f2(n int) int{
    if n == 0{
        return 0
    }

    if n == 1{
        return 2
    }

    return f2(n - 1) +f2(n - 2)
}

func s1(s string) int{
    if s == ""{
        return 0
    }

    n := 1
    for range s{
        n++
    }
    return n
}

func s2(s string) int{
    return len(s)
}
```

测试函数testMe_test.go如下



```go
package main

import "testing"

func TestS1(t *testing.T){
    if s1("123456789") != 9{
        t.Error(`s1("123456789") != 9`)
    }

    if s1("") != 0{
        t.Error(`s1("") != 0`)
    }
}

func TestS2(t *testing.T){
    if s2("123456789") != 9{
        t.Error(`s2("123456789") != 9`)
    }

    if s2("") != 0{
        t.Error(`s2("") != 0`)
    }
}

func TestF1(t *testing.T){
    if f1(0) != 0{
        t.Error(`f1(0) != 0`)
    }
    if f1(1) != 1{
        t.Error(`f1(1) != 1`)
    }

    if f1(2) != 1{
        t.Error(`f1(2) != 1`)
    }

    if f1(10) != 55{
        t.Error(`f1(10) != 55`)
    }
}

func TestF2(t *testing.T){
    if f2(0) != 0{
        t.Error(`f1(0) != 0`)
    }
    if f2(1) != 1{
        t.Error(`f1(1) != 1`)
    }

    if f2(2) != 1{
        t.Error(`f1(2) != 1`)
    }

    if f2(10) != 55{
        t.Error(`f1(10) != 55`)
    }
}
```

执行测试函数

```go
D:\git\Go>go test ./main/testMe.go ./main/testMe_test.go -v
=== RUN   TestS1
--- FAIL: TestS1 (0.00s)
        testMe_test.go:7: s1("123456789") != 9
=== RUN   TestS2
--- PASS: TestS2 (0.00s)
=== RUN   TestF1
--- PASS: TestF1 (0.00s)
=== RUN   TestF2
--- FAIL: TestF2 (0.00s)
        testMe_test.go:47: f1(1) != 1
        testMe_test.go:51: f1(2) != 1
        testMe_test.go:55: f1(10) != 55
FAIL
FAIL    command-line-arguments  0.108s
```

如果只想执行指定函数的测试函数如下

```go
D:\git\Go>go test ./main/testMe.go ./main/testMe_test.go -run='s1' -v
testing: warning: no tests to run
PASS
ok      command-line-arguments  0.106s [no tests to run]
```

# 详细使用



## 单元测试——测试和验证代码的框架

要开始一个单元测试， 在命名文件时需要让文件必须以`_test`结尾。默认的情况下，`go test`命令不需要任何的参数，它会自动把你源码包下面所有 test 文件测试完毕，当然你也可以带上参数。 

这里介绍几个常用的参数：

- -bench regexp 执行相应的 benchmarks，例如 -bench=.；
- -cover 开启测试覆盖率；
- -run regexp 只运行 regexp 匹配的函数，例如 -run=Array 那么就执行包含有 Array 开头的函数；
- -v 显示测试的详细命令。

 单元测试源码文件可以由多个测试用例组成，每个测试用例函数需要以`Test`为前缀，例如： 

```
 func TestXXX( t *testing.T ) 
```

- 测试用例文件不会参与正常源码编译，不会被包含到可执行文件中。
- 测试用例文件使用`go test`指令来执行，没有也不需要 main() 作为函数入口。所有在以`_test`结尾的源码内以`Test`开头的函数会自动被执行。
- 测试用例可以不传入 *testing.T 参数。

 helloworld 的测试代码 

```
package code11_3

import "testing"

func TestHelloWorld(t *testing.T) {
    t.Log("hello world")
}
```

代码说明如下：

- 第 5 行，单元测试文件 (*_test.go) 里的测试入口必须以 Test 开始，参数为 *testing.T 的函数。一个单元测试文件可以有多个测试入口。
- 第 6 行，使用 testing 包的 T 结构提供的 Log() 方法打印字符串。

#### 1) 单元测试命令行

单元测试使用 go test 命令启动，例如：

```
$ go test helloworld_test.go
ok          command-line-arguments        0.003s
$ go test -v helloworld_test.go
=== RUN   TestHelloWorld
--- PASS: TestHelloWorld (0.00s)
        helloworld_test.go:8: hello world
PASS
ok          command-line-arguments        0.004s
```

代码说明如下：

- 第 1 行，在 go test 后跟 helloworld_test.go 文件，表示测试这个文件里的所有测试用例。
- 第 2 行，显示测试结果，ok 表示测试通过，command-line-arguments 是测试用例需要用到的一个包名，0.003s 表示测试花费的时间。
- 第 3 行，显示在附加参数中添加了`-v`，可以让测试时显示详细的流程。
- 第 4 行，表示开始运行名叫 TestHelloWorld 的测试用例。
- 第 5 行，表示已经运行完 TestHelloWorld 的测试用例，PASS 表示测试成功。
- 第 6 行打印字符串 hello world。

#### 2) 运行指定单元测试用例

`go test`指定文件时默认执行文件内的所有测试用例。可以使用参数选择需要的测试用例单独执行，参考下面的代码。

 一个文件包含多个测试用例 



```
package code11_3

import "testing"

func TestA(t *testing.T) {
    t.Log("A")
}

func TestAK(t *testing.T) {
    t.Log("AK")
}

func TestB(t *testing.T) {
    t.Log("B")
}

func TestC(t *testing.T) {
    t.Log("C")
}
```

 这里指定 TestA 进行测试： 

```
$ go test -v -run TestA select_test.go
=== RUN   TestA
--- PASS: TestA (0.00s)
        select_test.go:6: A
=== RUN   TestAK
--- PASS: TestAK (0.00s)
        select_test.go:10: AK
PASS
ok          command-line-arguments        0.003s
```

TestA 和 TestAK 的测试用例都被执行，原因是`-run``-run TestA$`

#### 3) 标记单元测试结果

当需要终止当前测试用例时，可以使用 FailNow，参考下面的代码。
测试结果标记（具体位置是）

```
func TestFailNow(t *testing.T) {
    t.FailNow()
}
```

 还有一种只标记错误不终止测试的方法，代码如下： 

```
func TestFail(t *testing.T) {

    fmt.Println("before fail")

    t.Fail()

    fmt.Println("after fail")
}
```

 测试结果如下： 

```
=== RUN   TestFail
before fail
after fail
--- FAIL: TestFail (0.00s)
FAIL
exit status 1
FAIL        command-line-arguments        0.002s
```

从日志中看出，第 5 行调用 Fail() 后测试结果标记为失败，但是第 7 行依然被程序执行了。

#### 4) 单元测试日志

每个测试用例可能并发执行，使用 testing.T 提供的日志输出可以保证日志跟随这个测试上下文一起打印输出。testing.T 提供了几种日志输出方法，详见下表所示。

| 方  法 | 备  注                           |
| ------ | -------------------------------- |
| Log    | 打印日志，同时结束测试           |
| Logf   | 格式化打印日志，同时结束测试     |
| Error  | 打印错误日志，同时结束测试       |
| Errorf | 格式化打印错误日志，同时结束测试 |
| Fatal  | 打印致命日志，同时结束测试       |
| Fatalf | 格式化打印致命日志，同时结束测试 |

开发者可以根据实际需要选择合适的日志。

## 基准测试——获得代码内存占用和运行效率的性能数据

基准测试可以测试一段程序的运行性能及耗费 CPU 的程度。Go语言中提供了基准测试框架，使用方法类似于单元测试，使用者无须准备高精度的计时器和各种分析工具，基准测试本身即可以打印出非常标准的测试报告。

#### 1) 基础测试基本使用

下面通过一个例子来了解基准测试的基本使用方法。
基准测试

```
package code11_3

import "testing"

func Benchmark_Add(b *testing.B) {
    var n int
    for i := 0; i < b.N; i++ {
        n++
    }
}
```

 这段代码使用基准测试框架测试加法性能。第 7 行中的 b.N 由基准测试框架提供。测试代码需要保证函数可重入性及无状态，也就是说，测试代码不使用全局变量等带有记忆性质的[数据结构](http://c.biancheng.net/data_structure/)。避免多次运行同一段代码时的环境不一致，不能假设 N 值范围。

 使用如下命令行开启基准测试： 

```
$ go test -v -bench=. benchmark_test.go
goos: linux
goarch: amd64
Benchmark_Add-4           20000000         0.33 ns/op
PASS
ok          command-line-arguments        0.700s
```

代码说明如下：

- 第 1 行的`-bench=.`表示运行 benchmark_test.go 文件里的所有基准测试，和单元测试中的`-run`类似。
- 第 4 行中显示基准测试名称，2000000000 表示测试的次数，也就是 testing.B 结构中提供给程序使用的 N。“0.33 ns/op”表示每一个操作耗费多少时间（纳秒）。


注意：Windows 下使用 go test 命令行时，应写为。基准测试框架对一个测试用例的默认测试时间是 1 秒。开始测试时，当以 Benchmark 开头的基准测试用例函数返回时还不到 1 秒，那么 testing.B 中的 N 值将按 1、2、5、10、20、50……递增，同时以递增后的值重新调用基准测试用例函数。通过参数可以自定义测试时间，例如：

```
$ go test -v -bench=. -benchtime=5s benchmark_test.go
goos: linux
goarch: amd64
Benchmark_Add-4           10000000000                 0.33 ns/op
PASS
ok          command-line-arguments        3.380s
```

#### 4) 测试内存

基准测试可以对一段代码可能存在的内存分配进行统计，下面是一段使用字符串格式化的函数，内部会进行一些分配操作。

```
func Benchmark_Alloc(b *testing.B) {

    for i := 0; i < b.N; i++ {
        fmt.Sprintf("%d", i)
    }
}
```

 在命令行中添加`-benchmem`参数以显示内存分配情况，参见下面的指令： 

```
$ go test -v -bench=Alloc -benchmem benchmark_test.go
goos: linux
goarch: amd64
Benchmark_Alloc-4 20000000 109 ns/op 16 B/op 2 allocs/op
PASS
ok          command-line-arguments        2.311s
```

代码说明如下：

- 第 1 行的代码中`-bench`后添加了 Alloc，指定只测试 Benchmark_Alloc() 函数。
- 第 4 行代码的“16 B/op”表示每一次调用需要分配 16 个字节，“2 allocs/op”表示每一次调用有两次分配。

开发者根据这些信息可以迅速找到可能的分配点，进行优化和调整。

#### 5) 控制计时器

有些测试需要一定的启动和初始化时间，如果从 Benchmark() 函数开始计时会很大程度上影响测试结果的精准性。testing.B 提供了一系列的方法可以方便地控制计时器，从而让计时器只在需要的区间进行测试。我们通过下面的代码来了解计时器的控制。
基准测试中的计时器控制（具体位置是）：

```
func Benchmark_Add_TimerControl(b *testing.B) {

    // 重置计时器
    b.ResetTimer()

    // 停止计时器
    b.StopTimer()

    // 开始计时器
    b.StartTimer()

    var n int
    for i := 0; i < b.N; i++ {
        n++
    }
}
```

 从 Benchmark() 函数开始，Timer 就开始计数。StopTimer() 可以停止这个计数过程，做一些耗时的操作，通过 StartTimer() 重新开始计时。ResetTimer() 可以重置计数器的数据。

计数器内部不仅包含耗时数据，还包括内存分配的数据。 



# 参考

1. [GO语言学习](http://c.biancheng.net/golang/intro/)