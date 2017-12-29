# algorithms
Algorithms Fourth Edition

实验数据从这里下载：`https://github.com/ekis/oldAlgorithms-2012`

这里也是 `https://github.com/ekis/oldAlgorithms-2012`

https://raw.githubusercontent.com/ekis/oldAlgorithms-2012/master/z-algs4-common/data-sets/leipzig1M.txt

## golang 的类型别名 `type wocao int`

wocao 不继承 int 的方法，两者的方法互相隔离。

## 单引号，双引号的区别

## Golang 怎么读取各种文件

`http://wiki.jikexueyuan.com/project/the-way-to-go/12.2.html`

## Golang 的官方文档

`https://github.com/golang/go/wiki`

## 遇到的问题汇总

### 1 `Test killed with quit: ran too long`

参考：
+ https://stackoverflow.com/questions/27778280/test-killed-with-quit-ran-too-long
+ https://golang.org/pkg/time/#ParseDuration

go test -timeout 20m （Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".）

### 2 copy struct 

`*b = *a // copy a to b`

### 3 copy slice

```go
keys := make([]*Comparable, size) // 需要分配空间
copy(keys, this.keys[loRank:hiIndex])
```

### 4 slice 作为形参

有空看看这个 https://blog.golang.org/slices

slice是一个结构体，*slice 是指向结构体的指针，这一点需要区分。

### 5 计算key的hash值

https://stackoverflow.com/questions/13582519/how-to-generate-hash-number-of-a-string-in-go

```go
package main

import (
        "fmt"
        "hash/fnv"
)

func hash(s string) uint32 {
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}

func main() {
        fmt.Println(hash("HelloWorld"))
        fmt.Println(hash("HelloWorld."))
}
```

### 6 迭代器实现

参考 `https://studygolang.com/articles/1695`

### 7 Golang 的单引号

rune 类型的，rune 是 int32 的别名。可以用来访问字符串的原始字符，而不是字节。

```go
package main

import "fmt"

func main() {
 s := "我操你大爷1"
 
 fmt.Println(len(s)) // 16
 
 for _, value := range s {
     fmt.Println(value, len(s))
 } // 打印出来6个数
 
 fmt.Println(s[7]) // 189 第七个字节
 
 fmt.Println(len([]rune(s))) // 6
 
 for _, value := range []rune(s) {
     fmt.Println(string(value), len([]rune(s)))
 }
}

```

## git

### 1 删除敏感数据

`https://help.github.com/articles/removing-sensitive-data-from-a-repository/`

`git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch chapter_03_查找/exam/leipzig1M.txt'  --prune-empty --tag-name-filter cat -- --all `







