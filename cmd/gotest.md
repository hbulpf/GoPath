# Go Test

go 语言的test代码写起来和运行起来都相对简单。

- 文件名要以 '_test' 结尾
- 测试函数以 'Test' 开头


```
//执行测试函数
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