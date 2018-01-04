package compress

/*
   标准输入字节流
 */
type BinaryStdInInterface interface {
	ReadBool() bool
	ReadChar() uint8
	IsEmpty() bool
	Close()
}

/*
   标准输出字节流
 */
type BinaryStdOutInterface interface {
	Write(b bool)
	WriteChar(c uint8)
	Close()
}

/*
   字母表 Alphabet
   字符串表示使用 []rune
 */
type AlphabetInterface interface {
	ToChar(index int) (r rune) // 返回字母表index处的字符
	ToIndex(r rune) (index int) // 字符转化为在其中的索引
	Contains(r rune) bool // 字母表中是否有该字母
	R() int // 字母表字符数量
	LgR() int // 需要多少位表示索引。
	ToIndices(r []rune) []int // 转化为R位的整数数组 每一个字符执行 ToIndex...
	ToChars([]int) []rune     // 转化为字符串表示
}
