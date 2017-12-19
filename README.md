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

## git

### 1 删除敏感数据

`https://help.github.com/articles/removing-sensitive-data-from-a-repository/`

`git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch chapter_03_查找/exam/leipzig1M.txt'  --prune-empty --tag-name-filter cat -- --all `






