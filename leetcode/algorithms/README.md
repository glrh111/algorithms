
## 算法

### 1. 带余除法 欧几里德除法

参考 `https://zh.wikipedia.org/wiki/%E5%B8%A6%E4%BD%99%E9%99%A4%E6%B3%95`

### 2. pseudo-random number

参考
+ `https://zh.wikipedia.org/wiki/%E4%BC%AA%E9%9A%8F%E6%9C%BA%E6%80%A7`
+ math/rand

实现
```go
func rand(n int) int {
	var (
		seed = time.Now().Nanosecond() // 种子值
		rnd int
	)
	rnd = seed * 1103515245 + 12345
	rnd = (rnd / 65536) % 32768
	return rnd % n
}
```

## 分类问题

### 1. 括号匹配问题

+ 20 给定的括号是否匹配？ 20
+ 32 给定s的最大匹配的子串 32
  + stack方法 将索引存入stack中
  + 动态规划？
+ 22 生成可匹配括号的串 22  
  