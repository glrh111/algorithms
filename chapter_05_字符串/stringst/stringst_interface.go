package stringst

type StringSTInterface interface {
	Put(key string, value interface{}) bool // bool 代表是否新增了值
	Get(key string) interface{}
	Delete(key string) bool
	Contains(key string) bool
	IsEmpty() bool
	LongestPrefixOf(s string) string // 返回值是ST中的key，它与s有最长的契合前缀
	KeysWithPreffix(s string) chan string // 返回前缀为 s 的所有键
	KeysThatMatch(s string) chan string   // 所有和 s 匹配的键，. 可以匹配任意字符
	Size() int
	Keys() chan string
}
