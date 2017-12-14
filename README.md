# algorithms
Algorithms Fourth Edition

实验数据从这里下载：`https://github.com/ekis/oldAlgorithms-2012`

这里也是 `https://github.com/ekis/oldAlgorithms-2012`

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




